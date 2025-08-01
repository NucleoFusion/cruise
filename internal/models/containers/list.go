package containers

import (
	"sort"
	"time"

	"github.com/NucleoFusion/cruise/internal/colors"
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
	ti.Width = w - 9
	ti.Prompt = " Search: "
	ti.Placeholder = "Press '/' to search..."

	ti.PromptStyle = lipgloss.NewStyle().Foreground(colors.Load().Lavender)
	ti.PlaceholderStyle = lipgloss.NewStyle().Foreground(colors.Load().Surface2)
	ti.TextStyle = styles.TextStyle()

	vp := viewport.New(w+3, h+1)
	vp.Style = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(colors.Load().Lavender).
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
		return s, nil
	case tea.KeyMsg:
		if s.Ti.Focused() {
			if msg.String() == "esc" {
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
		case "/":
			s.Ti.Focus()
			return s, nil
		case "down":
			if len(s.FilteredItems)-1 > s.SelectedIndex {
				s.SelectedIndex += 1
			}
			if s.SelectedIndex > s.Vp.Height+s.Vp.YOffset-3 { // -2 for border and sosething else, idk breaks otherwise
				s.Vp.YOffset += 1
			}
			s.UpdateList()
			return s, nil
		case "up":
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
		return lipgloss.Place(s.Width, s.Height, lipgloss.Center, lipgloss.Center, "Error: "+s.Err.Error())
	}

	if len(s.Items) == 0 {
		return lipgloss.Place(s.Width, s.Height, lipgloss.Center, lipgloss.Center, "No Containers Found!")
	}

	style := lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(colors.Load().Lavender)

	s.UpdateList()

	return lipgloss.JoinVertical(lipgloss.Center,
		style.Render(s.Ti.View()),
		s.Vp.View())
}

func (s *ContainerList) UpdateList() {
	w := (s.Width)/9 - 1

	text := lipgloss.NewStyle().Bold(true).Render(docker.SummaryHeaders(w)+"\n") + "\n"

	for k, v := range s.FilteredItems {
		line := docker.FormattedSummary(v, w)

		if k == s.SelectedIndex {
			line = lipgloss.NewStyle().Background(colors.Load().Lavender).Foreground(colors.Load().Base).Render(line)
		} else {
			line = styles.TextStyle().Render(line)
		}

		text += line + "\n"
	}

	s.Vp.SetContent(text)
}

func (s *ContainerList) Filter(val string) {
	w := (s.Width)/9 - 1

	formatted := make([]string, len(s.Items))
	originals := make([]container.Summary, len(s.Items))

	for i, v := range s.Items {
		str := docker.FormattedSummary(v, w)
		formatted[i] = str
		originals[i] = v
	}

	ranked := fuzzy.RankFindFold(val, formatted)
	sort.Sort(ranked) // Best matches first

	result := make([]container.Summary, len(ranked))
	for i, r := range ranked {
		result[i] = originals[r.OriginalIndex]
	}

	s.FilteredItems = result // <- You’ll need to define this as []container.Summary

	if len(s.FilteredItems) <= s.SelectedIndex {
		s.SelectedIndex = len(s.FilteredItems) - 1
	}
}
