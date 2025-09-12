package composedash

import tea "github.com/charmbracelet/bubbletea"

type Dash struct {
	Width  int
	Height int
}

func NewComposeDash(w, h int) *Dash {
	return &Dash{
		Width:  w,
		Height: h,
	}
}

func (s *Dash) Init() tea.Cmd { return nil }

func (s *Dash) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return s, nil
}

func (s *Dash) View() string {
	return "daaashbored"
}
