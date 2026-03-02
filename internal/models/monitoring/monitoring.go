package monitoring

import (
	"context"
	"sort"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/cruise-org/cruise/internal/messages"
	styledhelp "github.com/cruise-org/cruise/internal/models/help"
	"github.com/cruise-org/cruise/internal/utils"
	"github.com/cruise-org/cruise/pkg/colors"
	"github.com/cruise-org/cruise/pkg/keymap"
	"github.com/cruise-org/cruise/pkg/runtimes"
	"github.com/cruise-org/cruise/pkg/styles"
	"github.com/cruise-org/cruise/pkg/types"
	"github.com/lithammer/fuzzysearch/fuzzy"
)

type LogStreamer struct {
	ctx    context.Context
	cancel context.CancelFunc
	lines  chan string
}

type Monitoring struct {
	Width     int
	Height    int
	Vp        viewport.Model
	Ti        textinput.Model
	Keymap    keymap.MonitorMap
	Help      styledhelp.StyledHelp
	Events    []types.Log
	Filtered  []types.Log
	Monitor   *types.Monitor
	IsLoading bool
	Length    int
}

var tickCmd = func() tea.Cmd {
	return tea.Tick(time.Second*3, func(_ time.Time) tea.Msg {
		return messages.MonitoringTick{}
	})
}

func NewMonitoring(w int, h int) *Monitoring {
	vp := viewport.New(w-2, h-8-strings.Count(styles.MonitoringText, "\n"))
	vp.Style = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(colors.Load().FocusedBorder).
		Padding(1).Foreground(colors.Load().Text)

	ti := textinput.New()
	ti.Width = w - 14
	ti.Prompt = " Search: "
	ti.Placeholder = "Press '/' to search..."

	ti.PromptStyle = lipgloss.NewStyle().Foreground(colors.Load().FocusedBorder)
	ti.PlaceholderStyle = lipgloss.NewStyle().Foreground(colors.Load().PlaceholderText)
	ti.TextStyle = styles.TextStyle()

	return &Monitoring{
		Width:     w,
		Height:    h,
		Help:      styledhelp.NewStyledHelp(keymap.NewMonitorMap().Bindings(), w-2),
		Keymap:    keymap.NewMonitorMap(),
		Vp:        vp,
		Ti:        ti,
		IsLoading: true,
		Events:    make([]types.Log, 0),
		Length:    h - 6 - strings.Count(styles.MonitoringText, "\n"), // -6 for styled help and ti
	}
}

func (s *Monitoring) Init() tea.Cmd {
	return tea.Batch(tickCmd(),
		func() tea.Msg {
			m, err := runtimes.RuntimeSrv.RuntimeLogs(context.Background())
			if err != nil {
				return messages.ErrorMsg{Msg: err.Error()}
			}
			return messages.MonitoringMonitor{Monitor: m}
		})
}

func (s *Monitoring) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case messages.MonitoringMonitor:
		s.Monitor = msg.Monitor
		s.IsLoading = false
		return s, nil

	case messages.MonitoringTick:
		if s.IsLoading {
			return s, tickCmd()
		}

		buf := make([]types.Log, 0)
		for {
			select {
			case v := <-s.Monitor.Incoming:
				buf = append(buf, v)
			default:
				goto done
			}
		}

	done:
		for _, v := range buf {
			s.Events = append(s.Events, v)
			if len(s.Events) == s.Height-5 {
				s.Events = s.Events[1:]
			}
		}

		return s, tickCmd()
	case tea.KeyMsg:
		if s.Ti.Focused() {
			if key.Matches(msg, s.Keymap.ExitSearch) {
				s.Ti.Blur()
				return s, nil
			}
			var cmd tea.Cmd
			s.Ti, cmd = s.Ti.Update(msg)
			s.Filter(s.Ti.Value())
			return s, cmd
		}
		switch {
		case key.Matches(msg, keymap.QuickQuitKey()):
			return s, tea.Quit
		case key.Matches(msg, s.Keymap.Search):
			s.Ti.Focus()
			return s, nil
		case key.Matches(msg, s.Keymap.Export):
			arr := make([]string, 0)
			for _, v := range s.Events {
				arr = append(arr, runtimes.FormatLog(v))
			}

			err := runtimes.Export(arr, "monitoring")
			if err != nil {
				return s, utils.ReturnError("Monitoring", "Error Exporting", err)
			}

			return s, utils.ReturnMsg("Monitoring", "Exported Successfully", "exported events to export dir.")
		}
	}
	return s, nil
}

func (s *Monitoring) View() string {
	s.Vp.SetContent(s.FormattedView())

	return styles.SceneStyle().Render(
		lipgloss.JoinVertical(lipgloss.Center,
			styles.TextStyle().Padding(1, 0).Render(styles.MonitoringText),
			styles.PageStyle().Render(s.Ti.View()),
			s.Vp.View(), s.Help.View()))
}

func (s *Monitoring) FormattedView() string {
	if s.IsLoading {
		return "Loading Logs....\n"
	}

	if len(s.Events) == 0 {
		return "No Logs yet, waiting....\n"
	}

	text := ""
	for _, msg := range s.GetLogs() {
		text += runtimes.FormatLog(msg) + "\n"
	}

	return text
}

func (s *Monitoring) GetLogs() []types.Log {
	if s.Ti.Value() == "" {
		return s.Events
	} else {
		return s.Filtered
	}
}

func (s *Monitoring) Filter(val string) {
	formatted := make([]string, len(s.Events))
	originals := make([]types.Log, len(s.Events))

	for i, v := range s.Events {
		str := runtimes.FormatLog(v)
		formatted[i] = str
		originals[i] = v
	}

	ranked := fuzzy.RankFindFold(val, formatted)
	sort.Sort(ranked)

	result := make([]types.Log, len(ranked))
	for i, r := range ranked {
		result[i] = originals[r.OriginalIndex]
	}

	s.Filtered = result
}
