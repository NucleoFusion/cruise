package home

import (
	"github.com/NucleoFusion/cruise/internal/styles"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type QuickStats struct {
	Width  int
	Height int
}

func NewQuickStats(w int, h int) *QuickStats {
	return &QuickStats{
		Width:  w,
		Height: h,
	}
}

func (s QuickStats) Init() tea.Cmd { return nil }

func (s QuickStats) Update(msg tea.Msg) (QuickStats, tea.Cmd) {
	return s, nil
}

func (s QuickStats) View() string {
	return styles.SubpageStyle().Render(lipgloss.Place(s.Width/3-2, s.Height/3-2,
		lipgloss.Center, lipgloss.Center, "Quick Stats"))
}
