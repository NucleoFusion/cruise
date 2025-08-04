package containers

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
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
	Width            int
	Height           int
	Items            []container.Summary
	Err              error
	FilteredItems    []container.Summary
	SelectedIndex    int
	IsExpanded       bool
	ExpandedIndex    int
	IsDetailsLoading bool
	StatsReader      container.StatsResponseReader
	Stats            container.StatsResponse
	Decoder          *json.Decoder
	Ti               textinput.Model
	Vp               viewport.Model
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
		Width:            w,
		Height:           h,
		Ti:               ti,
		SelectedIndex:    0,
		IsExpanded:       false,
		ExpandedIndex:    0,
		Vp:               vp,
		IsDetailsLoading: true,
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
		return s, tea.Tick(2*time.Second, func(_ time.Time) tea.Msg {
			items, err := docker.GetContainers()
			return messages.ContainerReadyMsg{
				Items: items,
				Err:   err,
			}
		})
	case messages.NewContainerDetails:
		s.IsDetailsLoading = false
		s.StatsReader = msg.Stats
		s.Decoder = msg.Decoder

		var m container.StatsResponse
		s.Decoder.Decode(&m)
		s.Stats = m

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
		case "enter":
			if s.IsExpanded {
				if s.SelectedIndex == s.ExpandedIndex {
					s.IsExpanded = false
					return s, nil
				}

				s.IsDetailsLoading = true

				s.ExpandedIndex = s.SelectedIndex
				return s, s.NewStats()
			}

			s.IsDetailsLoading = true

			s.ExpandedIndex = s.SelectedIndex
			s.IsExpanded = true
			return s, s.NewStats()
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

		if s.IsExpanded && k == s.ExpandedIndex {
			dropdown := ""
			if s.IsDetailsLoading {
				dropdown += lipgloss.Place(s.Width-14, 6, lipgloss.Center, lipgloss.Center, "Loading")
			} else {
				var totalRxBytes, totalTxBytes uint64

				for _, netStats := range s.Stats.Networks {
					totalRxBytes += netStats.RxBytes
					totalTxBytes += netStats.TxBytes
				}

				dropdown += fmt.Sprintf(" %s | [CPU%%: %d | Mem: %d | Network: %.2fMB Rx / %.2fMB Tx ]\n",
					strings.Trim(s.Stats.Name, "/"), s.Stats.CPUStats.CPUUsage.TotalUsage, s.Stats.MemoryStats.Usage, float64(totalRxBytes)/(1024*1024), float64(totalTxBytes)/(1024*1024))

				dropdown += strings.Repeat("─", s.Width-14)

				dropdown += "\n\n\n"
			}

			line += lipgloss.NewStyle().MarginLeft(3).Border(styles.DropdownBorder()).Render(dropdown)
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

func (s *ContainerList) NewStats() tea.Cmd {
	return tea.Tick(0, func(_ time.Time) tea.Msg {
		stats, err := docker.GetContainerStats(s.GetCurrentItem().ID)
		if err != nil {
			return messages.ErrorMsg{
				Title: "Error Querying Stats",
				Locn:  "Container Page",
				Msg:   err.Error(),
			}
		}
		decoder := json.NewDecoder(stats.Body)

		return messages.NewContainerDetails{
			Stats:   stats,
			Decoder: decoder,
		}
	})
}
