package fzf

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
	"github.com/cruise-org/cruise/pkg/colors"
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
