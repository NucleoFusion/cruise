package home

import (
	"fmt"

	"github.com/NucleoFusion/cruise/internal/docker"
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
	return styles.SubpageStyle().PaddingTop(1).PaddingLeft(4).Render(lipgloss.JoinVertical(lipgloss.Center,
		styles.TitleStyle().Render("Docker Stats"),
		lipgloss.Place(s.Width/4-2, s.Height/3-4,
			lipgloss.Left, lipgloss.Center, s.GetFormattedView())))
}

func (s QuickStats) GetFormattedView() string {
	vols := docker.GetNumVolumes()
	cntnr := docker.GetNumContainers()
	imgs := docker.GetNumImages()
	ntwrks := docker.GetNumNetworks()

	return fmt.Sprintf(`
Containers: %d

Images:     %d

Volumes:    %d

Networks:   %d
		`, cntnr, imgs, vols, ntwrks)
}
