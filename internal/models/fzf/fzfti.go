package fzf

import (
	"github.com/NucleoFusion/cruise/internal/colors"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
)

func TextStyle() lipgloss.Style {
	return lipgloss.NewStyle().Foreground(colors.Load().Text)
}

func NewTI(w int) textinput.Model {
	ti := textinput.New()
	ti.Placeholder = "Search..."
	ti.Prompt = "Filter : "
	ti.Width = w - 12
	ti.TextStyle = TextStyle()
	ti.PromptStyle = TextStyle()
	ti.Focus()

	return ti
}
