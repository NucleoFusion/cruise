package home

import (
	"fmt"

	"github.com/NucleoFusion/cruise/internal/docker"
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
	return styles.SubpageStyle().PaddingTop(1).PaddingLeft(4).Render(lipgloss.JoinVertical(lipgloss.Center,
		styles.TitleStyle().Render("Daemon Status"),
		lipgloss.Place(s.Width/4, s.Height/3-4,
			lipgloss.Left, lipgloss.Center, s.FormattedView())))
}

func (s Daemon) FormattedView() string {
	info, err := docker.GetDaemonInfo()
	if err != nil {
		return "Error Getting Daemon Status"
	}

	return fmt.Sprintf(`
Docker Daemon Status: Running (%s)

Uptime: %s | OS: %s
		`, info.Version, info.Uptime, info.OS)
}
