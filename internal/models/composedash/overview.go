package composedash

import (
	"fmt"

	"github.com/NucleoFusion/cruise/internal/styles"
	"github.com/charmbracelet/bubbles/viewport"
)

type OverviewDetails struct {
	Projects   int
	Services   int
	Containers int
	Volumes    int
	Networks   int
}

func NewOverview(w int, h int) viewport.Model {
	vp := viewport.New(w, h)
	vp.Style = styles.PageStyle()

	vp.SetContent("Loading Overview Details...")

	return vp
}

func (s *OverviewDetails) View() string {
	return fmt.Sprintf("%s: %d\n%s: %d\n%s: %d\n%s: %d\n%s: %d",
		styles.DetailKeyStyle().Render("Projects"), s.Projects,
		styles.DetailKeyStyle().Render("Services"), s.Services,
		styles.DetailKeyStyle().Render("Containers"), s.Containers,
		styles.DetailKeyStyle().Render("Volumes"), s.Volumes,
		styles.DetailKeyStyle().Render("Networks"), s.Networks,
	)
}
