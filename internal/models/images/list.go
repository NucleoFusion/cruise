package images

import (
	"fmt"
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
	"github.com/docker/docker/api/types/image"
	"github.com/lithammer/fuzzysearch/fuzzy"
)

type ImageList struct {
	Width         int
	Height        int
	Items         []image.Summary
	FilteredItems []image.Summary
	SelectedIndex int
	Ti            textinput.Model
	Vp            viewport.Model
}

func NewImageList(w int, h int) *ImageList {
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

	return &ImageList{
		Width:         w,
		Height:        h,
		Ti:            ti,
		SelectedIndex: 0,
		Vp:            vp,
	}
}

func (s *ImageList) Init() tea.Cmd {
	return tea.Tick(0, func(_ time.Time) tea.Msg {
		images, err := docker.GetImages()
		if err != nil {
			return messages.ErrorMsg{
				Locn:  "Images Page",
				Msg:   err.Error(),
				Title: "Error Querying Images",
			}
		}
		fmt.Println("Returning ImagesReady")
		return messages.ImagesReadyMsg{Items: images}
	})
}

func (s *ImageList) Update(msg tea.Msg) (*ImageList, tea.Cmd) {
	switch msg := msg.(type) {
	case messages.ImagesReadyMsg:
		s.Items = msg.Items
		s.FilteredItems = msg.Items
		return s, nil // TODO: Return tick
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
			if s.SelectedIndex > s.Vp.Height+s.Vp.YOffset-7 { // -2 for border and sosething else, idk breaks otherwise
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
			return s, nil
		}
	}
	return s, nil
}

func (s *ImageList) View() string {
	if len(s.Items) == 0 {
		return lipgloss.Place(s.Width, s.Height, lipgloss.Center, lipgloss.Center, "No Containers Found!")
	}

	style := lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(colors.Load().Lavender)

	s.UpdateList()

	return lipgloss.JoinVertical(lipgloss.Center,
		style.Render(s.Ti.View()),
		s.Vp.View())
}

func (s *ImageList) UpdateList() {
	w := (s.Width)/5 - 1

	text := lipgloss.NewStyle().Bold(true).Render(docker.ImagesHeaders(w)+"\n") + "\n"

	for k, v := range s.FilteredItems {
		line := docker.ImagesFormattedSummary(v, w)

		if k == s.SelectedIndex {
			line = lipgloss.NewStyle().Background(colors.Load().Lavender).Foreground(colors.Load().Base).Render(line)
		} else {
			line = styles.TextStyle().Render(line)
		}

		text += line + "\n"
	}

	s.Vp.SetContent(text)
}

func (s *ImageList) Filter(val string) {
	w := (s.Width)/9 - 1

	formatted := make([]string, len(s.Items))
	originals := make([]image.Summary, len(s.Items))

	for i, v := range s.Items {
		str := docker.ImagesFormattedSummary(v, w)
		formatted[i] = str
		originals[i] = v
	}

	ranked := fuzzy.RankFindFold(val, formatted)
	sort.Sort(ranked)

	result := make([]image.Summary, len(ranked))
	for i, r := range ranked {
		result[i] = originals[r.OriginalIndex]
	}

	s.FilteredItems = result

	if len(s.FilteredItems) <= s.SelectedIndex {
		s.SelectedIndex = len(s.FilteredItems) - 1
	}
}

func (s *ImageList) GetCurrentItem() image.Summary {
	return s.FilteredItems[s.SelectedIndex]
}
