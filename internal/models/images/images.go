package images

import (
	"strings"
	"time"

	"github.com/NucleoFusion/cruise/internal/docker"
	"github.com/NucleoFusion/cruise/internal/keymap"
	"github.com/NucleoFusion/cruise/internal/messages"
	"github.com/NucleoFusion/cruise/internal/styles"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Prune, Image Layers, CRUD images
// TODO: Search Repo's for image
// TODO: w/ Registry push/pull/build
// TODO: Vulnerability Scanning
// TODO: Rename Tags

type Images struct {
	Width     int
	Height    int
	List      *ImageList
	Keymap    keymap.ImagesMap
	Help      help.Model
	IsLoading bool
}

func NewImages(w int, h int) *Images {
	return &Images{
		Width:     w,
		Height:    h,
		IsLoading: true,
		List:      NewImageList(w-4, h-7-strings.Count(styles.ImagesText, "\n")),
		Keymap:    keymap.NewImagesMap(),
		Help:      help.New(),
	}
}

func (s *Images) Init() tea.Cmd {
	return s.List.Init()
}

func (s *Images) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case messages.ImagesReadyMsg:
		s.IsLoading = false

		var cmd tea.Cmd
		s.List, cmd = s.List.Update(msg)
		return s, cmd
	case messages.UpdateImagesMsg:
		var cmd tea.Cmd
		s.List, cmd = s.List.Update(msg)
		return s, cmd
	case tea.KeyMsg:
		if s.List.Ti.Focused() {
			var cmd tea.Cmd
			s.List, cmd = s.List.Update(msg)
			return s, cmd
		}
		switch {
		case key.Matches(msg, s.Keymap.Remove):
			err := docker.RemoveImage(s.List.GetCurrentItem().ID)
			if err != nil {
				return s, func() tea.Msg {
					return messages.ErrorMsg{
						Msg:   err.Error(),
						Locn:  "Images Page",
						Title: "Error Removing Image",
					}
				}
			}
			return s, nil
		case key.Matches(msg, s.Keymap.Prune):
			err := docker.PruneImages()
			if err != nil {
				return s, func() tea.Msg {
					return messages.ErrorMsg{
						Msg:   err.Error(),
						Locn:  "Images Page",
						Title: "Error Pruning Images",
					}
				}
			}
			return s, nil
			// case key.Matches(msg, s.Keymap.Layers):
			// 	// remove
		}
	}

	var cmd tea.Cmd
	s.List, cmd = s.List.Update(msg)
	return s, cmd
}

func (s *Images) View() string {
	return lipgloss.JoinVertical(lipgloss.Center,
		styles.TextStyle().Render(styles.ImagesText), s.GetListText(), s.Help.View(keymap.NewDynamic(s.Keymap.Bindings())))
}

func (s *Images) GetListText() string {
	if s.IsLoading {
		return lipgloss.Place(s.Width-2, s.Height-4-strings.Count(styles.ImagesText, "\n"),
			lipgloss.Center, lipgloss.Top, "Loading...")
	}

	return lipgloss.NewStyle().Padding(1).Render(s.List.View())
}

func (s *Images) Refresh() tea.Cmd {
	return tea.Tick(3*time.Second, func(_ time.Time) tea.Msg {
		items, err := docker.GetImages()
		if err != nil {
			return messages.ErrorMsg{
				Locn:  "Images Page",
				Msg:   err.Error(),
				Title: "Error Querying Images",
			}
		}
		return messages.ImagesReadyMsg{Items: items}
	})
}
