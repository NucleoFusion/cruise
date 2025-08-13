package styles

import (
	"github.com/NucleoFusion/cruise/internal/colors"
	"github.com/charmbracelet/lipgloss"
)

func ErrorStyle() lipgloss.Style {
	return lipgloss.NewStyle().Background(colors.Load().Red).Foreground(colors.Load().Base)
}

func PageStyle() lipgloss.Style {
	return lipgloss.NewStyle().Background(colors.Load().Base).Foreground(colors.Load().Text).
		Border(lipgloss.RoundedBorder()).BorderForeground(colors.Load().Lavender)
}

func SubpageStyle() lipgloss.Style {
	return lipgloss.NewStyle().Background(colors.Load().Base).Foreground(colors.Load().Text).
		Border(lipgloss.RoundedBorder()).BorderForeground(colors.Load().Surface1)
}

func TitleStyle() lipgloss.Style {
	return lipgloss.NewStyle().Foreground(colors.Load().Sapphire)
}
