package composedash

import (
	"github.com/NucleoFusion/cruise/internal/styles"
	"github.com/charmbracelet/bubbles/viewport"
)

func NewTrends(w int, h int) viewport.Model {
	vp := viewport.New(w, h)
	vp.Style = styles.PageStyle()

	vp.SetContent("Loading Trends data...")
	return vp
}

func TrendsView() string {
	return "Trends View"
}
