package root

import (
	"github.com/NucleoFusion/cruise/internal/enums"
	tea "github.com/charmbracelet/bubbletea"
)

type Root struct {
	Width       int
	Height      int
	CurrentPage enums.PageType
}

func NewRoot() Root { return Root{} }

func (s Root) Init() tea.Cmd { return nil }

func (s Root) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return s, tea.Quit
		}
	}

	return s, nil
}

func (s Root) View() string { return "This is cruise" }
