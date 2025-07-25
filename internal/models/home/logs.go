package home

import (
	"github.com/NucleoFusion/cruise/internal/styles"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Logs struct {
	Width  int
	Height int
}

func NewLogs(w int, h int) *Logs {
	return &Logs{
		Width:  w,
		Height: h,
	}
}

func (s Logs) Init() tea.Cmd { return nil }

func (s Logs) Update(msg tea.Msg) (Logs, tea.Cmd) {
	return s, nil
}

func (s Logs) View() string {
	return styles.SubpageStyle().Render(lipgloss.Place(s.Width*2/3-1, s.Height/3-2,
		lipgloss.Center, lipgloss.Center, "Logs"))
}
