package projects

import (
	"log"
	"strings"

	"github.com/NucleoFusion/cruise/internal/keymap"
	"github.com/NucleoFusion/cruise/internal/messages"
	styledhelp "github.com/NucleoFusion/cruise/internal/models/help"
	"github.com/NucleoFusion/cruise/internal/styles"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Projects struct {
	Width     int
	Height    int
	List      *ProjectList
	Keymap    keymap.ContainersMap
	Help      styledhelp.StyledHelp
	IsLoading bool
}

func NewProjects(w int, h int) *Projects {
	return &Projects{
		Width:     w,
		Height:    h,
		IsLoading: true,
		List:      NewProjectList(w-4, h-3-strings.Count(styles.ProjectsText, "\n")),
		Help:      styledhelp.NewStyledHelp(keymap.NewContainersMap().Bindings(), w),
	}
}

func (s *Projects) Init() tea.Cmd {
	return s.List.Init()
}

func (s *Projects) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	log.Println("Updating Projects")
	defer log.Println("Completed Projects Update")
	switch msg := msg.(type) {
	case messages.ProjectsReadyMsg:
		s.IsLoading = false

		var cmd tea.Cmd
		s.List, cmd = s.List.Update(msg)
		return s, cmd
	case tea.KeyMsg:
		if s.List.Ti.Focused() {
			var cmd tea.Cmd
			s.List, cmd = s.List.Update(msg)
			return s, cmd
		}
	}

	var cmd tea.Cmd
	s.List, cmd = s.List.Update(msg)
	return s, cmd
}

func (s *Projects) View() string {
	log.Println("Viewing Projects")
	defer log.Println("Printed Projects")

	return lipgloss.JoinVertical(lipgloss.Center,
		styles.TextStyle().Render(styles.ProjectsText), s.GetListText(), s.Help.View())
}

func (s *Projects) GetListText() string {
	log.Println("Got Text")
	defer log.Println("Returned Text")
	if s.IsLoading {
		return lipgloss.Place(s.Width-2, s.Height-4-strings.Count(styles.ContainersText, "\n"),
			lipgloss.Center, lipgloss.Top, "Loading...")
	}

	return lipgloss.NewStyle().PaddingLeft(1).Render(s.List.View())
}
