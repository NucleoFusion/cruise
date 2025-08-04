package home

import (
	"time"

	"github.com/NucleoFusion/cruise/internal/docker"
	"github.com/NucleoFusion/cruise/internal/messages"
	"github.com/NucleoFusion/cruise/internal/styles"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/docker/docker/api/types/events"
)

type Logs struct {
	Width     int
	Height    int
	Events    []*events.Message
	EventChan chan *events.Message
	IsLoading bool
	Length    int
}

func NewLogs(w int, h int) *Logs {
	return &Logs{
		Width:     w,
		Height:    h,
		IsLoading: true,
		Length:    h/3 - 6,
		EventChan: docker.RecentEventStream(h/3 - 6),
	}
}

func (s *Logs) Init() tea.Cmd {
	return tea.Batch(s.Sub())
}

func (s *Logs) Sub() tea.Cmd {
	return tea.Every(2*time.Second, func(_ time.Time) tea.Msg {
		return s.PollEvents()()
	})
}

func (s *Logs) Update(msg tea.Msg) (*Logs, tea.Cmd) {
	switch msg := msg.(type) {
	case messages.NewEvents:
		s.Events = append(s.Events, msg.Events...)

		if len(s.Events) > s.Length {
			s.Events = s.Events[len(s.Events)-s.Length:]
		}
		if s.IsLoading {
			s.IsLoading = false
		}

		return s, s.Sub()
	}
	return s, nil
}

func (s Logs) View() string {
	return styles.SubpageStyle().PaddingTop(1).PaddingLeft(4).Render(lipgloss.JoinVertical(lipgloss.Center,
		styles.TitleStyle().Render("Event Logs"),
		lipgloss.Place(s.Width*2/3-4, s.Height/3-4,
			lipgloss.Left, lipgloss.Bottom, s.FormattedView())))
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
		text += docker.FormatDockerEvent(*msg) + "\n"
	}

	return text
}

func (s *Logs) PollEvents() tea.Cmd {
	return func() tea.Msg {
		evs := make([]*events.Message, 0, s.Length)
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
