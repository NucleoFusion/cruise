package page

import tea "github.com/charmbracelet/bubbletea"

type Page interface {
	Init() tea.Cmd
	Update(msg tea.Msg) (tea.Model, tea.Cmd)
	View() string

	// Called when page needs to clean up routines and more
	// mostly on page switch
	// Cleanup() tea.Cmd
}
