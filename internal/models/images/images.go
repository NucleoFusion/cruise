package images

import (
	"strings"
	"time"

	"github.com/NucleoFusion/cruise/internal/docker"
	"github.com/NucleoFusion/cruise/internal/messages"
	"github.com/NucleoFusion/cruise/internal/styles"
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Images struct {
	Width  int
	Height int
	List   *ImageList
	// Keymap    keymap.ImagesMap
	Help      help.Model
	IsLoading bool
}

func NewImages(w int, h int) *Images {
	return &Images{
		Width:     w,
		Height:    h,
		IsLoading: true,
		List:      NewImageList(w-4, h-7-strings.Count(styles.ImagesText, "\n")),
		// Keymap:    keymap.NewImagesMap(),
		Help: help.New(),
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
	case tea.KeyMsg:
		if s.List.Ti.Focused() {
			var cmd tea.Cmd
			s.List, cmd = s.List.Update(msg)
			return s, cmd
		}
	}

	var cmd tea.Cmd
	s.List, cmd = s.List.Update(msg)
	return s, cmd
}

func (s *Images) View() string {
	return lipgloss.JoinVertical(lipgloss.Center,
		styles.TextStyle().Render(styles.ImagesText), s.GetListText(), "help")
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
