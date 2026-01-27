package images

import (
	"sort"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/cruise-org/cruise/internal/messages"
	"github.com/cruise-org/cruise/internal/runtimes/docker"
	"github.com/cruise-org/cruise/internal/utils"
	"github.com/cruise-org/cruise/pkg/colors"
	"github.com/cruise-org/cruise/pkg/config"
	"github.com/cruise-org/cruise/pkg/styles"
	"github.com/docker/docker/api/types/image"
	"github.com/lithammer/fuzzysearch/fuzzy"
)

type ImageList struct {
	Width         int
	Height        int
	ImageMap      map[string]image.Summary
	Items         []string
	FilteredItems []string
	SelectedIndex int
	Ti            textinput.Model
	Vp            viewport.Model
}

func NewImageList(w int, h int) *ImageList {
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

	return &ImageList{
		Width:         w,
		Height:        h,
		Ti:            ti,
		SelectedIndex: 0,
		Vp:            vp,
		ImageMap:      make(map[string]image.Summary),
	}
}

func (s *ImageList) Init() tea.Cmd {
	return tea.Tick(0, func(_ time.Time) tea.Msg {
		images, err := docker.GetImages()

		m := make(map[string]image.Summary)
		for _, v := range images {
			m[v.ID] = v
		}

		if err != nil {
			return utils.ReturnError("Images Page", "Error Querying Images", err)
		}
		return messages.ImagesReadyMsg{Map: m}
	})
}

func (s *ImageList) Update(msg tea.Msg) (*ImageList, tea.Cmd) {
	switch msg := msg.(type) {
	case messages.ImagesReadyMsg:
		s.ImageMap = msg.Map

		items := make([]string, 0, len(msg.Map))
		for k := range s.ImageMap {
			items = append(items, k)
		}

		sort.Strings(items)

		s.Items = items
		s.FilteredItems = items

		return s, nil

	case messages.UpdateImagesMsg:
		items := make([]string, 0, len(msg.Items))
		for _, v := range msg.Items {
			s.ImageMap[v.ID] = v
			items = append(items, v.ID)
		}

		sort.Strings(items)

		s.Items = items
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

func (s *ImageList) View() string {
	if len(s.Items) == 0 {
		return lipgloss.Place(s.Width-2, s.Height, lipgloss.Center, lipgloss.Center, "No Containers Found!")
	}

	style := lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(colors.Load().FocusedBorder)

	s.UpdateList()

	return lipgloss.JoinVertical(lipgloss.Center,
		style.Render(s.Ti.View()),
		s.Vp.View())
}

func (s *ImageList) UpdateList() {
	w := (s.Width-2)/9 - 1

	text := lipgloss.NewStyle().Bold(true).Render(docker.ImagesHeaders(w)+"\n") + "\n"

	for k, v := range s.FilteredItems {
		line := docker.ImagesFormattedSummary(s.ImageMap[v], w)

		if k == s.SelectedIndex {
			line = lipgloss.NewStyle().Background(colors.Load().MenuSelectedBg).Foreground(colors.Load().MenuSelectedText).Render(line)
		} else {
			line = styles.TextStyle().Render(line)
		}

		text += line + "\n"
	}

	s.Vp.SetContent(text)
}

func (s *ImageList) Filter(val string) {
	w := (s.Width-2)/9 - 1

	formatted := make([]string, len(s.Items))
	originals := make([]image.Summary, len(s.Items))

	for i, v := range s.Items {
		str := docker.ImagesFormattedSummary(s.ImageMap[v], w)
		formatted[i] = str
		originals[i] = s.ImageMap[v]
	}

	ranked := fuzzy.RankFindFold(val, formatted)
	sort.Sort(ranked)

	result := make([]string, len(ranked))
	for i, r := range ranked {
		result[i] = originals[r.OriginalIndex].ID
	}

	s.FilteredItems = result

	if len(s.FilteredItems) <= s.SelectedIndex {
		s.SelectedIndex = len(s.FilteredItems) - 1
	}
}

func (s *ImageList) GetCurrentItem() image.Summary {
	return s.ImageMap[s.FilteredItems[s.SelectedIndex]]
}
