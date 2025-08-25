package volumes

import (
	"fmt"
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
	"github.com/docker/docker/api/types/volume"
	"github.com/lithammer/fuzzysearch/fuzzy"
)

type VolumeList struct {
	Width         int
	Height        int
	Items         []*volume.Volume
	FilteredItems []*volume.Volume
	SelectedIndex int
	Ti            textinput.Model
	Vp            viewport.Model
}

func NewVolumeList(w int, h int) *VolumeList {
	ti := textinput.New()
	ti.Width = w - 9
	ti.Prompt = " Search: "
	ti.Placeholder = "Press '/' to search..."

	ti.PromptStyle = lipgloss.NewStyle().Foreground(colors.Load().FocusedBorder)
	ti.PlaceholderStyle = lipgloss.NewStyle().Foreground(colors.Load().PlaceholderText)
	ti.TextStyle = styles.TextStyle()

	vp := viewport.New(w+3, h+1)
	vp.Style = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(colors.Load().FocusedBorder).
		Padding(1).Foreground(colors.Load().Text)

	return &VolumeList{
		Width:         w,
		Height:        h,
		Ti:            ti,
		SelectedIndex: 0,
		Vp:            vp,
	}
}

func (s *VolumeList) Init() tea.Cmd {
	return tea.Tick(0, func(_ time.Time) tea.Msg {
		vols, err := docker.GetVolumes()
		if err != nil {
			fmt.Println(err)
			return utils.ReturnError("Volumes Page", "Error Querying Volumes", err)
		}
		return messages.VolumesReadyMsg{Items: vols.Volumes}
	})
}

func (s *VolumeList) Update(msg tea.Msg) (*VolumeList, tea.Cmd) {
	switch msg := msg.(type) {
	case messages.VolumesReadyMsg:
		s.Items = msg.Items
		s.FilteredItems = msg.Items
		return s, nil
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

func (s *VolumeList) View() string {
	if len(s.Items) == 0 {
		return lipgloss.Place(s.Width, s.Height, lipgloss.Center, lipgloss.Center, "No Volumes Found!")
	}

	style := lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(colors.Load().FocusedBorder)

	s.UpdateList()

	return lipgloss.JoinVertical(lipgloss.Center,
		style.Render(s.Ti.View()),
		s.Vp.View())
}

func (s *VolumeList) UpdateList() {
	text := lipgloss.NewStyle().Bold(true).Render(docker.VolumesHeaders(s.Width)+"\n") + "\n"

	for k, v := range s.FilteredItems {
		if v == nil {
			continue
		}

		line := docker.VolumesFormattedSummary(*v, s.Width)

		if k == s.SelectedIndex {
			line = lipgloss.NewStyle().Background(colors.Load().MenuSelectedText).Foreground(colors.Load().MenuSelectedBg).Render(line)
		} else {
			line = styles.TextStyle().Render(line)
		}

		text += line + "\n"
	}

	s.Vp.SetContent(text)
}

func (s *VolumeList) Filter(val string) {
	formatted := make([]string, len(s.Items))
	originals := make([]*volume.Volume, len(s.Items))

	for i, v := range s.Items {
		str := docker.VolumesFormattedSummary(*v, s.Width)
		formatted[i] = str
		originals[i] = v
	}

	ranked := fuzzy.RankFindFold(val, formatted)
	sort.Sort(ranked)

	result := make([]*volume.Volume, len(ranked))
	for i, r := range ranked {
		result[i] = originals[r.OriginalIndex]
	}

	s.FilteredItems = result

	if len(s.FilteredItems) <= s.SelectedIndex {
		s.SelectedIndex = len(s.FilteredItems) - 1
	}
}

func (s *VolumeList) GetCurrentItem() volume.Volume {
	return *s.FilteredItems[s.SelectedIndex]
}
