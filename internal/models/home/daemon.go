package home

import (
	"github.com/NucleoFusion/cruise/internal/styles"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Daemon struct {
	Width  int
	Height int
}

func NewDaemon(w int, h int) *Daemon {
	return &Daemon{
		Width:  w,
		Height: h,
	}
}

func (s Daemon) Init() tea.Cmd { return nil }

func (s Daemon) Update(msg tea.Msg) (Daemon, tea.Cmd) {
	return s, nil
}

func (s Daemon) View() string {
	return styles.SubpageStyle().Render(lipgloss.Place(s.Width/3-2, s.Height/3-2,
		lipgloss.Center, lipgloss.Center, "Daemon"))
}
