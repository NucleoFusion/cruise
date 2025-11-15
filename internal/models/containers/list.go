package containers

import (
	"context"
	"sort"
	"time"

	"github.com/NucleoFusion/cruise/internal/colors"
	"github.com/NucleoFusion/cruise/internal/config"
	"github.com/NucleoFusion/cruise/internal/docker"
	"github.com/NucleoFusion/cruise/internal/messages"
	"github.com/NucleoFusion/cruise/internal/styles"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/docker/docker/api/types/container"
	"github.com/lithammer/fuzzysearch/fuzzy"
)

type LogStreamer struct {
	ctx    context.Context
	cancel context.CancelFunc
	lines  chan string
}

type ContainerList struct {
	Width         int
	Height        int
	Items         []container.Summary
	Err           error
	FilteredItems []container.Summary
	SelectedIndex int
	Ti            textinput.Model
	Vp            viewport.Model
}

func NewContainerList(w int, h int) *ContainerList {
	ti := textinput.New()
	ti.Width = w - 12
	ti.Prompt = " Search: "
	ti.Placeholder = "Press '/' to search..."

	ti.PromptStyle = lipgloss.NewStyle().Foreground(colors.Load().FocusedBorder)
	ti.PlaceholderStyle = lipgloss.NewStyle().Foreground(colors.Load().PlaceholderText)
	ti.TextStyle = styles.TextStyle()

	vp := viewport.New(w, h-4) //h-4 to account for searchbar
	vp.Style = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(colors.Load().FocusedBorder).
		Padding(1).Foreground(colors.Load().Text)

	return &ContainerList{
		Width:         w,
		Height:        h,
		Ti:            ti,
		SelectedIndex: 0,
		Vp:            vp,
	}
}

func (s *ContainerList) Init() tea.Cmd {
	return tea.Tick(0, func(_ time.Time) tea.Msg {
		items, err := docker.GetContainers()
		return messages.ContainerReadyMsg{
			Items: items,
			Err:   err,
		}
	})
}

func (s *ContainerList) Update(msg tea.Msg) (*ContainerList, tea.Cmd) {
	switch msg := msg.(type) {
	case messages.ContainerReadyMsg:
		s.Items = msg.Items
		s.FilteredItems = msg.Items
		s.Err = msg.Err
		return s, tea.Tick(3*time.Second, func(_ time.Time) tea.Msg {
			items, err := docker.GetContainers()
			return messages.ContainerReadyMsg{
				Items: items,
				Err:   err,
			}
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
			if s.SelectedIndex > s.Vp.Height+s.Vp.YOffset-3 { // -2 for border and sosething else, idk breaks otherwise
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

func (s *ContainerList) View() string {
	if s.Err != nil {
		return styles.PageStyle().Render(lipgloss.Place(s.Width-2, s.Height, lipgloss.Center, lipgloss.Center, "Error: "+s.Err.Error()))
	}

	if len(s.Items) == 0 {
		return styles.PageStyle().Render(lipgloss.Place(s.Width-2, s.Height, lipgloss.Center, lipgloss.Center, "No Containers Found!"))
	}

	style := lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(colors.Load().FocusedBorder)

	s.UpdateList()

	return lipgloss.JoinVertical(lipgloss.Center,
		style.Render(s.Ti.View()),
		s.Vp.View())
}

func (s *ContainerList) UpdateList() {
	w := (s.Width-2)/9 - 1

	text := lipgloss.NewStyle().Bold(true).Render(docker.ContainerHeaders(w)+"\n") + "\n"

	for k, v := range s.FilteredItems {
		line := docker.ContainerFormattedSummary(v, w)

		if k == s.SelectedIndex {
			line = styles.SelectedStyle().Render(line)
		} else {
			line = styles.TextStyle().Render(line)
		}

		text += line + "\n"
	}

	s.Vp.SetContent(text)
}

func (s *ContainerList) Filter(val string) {
	w := (s.Width-2)/9 - 1

	formatted := make([]string, len(s.Items))
	originals := make([]container.Summary, len(s.Items))

	for i, v := range s.Items {
		str := docker.ContainerFormattedSummary(v, w)
		formatted[i] = str
		originals[i] = v
	}

	ranked := fuzzy.RankFindFold(val, formatted)
	sort.Sort(ranked)

	result := make([]container.Summary, len(ranked))
	for i, r := range ranked {
		result[i] = originals[r.OriginalIndex]
	}

	s.FilteredItems = result

	if len(s.FilteredItems) <= s.SelectedIndex {
		s.SelectedIndex = len(s.FilteredItems) - 1
	}
}

func (s *ContainerList) GetCurrentItem() container.Summary {
	return s.FilteredItems[s.SelectedIndex]
}
