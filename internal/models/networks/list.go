package networks

import (
	"sort"
	"time"

	"github.com/NucleoFusion/cruise/internal/colors"
	"github.com/NucleoFusion/cruise/internal/config"
	"github.com/NucleoFusion/cruise/internal/docker"
	"github.com/NucleoFusion/cruise/internal/messages"
	"github.com/NucleoFusion/cruise/internal/styles"
	"github.com/NucleoFusion/cruise/internal/utils"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/docker/docker/api/types/network"
	"github.com/lithammer/fuzzysearch/fuzzy"
)

type NetworkList struct {
	Width         int
	Height        int
	Items         []network.Summary
	FilteredItems []network.Summary
	SelectedIndex int
	Ti            textinput.Model
	Vp            viewport.Model
}

func NewNetworkList(w int, h int) *NetworkList {
	ti := textinput.New()
	ti.Width = w - 12
	ti.Prompt = " Search: "
	ti.Placeholder = "Press '/' to search..."

	ti.PromptStyle = lipgloss.NewStyle().Foreground(colors.Load().FocusedBorder)
	ti.PlaceholderStyle = lipgloss.NewStyle().Foreground(colors.Load().PlaceholderText)
	ti.TextStyle = styles.TextStyle()

	vp := viewport.New(w, h-3)
	vp.Style = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(colors.Load().FocusedBorder).
		Padding(1).Foreground(colors.Load().Text)

	return &NetworkList{
		Width:         w,
		Height:        h,
		Ti:            ti,
		SelectedIndex: 0,
		Vp:            vp,
	}
}

func (s *NetworkList) Init() tea.Cmd {
	return tea.Tick(0, func(_ time.Time) tea.Msg {
		images, err := docker.GetNetworks()
		if err != nil {
			return utils.ReturnError("Networks Page", "Error Querying Networks", err)
		}
		return messages.NetworksReadyMsg{Items: images}
	})
}

func (s *NetworkList) Update(msg tea.Msg) (*NetworkList, tea.Cmd) {
	switch msg := msg.(type) {
	case messages.NetworksReadyMsg:
		s.Items = msg.Items
		s.FilteredItems = msg.Items
		return s, tea.Tick(3*time.Second, func(_ time.Time) tea.Msg {
			images, err := docker.GetNetworks()
			if err != nil {
				return utils.ReturnError("Networks Page", "Error Querying Networks", err)
			}
			return messages.NetworksReadyMsg{Items: images}
		})

	case tea.KeyMsg:
		if s.Ti.Focused() {
			if msg.String() == config.Cfg.Keybinds.Global.UnfocusSearch {
				s.Ti.Blur()
				return s, nil
			}
			var cmd tea.Cmd
			s.Ti, cmd = s.Ti.Update(msg)
			s.Filter(s.Ti.Value())
			s.UpdateList()
			return s, cmd
		}
		switch msg.String() {
		case config.Cfg.Keybinds.Global.FocusSearch:
			s.Ti.Focus()
			return s, nil
		case config.Cfg.Keybinds.Global.ListDown:
			if len(s.FilteredItems)-1 > s.SelectedIndex {
				s.SelectedIndex += 1
			}
			if s.SelectedIndex > s.Vp.Height+s.Vp.YOffset-7 { // -2 for border and sosething else, idk breaks otherwise
				s.Vp.YOffset += 1
			}
			s.UpdateList()
			return s, nil
		case config.Cfg.Keybinds.Global.ListUp:
			if 0 < s.SelectedIndex {
				s.SelectedIndex -= 1
			}
			if s.SelectedIndex < s.Vp.YOffset {
				s.Vp.YOffset -= 1
			}
			s.UpdateList()
			return s, nil
		}
	}
	return s, nil
}

func (s *NetworkList) View() string {
	if len(s.Items) == 0 {
		return lipgloss.Place(s.Width-2, s.Height, lipgloss.Center, lipgloss.Center, "No Containers Found!")
	}

	style := lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(colors.Load().FocusedBorder)

	s.UpdateList()

	return lipgloss.JoinVertical(lipgloss.Center,
		style.Render(s.Ti.View()),
		s.Vp.View())
}

func (s *NetworkList) UpdateList() {
	text := lipgloss.NewStyle().Bold(true).Render(docker.NetworksHeaders(s.Width-2)+"\n") + "\n"

	for k, v := range s.FilteredItems {
		line := docker.NetworksFormattedSummary(v, s.Width-2)

		if k == s.SelectedIndex {
			line = lipgloss.NewStyle().Background(colors.Load().MenuSelectedBg).Foreground(colors.Load().MenuSelectedText).Render(line)
		} else {
			line = styles.TextStyle().Render(line)
		}

		text += line + "\n"
	}

	s.Vp.SetContent(text)
}

func (s *NetworkList) Filter(val string) {
	formatted := make([]string, len(s.Items))
	originals := make([]network.Summary, len(s.Items))

	for i, v := range s.Items {
		str := docker.NetworksFormattedSummary(v, s.Width-2)
		formatted[i] = str
		originals[i] = v
	}

	ranked := fuzzy.RankFindFold(val, formatted)
	sort.Sort(ranked)

	result := make([]network.Summary, len(ranked))
	for i, r := range ranked {
		result[i] = originals[r.OriginalIndex]
	}

	s.FilteredItems = result

	if len(s.FilteredItems) <= s.SelectedIndex {
		s.SelectedIndex = len(s.FilteredItems) - 1
	}
}

func (s *NetworkList) GetCurrentItem() network.Summary {
	return s.FilteredItems[s.SelectedIndex]
}
