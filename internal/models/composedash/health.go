package composedash

import (
	"fmt"

	"github.com/NucleoFusion/cruise/internal/styles"
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/lipgloss"
)

type HealthDetails struct {
	Width      int
	Projects   int
	Services   int
	Containers int
	Volumes    int
	Networks   int
}

func NewHealth(w int, h int) viewport.Model {
	vp := viewport.New(w, h)
	vp.Style = styles.PageStyle().Padding(1, 3)

	vp.SetContent("Loading Health Details...")

	return vp
}

func (s *HealthDetails) View() string {
	title := lipgloss.PlaceHorizontal(s.Width-5, lipgloss.Center, styles.TitleStyle().Render("Overview"))

	dets := fmt.Sprintf("%s: %d\n\n%s: %d\n\n%s: %d\n\n%s: %d\n\n%s: %d",
		styles.DetailKeyStyle().Render("Projects"), s.Projects,
		styles.DetailKeyStyle().Render("Services"), s.Services,
		styles.DetailKeyStyle().Render("Containers"), s.Containers,
		styles.DetailKeyStyle().Render("Volumes"), s.Volumes,
		styles.DetailKeyStyle().Render("Networks"), s.Networks,
	)

	return lipgloss.JoinVertical(lipgloss.Left, title, "\n\n", dets)
}
