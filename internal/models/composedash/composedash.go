package composedash

import (
	"strings"
	"time"

	"github.com/NucleoFusion/cruise/internal/compose"
	"github.com/NucleoFusion/cruise/internal/docker"
	"github.com/NucleoFusion/cruise/internal/messages"
	"github.com/NucleoFusion/cruise/internal/styles"
	"github.com/NucleoFusion/cruise/internal/utils"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Dash struct {
	Width    int
	Height   int
	OverVp   viewport.Model
	HealthVp viewport.Model
	TrendVp  viewport.Model
	// Details
	OverviewDetails *OverviewDetails
}

func NewComposeDash(w, h int) *Dash {
	ht := h - strings.Count(styles.ComposeText, "\n") - 4 - 1 // removing help height and the spacing respectively
	return &Dash{
		Width:           w,
		Height:          h,
		OverVp:          NewOverview(w/3, ht),
		TrendVp:         NewTrends(w/3, ht),
		HealthVp:        NewHealth(w/3, ht),
		OverviewDetails: &OverviewDetails{Width: w / 3},
	}
}

func (s *Dash) Init() tea.Cmd {
	return tea.Tick(0, func(_ time.Time) tea.Msg {
		prj, err := compose.GetNumProjects()
		if err != nil {
			return utils.ReturnError("Compose Dashboard", "Error Getting Projects", err)
		}

		srv, err := compose.GetNumServices()
		if err != nil {
			return utils.ReturnError("Compose Dashboard", "Error Getting Services", err)
		}

		ntw, err := docker.GetNumNetworks()
		if err != nil {
			return utils.ReturnError("Compose Dashboard", "Error Getting Networks", err)
		}

		return messages.ComposeOverview{
			Projects:   prj,
			Services:   srv,
			Containers: docker.GetNumContainers(),
			Volumes:    docker.GetNumVolumes(),
			Networks:   ntw,
		}
	})
}

func (s *Dash) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case messages.ComposeOverview:
		s.OverviewDetails.Containers = msg.Containers
		s.OverviewDetails.Services = msg.Services
		s.OverviewDetails.Projects = msg.Projects
		s.OverviewDetails.Networks = msg.Networks
		s.OverviewDetails.Volumes = msg.Volumes

		s.OverVp.SetContent(s.OverviewDetails.View())

		return s, nil
	}

	return s, nil
}

func (s *Dash) View() string {
	return lipgloss.JoinVertical(lipgloss.Center,
		lipgloss.PlaceHorizontal(s.Width, lipgloss.Center, styles.TextStyle().Render(styles.ComposeText)),
		"\n",
		lipgloss.JoinHorizontal(lipgloss.Center, s.OverVp.View(), s.TrendVp.View(), s.HealthVp.View()),
	)
}
