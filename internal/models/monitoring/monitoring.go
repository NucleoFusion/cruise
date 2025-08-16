package monitoring

import (
	"context"
	"strings"
	"time"

	"github.com/NucleoFusion/cruise/internal/colors"
	"github.com/NucleoFusion/cruise/internal/docker"
	"github.com/NucleoFusion/cruise/internal/keymap"
	"github.com/NucleoFusion/cruise/internal/messages"
	"github.com/NucleoFusion/cruise/internal/styles"
	"github.com/NucleoFusion/cruise/internal/utils"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/docker/docker/api/types/events"
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
	Help      help.Model
	Events    []*events.Message
	EventChan <-chan *events.Message
	ErrChan   <-chan error
	IsLoading bool
	Length    int
}

func NewMonitoring(w int, h int) *Monitoring {
	vp := viewport.New(w, h-3-strings.Count(styles.MonitoringText, "\n"))
	vp.Style = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(colors.Load().Lavender).
		Padding(1).Foreground(colors.Load().Text)

	eventChan, errChan := docker.RecentEventStream(h/3 - 6)

	return &Monitoring{
		Width:     w,
		Height:    h,
		Help:      help.New(),
		Vp:        vp,
		Length:    h - 5 - strings.Count(styles.MonitoringText, "\n"),
		EventChan: eventChan,
		ErrChan:   errChan,
	}
}

func (s *Monitoring) Init() tea.Cmd {
	return tea.Batch(s.Sub())
}

func (s *Monitoring) Sub() tea.Cmd {
	return tea.Every(2*time.Second, func(_ time.Time) tea.Msg {
		return s.PollEvents()()
	})
}

func (s *Monitoring) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case messages.NewEvents:
		s.Events = append(s.Events, msg.Events...)

		if s.IsLoading {
			s.IsLoading = false
		}

		s.Vp.SetYOffset(len(s.Events))

		return s, s.Sub()
	}
	return s, nil
}

func (s *Monitoring) View() string {
	if s.IsLoading {
		s.Vp.SetContent("Loading...")
	} else {
		s.Vp.SetContent(s.FormattedView())
	}

	return lipgloss.JoinVertical(lipgloss.Center,
		styles.TextStyle().Render(styles.MonitoringText), s.Vp.View(), s.Help.View(keymap.NewDynamic([]key.Binding{})))
}

func (s *Monitoring) FormattedView() string {
	if s.IsLoading {
		return "Loading Logs....\n"
	}

	if len(s.Events) == 0 {
		return "No Logs yet, waiting....\n"
	}

	text := ""
	events := s.Events
	for _, msg := range events {
		text += docker.FormatDockerEventVerbose(*msg) + "\n"
	}

	return text
}

func (s *Monitoring) PollEvents() tea.Cmd {
	return func() tea.Msg {
		evs := make([]*events.Message, 0, s.Length)

		select {
		case err := <-s.ErrChan:
			return utils.ReturnError("Monitoring Page", "Error Querying Events", err)
		default:
			for i := 0; i < s.Length; i++ {
				select {
				case ev := <-s.EventChan:
					evs = append(evs, ev)
				default:
					return messages.NewEvents{
						Events: evs,
					}
				}
			}

			return messages.NewEvents{
				Events: evs,
			}
		}
	}
}
