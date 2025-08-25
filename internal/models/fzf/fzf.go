package fzf

import (
	"sort"

	"github.com/NucleoFusion/cruise/internal/colors"
	"github.com/NucleoFusion/cruise/internal/keymap"
	"github.com/NucleoFusion/cruise/internal/messages"
	styledhelp "github.com/NucleoFusion/cruise/internal/models/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/lithammer/fuzzysearch/fuzzy"
)

type FuzzyFinder struct {
	Width         int
	Height        int
	Keymap        keymap.FuzzyMap
	Help          styledhelp.StyledHelp
	Vp            viewport.Model
	Ti            textinput.Model
	Items         []string
	Filtered      []string
	SelectedIndex int
}

func NewFzf(items []string, w int, h int) FuzzyFinder {
	return FuzzyFinder{
		Width:         w,
		Height:        h,
		Help:          styledhelp.NewStyledHelp(keymap.NewFuzzyMap().Bindings(), w),
		Keymap:        keymap.NewFuzzyMap(),
		Items:         items,
		Filtered:      items,
		Vp:            NewVP(w, h, items),
		Ti:            NewTI(w / 3),
		SelectedIndex: 0,
	}
}

func (m *FuzzyFinder) Init() tea.Cmd {
	return nil
}

func (m *FuzzyFinder) Update(msg tea.Msg) (FuzzyFinder, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.Keymap.StartWriting):
			m.Ti.Placeholder = "Search..."
			m.Ti.Focus()
		case key.Matches(msg, m.Keymap.Down):
			if len(m.Filtered)-1 > m.SelectedIndex {
				m.SelectedIndex += 1
			}
			if m.SelectedIndex > m.Vp.Height+m.Vp.YOffset-3 { // -2 for border and something else, idk breaks otherwise
				m.Vp.YOffset += 1
			}
			m.UpdateVP()
			return *m, nil
		case key.Matches(msg, m.Keymap.Up):
			if 0 < m.SelectedIndex {
				m.SelectedIndex -= 1
			}
			if m.SelectedIndex < m.Vp.YOffset {
				m.Vp.YOffset -= 1
			}
			m.UpdateVP()
			return *m, nil
		case key.Matches(msg, m.Keymap.Enter):
			return *m, func() tea.Msg {
				return messages.FzfSelection{
					Selection: m.Filtered[m.SelectedIndex],
					Exited:    false,
				}
			}
		case key.Matches(msg, m.Keymap.Exit):
			return *m, func() tea.Msg {
				return messages.FzfSelection{
					Selection: "",
					Exited:    true,
				}
			}
		default:
			if m.Ti.Focused() {
				var cmd tea.Cmd
				m.Ti, cmd = m.Ti.Update(msg)
				m.Filter(m.Ti.Value())
				m.UpdateVP()
				return *m, cmd
			}
		}
	}

	return *m, nil
}

func (m *FuzzyFinder) View() string {
	pg := lipgloss.Place(m.Width, m.Height-1, lipgloss.Center, lipgloss.Center,
		lipgloss.JoinVertical(lipgloss.Center,
			lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(colors.Load().FocusedBorder).Render(m.Ti.View()),
			m.Vp.View()))
	hlp := lipgloss.NewStyle().PaddingLeft(2).Render(m.Help.View())
	return lipgloss.JoinVertical(lipgloss.Left, pg, hlp)
}

func (m *FuzzyFinder) UpdateVP() {
	text := ""
	for k, v := range m.Filtered {
		if k == m.SelectedIndex {
			text += SelectedItemStyle(m.Width/3).Render(v) + "\n"
			continue
		}

		text += ItemLineStyle(m.Width/3).Render(v) + "\n"
	}

	m.Vp.SetContent(text)
}

func (m *FuzzyFinder) Filter(val string) {
	ranked := fuzzy.RankFindFold(val, m.Items)
	sort.Sort(ranked) // Sort by ascending distance
	result := make([]string, len(ranked))
	for i, r := range ranked {
		result[i] = r.Target
	}

	m.Filtered = result

	if len(m.Filtered) <= m.SelectedIndex { // So that index doesnt go out of bounds
		m.SelectedIndex = len(m.Filtered) - 1
	}
}
