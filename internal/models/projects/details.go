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
	IsLoading bool
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

	return &ProjectDetails{
		Width:      w,
		Height:     h,
		Summary:    s,
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

		s.DetailsVp.SetContent(s.GetDetailsView())

		s.IsLoading = false
		return s, nil
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
		styles.DetailKeyStyle().Render(" Summary: "), styles.TextStyle().Render(utils.Shorten(fmt.Sprintf("%d", s.Summary.NumServices()), w/3-15)),
		styles.DetailKeyStyle().Render(" Networks: "), styles.TextStyle().Render(utils.Shorten(fmt.Sprintf("%d", s.Summary.Networks), w/3-15)),
		styles.DetailKeyStyle().Render(" Volumes: "), styles.TextStyle().Render(utils.Shorten(fmt.Sprintf("%d", s.Summary.Volumes), w/3-15)))

	return lipgloss.JoinVertical(lipgloss.Left, lipgloss.PlaceHorizontal(w/3-4, lipgloss.Center, styles.TitleStyle().Render(" Volume Details ")), "\n\n", text)
}
