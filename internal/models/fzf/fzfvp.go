package fzf

import (
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/lipgloss"
	"github.com/cruise-org/cruise/pkg/colors"
)

func ItemLineStyle(w int) lipgloss.Style {
	return lipgloss.NewStyle().Width(w).Foreground(colors.Load().Text)
}

func SelectedItemStyle(w int) lipgloss.Style {
	return ItemLineStyle(w).Background(colors.Load().MenuSelectedBg).Foreground(colors.Load().MenuSelectedText)
}

func VPStyle() lipgloss.Style {
	return lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(colors.Load().FocusedBorder).
		Foreground(colors.Load().Text)
}

func NewVP(w int, h int, items []string) viewport.Model {
	vp := viewport.New(w/3, h/2)
	vp.Style = VPStyle()

	text := ""
	for k, v := range items {
		if k == 0 {
			text += SelectedItemStyle(w/3).Render(v) + "\n"
			continue
		}

		text += ItemLineStyle(w/3).Render(v) + "\n"
	}

	vp.SetContent(text)

	return vp
}
