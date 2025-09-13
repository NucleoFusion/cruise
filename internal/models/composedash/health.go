package composedash

import (
	"github.com/NucleoFusion/cruise/internal/styles"
	"github.com/charmbracelet/bubbles/viewport"
)

func NewHealth(w int, h int) viewport.Model {
	vp := viewport.New(w, h)
	vp.Style = styles.PageStyle()

	vp.SetContent("Loading Compose Health...")

	return vp
}

func HealthView() string {
	return "Health View"
}
