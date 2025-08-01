package containers

import (
	"strings"

	"github.com/NucleoFusion/cruise/internal/messages"
	"github.com/NucleoFusion/cruise/internal/styles"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Containers struct {
	Width     int
	Height    int
	List      *ContainerList
	IsLoading bool
}

func NewContainers(w int, h int) *Containers {
	return &Containers{
		Width:     w,
		Height:    h,
		IsLoading: true,
		List:      NewContainerList(w-4, h-7-strings.Count(styles.ContainersText, "\n")),
	}
}

func (s *Containers) Init() tea.Cmd {
	return s.List.Init()
}

func (s *Containers) Update(msg tea.Msg) (*Containers, tea.Cmd) {
	switch msg := msg.(type) {
	case messages.ContainerReadyMsg:
		s.IsLoading = false

		var cmd tea.Cmd
		s.List, cmd = s.List.Update(msg)
		return s, cmd
	}

	var cmd tea.Cmd
	s.List, cmd = s.List.Update(msg)
	return s, cmd
}

func (s *Containers) View() string {
	return lipgloss.JoinVertical(lipgloss.Center,
		styles.TextStyle().Render(styles.ContainersText), s.GetListText())
}

func (s *Containers) GetListText() string {
	if s.IsLoading {
		return lipgloss.Place(s.Width-2, s.Height-3-strings.Count(styles.ContainersText, "\n"),
			lipgloss.Center, lipgloss.Top, "Loading...")
	}

	return lipgloss.NewStyle().Padding(1).Render(s.List.View())
}
