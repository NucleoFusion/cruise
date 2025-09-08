package nav

import (
	"github.com/NucleoFusion/cruise/internal/enums"
	"github.com/NucleoFusion/cruise/internal/keymap"
	"github.com/NucleoFusion/cruise/internal/messages"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type Nav struct {
	Width  int
	Height int
	Pages  map[string][]enums.PageType
	Keymap keymap.NavMap
}

func NewNav(w, h int) *Nav {
	pgs := map[string][]enums.PageType{
		"System":    {enums.Home},
		"Artifacts": {enums.Containers, enums.Images, enums.Networks, enums.Volumes},
		"Ops":       {enums.Vulnerability, enums.Monitoring},
	}
	return &Nav{
		Width:  w,
		Height: h,
		Pages:  pgs,
	}
}

func (s *Nav) Init() tea.Cmd { return nil }

func (s *Nav) Update(msg tea.Msg) (*Nav, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, s.Keymap.Dashboard):
			return s, func() tea.Msg { return messages.ChangePg{Pg: enums.Home, Exited: false} }
		case key.Matches(msg, s.Keymap.Dashboard):
			return s, func() tea.Msg { return messages.ChangePg{Pg: enums.Home, Exited: false} }
		case key.Matches(msg, s.Keymap.Containers):
			return s, func() tea.Msg { return messages.ChangePg{Pg: enums.Containers, Exited: false} }
		case key.Matches(msg, s.Keymap.Images):
			return s, func() tea.Msg { return messages.ChangePg{Pg: enums.Images, Exited: false} }
		case key.Matches(msg, s.Keymap.Networks):
			return s, func() tea.Msg { return messages.ChangePg{Pg: enums.Networks, Exited: false} }
		case key.Matches(msg, s.Keymap.Volumes):
			return s, func() tea.Msg { return messages.ChangePg{Pg: enums.Volumes, Exited: false} }
		case key.Matches(msg, s.Keymap.Monitoring):
			return s, func() tea.Msg { return messages.ChangePg{Pg: enums.Monitoring, Exited: false} }
		case key.Matches(msg, s.Keymap.Vulnerability):
			return s, func() tea.Msg { return messages.ChangePg{Pg: enums.Vulnerability, Exited: false} }
		}
	}
	return s, nil
}

// TODO: uild View
func View() string {
	return "Navigaaaayshon"
}
