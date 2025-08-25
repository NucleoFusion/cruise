package styles

import (
	"github.com/NucleoFusion/cruise/internal/colors"
	"github.com/charmbracelet/lipgloss"
)

func ErrorStyle() lipgloss.Style {
	return lipgloss.NewStyle().Background(colors.Load().ErrorBg).Foreground(colors.Load().ErrorText)
}

func PageStyle() lipgloss.Style {
	return lipgloss.NewStyle().Foreground(colors.Load().Text).
		Border(lipgloss.RoundedBorder()).BorderForeground(colors.Load().FocusedBorder)
}

func SubpageStyle() lipgloss.Style {
	return lipgloss.NewStyle().Foreground(colors.Load().Text).
		Border(lipgloss.RoundedBorder()).BorderForeground(colors.Load().UnfocusedBorder)
}

func TitleStyle() lipgloss.Style {
	return lipgloss.NewStyle().Background(colors.Load().SubtitleBg).Foreground(colors.Load().SubtitleText).Padding(0, 1)
}

func SelectedStyle() lipgloss.Style {
	return lipgloss.NewStyle().Background(colors.Load().MenuSelectedText).Foreground(colors.Load().MenuSelectedText)
}

func DetailKeyStyle() lipgloss.Style {
	return lipgloss.NewStyle().Foreground(colors.Load().HelpKeyText).Background(colors.Load().HelpKeyBg)
}
