package root

import (
	"github.com/NucleoFusion/cruise/internal/enums"
	"github.com/NucleoFusion/cruise/internal/messages"
	"github.com/NucleoFusion/cruise/internal/models/fzf"
	"github.com/NucleoFusion/cruise/internal/models/home"
	tea "github.com/charmbracelet/bubbletea"
)

type Root struct {
	Width          int
	Height         int
	CurrentPage    enums.PageType
	IsLoading      bool
	PageFzf        fzf.FuzzyFinder
	IsChangingPage bool
	Home           *home.Home
}

func NewRoot() *Root {
	return &Root{
		CurrentPage: enums.Home,
		IsLoading:   true,
	}
}

func (s *Root) Init() tea.Cmd { return nil }

func (s *Root) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case messages.FzfSelection:
		s.IsChangingPage = false
		return s, nil
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return s, tea.Quit
		case tea.KeyShiftTab:
			s.IsChangingPage = true
			return s, nil
		}
	case tea.WindowSizeMsg:
		s.Width = msg.Width
		s.Height = msg.Height

		s.PageFzf = fzf.NewFzf([]string{"Home", "TBD", "TBD"}, msg.Width, msg.Height)
		s.Home = home.NewHome(msg.Width, msg.Height)

		cmd := tea.Batch(s.Home.Init())

		s.IsLoading = false
		return s, cmd
	}

	if s.IsChangingPage {
		var cmd tea.Cmd
		s.PageFzf, cmd = s.PageFzf.Update(msg)
		return s, cmd
	}

	switch s.CurrentPage {
	case enums.Home:
		var cmd tea.Cmd
		s.Home, cmd = s.Home.Update(msg)
		return s, cmd
	}

	return s, nil
}

func (s *Root) View() string {
	if s.IsLoading {
		return "\nLoading..."
	}

	if s.IsChangingPage {
		return s.PageFzf.View()
	}

	switch s.CurrentPage {
	case enums.Home:
		return s.Home.View()
	}

	return "Cruise - A TUI Docker Client"
}
