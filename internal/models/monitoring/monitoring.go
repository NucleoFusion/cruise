package monitoring

import (
	"strings"

	"github.com/NucleoFusion/cruise/internal/colors"
	"github.com/NucleoFusion/cruise/internal/keymap"
	"github.com/NucleoFusion/cruise/internal/styles"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Monitoring struct {
	Width  int
	Height int
	Vp     viewport.Model
	Help   help.Model
}

func NewMonitoring(w int, h int) *Monitoring {
	vp := viewport.New(w, h-3-strings.Count(styles.MonitoringText, "\n"))
	vp.Style = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(colors.Load().Lavender).
		Padding(1).Foreground(colors.Load().Text)
	vp.SetContent("Logs Page")
	return &Monitoring{
		Width:  w,
		Height: h,
		Help:   help.New(),
		Vp:     vp,
	}
}

func (s *Monitoring) Init() tea.Cmd {
	return nil
}

func (s *Monitoring) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return s, nil
}

func (s *Monitoring) View() string {
	return lipgloss.JoinVertical(lipgloss.Center,
		styles.TextStyle().Render(styles.MonitoringText), s.Vp.View(), s.Help.View(keymap.NewDynamic([]key.Binding{})))
}
