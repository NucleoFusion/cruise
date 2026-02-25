package home

import (
	"context"
	"log"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/cruise-org/cruise/internal/messages"
	"github.com/cruise-org/cruise/pkg/runtimes"
	"github.com/cruise-org/cruise/pkg/styles"
	"github.com/cruise-org/cruise/pkg/types"
)

var tickCmd = func() tea.Cmd {
	return tea.Tick(time.Second*3, func(_ time.Time) tea.Msg {
		return messages.HomeLogsTick{}
	})
}

type Logs struct {
	Width        int
	Height       int
	Events       []types.Log
	EventMonitor *types.Monitor
	IsLoading    bool
	Length       int
}

func NewLogs(w int, h int) *Logs {
	return &Logs{
		Width:     w,
		Height:    h,
		IsLoading: true,
		Length:    h - 6,
		Events:    make([]types.Log, 0),
	}
}

func (s *Logs) Init() tea.Cmd {
	return tea.Batch(tickCmd(),
		func() tea.Msg {
			m, err := runtimes.RuntimeSrv.RuntimeLogs(context.Background())
			if err != nil {
				return messages.ErrorMsg{Msg: err.Error()}
			}
			return messages.HomeLogsMonitor{Monitor: m}
		})
}

func (s *Logs) Update(msg tea.Msg) (*Logs, tea.Cmd) {
	switch msg := msg.(type) {
	case messages.HomeLogsMonitor:
		s.EventMonitor = msg.Monitor
		s.IsLoading = false
		return s, nil

	case messages.HomeLogsTick:
		log.Println("[Home UI] Home Logs Ticked")
		if s.IsLoading {
			return s, tickCmd()
		}

		buf := make([]types.Log, 0)
		for {
			select {
			case v := <-s.EventMonitor.Incoming:
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
	}
	return s, nil
}

func (s Logs) View() string {
	return styles.SubpageStyle().PaddingTop(1).PaddingLeft(4).Render(lipgloss.JoinVertical(lipgloss.Center,
		styles.TitleStyle().Render("Event Logs"),
		lipgloss.NewStyle().
			PaddingTop(1).
			Width(s.Width-6).   //-6 from padding(4) and border(2)
			Height(s.Height-4). //-4 from title(1) border(2) and padding(1)
			Align(lipgloss.Left, lipgloss.Center).
			Render(s.FormattedView())))
}

func (s *Logs) FormattedView() string {
	if s.IsLoading {
		return "Loading Logs....\n"
	}

	if len(s.Events) == 0 {
		return "No Logs yet, waiting....\n"
	}

	text := ""
	events := s.Events
	for _, msg := range events {
		text += runtimes.FormatLog(msg) + "\n"
	}

	return text
}
