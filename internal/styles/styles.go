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
	return lipgloss.NewStyle().Background(colors.Load().Surface0).Foreground(colors.Load().Sapphire).Padding(0, 1)
}

func SelectedStyle() lipgloss.Style {
	return lipgloss.NewStyle().Background(colors.Load().Lavender).Foreground(colors.Load().Base)
}

func DetailKeyStyle() lipgloss.Style {
	return lipgloss.NewStyle().Foreground(colors.Load().Text).Background(colors.Load().Surface0)
}
