package projects

import (
	"fmt"
	"time"

	"github.com/NucleoFusion/cruise/internal/compose"
	"github.com/NucleoFusion/cruise/internal/messages"
	"github.com/NucleoFusion/cruise/internal/styles"
	"github.com/NucleoFusion/cruise/internal/types"
	"github.com/NucleoFusion/cruise/internal/utils"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ProjectDetails struct {
	Width     int
	Height    int
	Summary   *types.ProjectSummary
	Inspect   *types.Project
	Services  []string
	IsLoading bool
	Current   int
	// Viewports
	DetailsVp  viewport.Model
	ServicesVp viewport.Model
	ResourceVp viewport.Model
	YmlVp      viewport.Model
}

func NewProjectDetails(w, h int, s *types.ProjectSummary) *ProjectDetails {
	// Details VP
	dvp := viewport.New(w/3, h/2)
	dvp.Style = styles.PageStyle().Padding(1, 2)
	// dvp.SetContent(getLabelView(v, labels, w))

	// Services VP
	svp := viewport.New(w, h/2)
	svp.Style = styles.PageStyle().Padding(1, 2)
	// dvp.SetContent(getLabelView(v, labels, w))

	// Resources VP
	rvp := viewport.New(w/3, h/2)
	rvp.Style = styles.PageStyle().Padding(1, 2)
	// dvp.SetContent(getLabelView(v, labels, w))

	// YAML Vp
	yvp := viewport.New(w/3, h/2)
	yvp.Style = styles.PageStyle().Padding(1, 2)
	// dvp.SetContent(getLabelView(v, labels, w))

	srvc := make([]string, 0, len(s.Services))
	for k := range s.Services {
		srvc = append(srvc, k)
	}

	return &ProjectDetails{
		Width:      w,
		Height:     h,
		Summary:    s,
		Current:    0,
		Services:   srvc,
		ResourceVp: rvp,
		ServicesVp: svp,
		DetailsVp:  dvp,
		YmlVp:      yvp,
		IsLoading:  true,
	}
}

func (s *ProjectDetails) Init() tea.Cmd {
	return tea.Tick(0, func(_ time.Time) tea.Msg {
		project, err := compose.Inspect(s.Summary)
		if err != nil {
			return utils.ReturnError("Project Details ", "Error Inspecting Project "+s.Summary.Name, err)
		}

		return messages.ProjectInspectResult{
			Project: project,
		}
	})
}

func (s *ProjectDetails) Update(msg tea.Msg) (*ProjectDetails, tea.Cmd) {
	switch msg := msg.(type) {
	case messages.ProjectInspectResult:
		s.Inspect = msg.Project

		s.YmlVp.SetContent(s.GetYAMLView())
		s.ResourceVp.SetContent(s.GetResourcesView())
		s.DetailsVp.SetContent(s.GetDetailsView())
		s.ServicesVp.SetContent(s.GetServicesView())

		s.IsLoading = false
		return s, nil
	case tea.KeyMsg:
		// TODO: Use Keymap
		switch msg.String() {
		case "up":
			if s.Current > 0 {
				s.Current--
				s.ServicesVp.SetContent(s.GetServicesView())
			}
			return s, nil
		case "down":
			if s.Current < len(s.Services)-1 {
				s.Current++
				s.ServicesVp.SetContent(s.GetServicesView())
			}
			return s, nil
		}
	}
	return s, nil
}

func (s *ProjectDetails) View() string {
	if s.IsLoading {
		return lipgloss.Place(s.Width, s.Height, lipgloss.Center, lipgloss.Center, "Loading")
	}

	return lipgloss.JoinVertical(lipgloss.Center,
		lipgloss.JoinHorizontal(lipgloss.Center, s.DetailsVp.View(), s.ResourceVp.View(), s.YmlVp.View()),
		s.ServicesVp.View(),
	)
}

func (s *ProjectDetails) GetDetailsView() string {
	w := s.Width

	text := fmt.Sprintf("%s %s \n\n%s %s \n\n%s %s \n\n%s %s \n\n%s %s \n\n%s %s \n\n%s %s",
		styles.DetailKeyStyle().Render(" Name: "), styles.TextStyle().Render(s.Inspect.Name),
		styles.DetailKeyStyle().Render(" Status: "), styles.TextStyle().Render(utils.Shorten(compose.Status(s.Inspect), w/3-15)),
		styles.DetailKeyStyle().Render(" Last Updated: "), styles.TextStyle().Render(utils.Shorten(compose.StartedAt(s.Inspect), w/3-15)),
		styles.DetailKeyStyle().Render(" Containers: "), styles.TextStyle().Render(utils.Shorten(fmt.Sprintf("%d", s.Summary.Containers), w/3-15)),
		styles.DetailKeyStyle().Render(" Services: "), styles.TextStyle().Render(utils.Shorten(fmt.Sprintf("%d", s.Summary.NumServices()), w/3-15)),
		styles.DetailKeyStyle().Render(" Networks: "), styles.TextStyle().Render(utils.Shorten(fmt.Sprintf("%d", s.Summary.Networks), w/3-15)),
		styles.DetailKeyStyle().Render(" Volumes: "), styles.TextStyle().Render(utils.Shorten(fmt.Sprintf("%d", s.Summary.Volumes), w/3-15)))

	return lipgloss.JoinVertical(lipgloss.Left, lipgloss.PlaceHorizontal(w/3-4, lipgloss.Center, styles.TitleStyle().Render(" Project Details ")), "\n\n", text)
}

func (s *ProjectDetails) GetResourcesView() string {
	w := s.Width

	text := fmt.Sprintf("%s %s \n\n%s %s \n\n%s %s \n\n%s %s \n\n%s %s \n\n%s %s \n\n%s %s",
		styles.DetailKeyStyle().Render(" CPU: "), styles.TextStyle().Render(fmt.Sprintf("%d", s.Inspect.AggregatedStats.CPU)),
		styles.DetailKeyStyle().Render(" Memory: "), styles.TextStyle().Render(fmt.Sprintf("%d", s.Inspect.AggregatedStats.Mem)),
		styles.DetailKeyStyle().Render(" Memory Limit: "), styles.TextStyle().Render(fmt.Sprintf("%d", s.Inspect.AggregatedStats.MemLimit)),
		styles.DetailKeyStyle().Render(" Net Rx: "), styles.TextStyle().Render(fmt.Sprintf("%d", s.Inspect.AggregatedStats.NetRx)),
		styles.DetailKeyStyle().Render(" Net Tx: "), styles.TextStyle().Render(fmt.Sprintf("%d", s.Inspect.AggregatedStats.NetTx)),
		styles.DetailKeyStyle().Render(" Block Read: "), styles.TextStyle().Render(fmt.Sprintf("%d", s.Inspect.AggregatedStats.BlkRead)),
		styles.DetailKeyStyle().Render(" Block Write: "), styles.TextStyle().Render(fmt.Sprintf("%d", s.Inspect.AggregatedStats.BlkWrite)))

	return lipgloss.JoinVertical(lipgloss.Left, lipgloss.PlaceHorizontal(w/3-4, lipgloss.Center, styles.TitleStyle().Render(" Project Details ")), "\n\n", text)
}

// TODO: Configure Registry
func (s *ProjectDetails) GetYAMLView() string {
	w := s.Width

	var text string
	if !s.Summary.RegistryConfigured {
		text = lipgloss.PlaceHorizontal(w/3-2, lipgloss.Center, "Registry Not Configured")
	} else {
		text = lipgloss.PlaceHorizontal(w/3-2, lipgloss.Center, "TODO: Configure Registry")
	}

	return lipgloss.JoinVertical(lipgloss.Left, lipgloss.PlaceHorizontal(w/3-4, lipgloss.Center, styles.TitleStyle().Render(" YAML Details ")), "\n\n", text)
}

func (s *ProjectDetails) GetServicesView() string {
	w := s.Width

	text := compose.ServiceHeaders(w) + "\n\n"
	for k, v := range s.Services {
		if k == s.Current {
			text += styles.SelectedStyle().Render(compose.ServiceFormatted(w, s.Inspect.Services[v]))
		} else {
			text += styles.TextStyle().Render(compose.ServiceFormatted(w, s.Inspect.Services[v]))
		}

		text += "\n"
	}

	return lipgloss.JoinVertical(lipgloss.Center, lipgloss.PlaceHorizontal(w-4, lipgloss.Center, styles.TitleStyle().Render(" Service Details ")), "\n\n", text)
}
