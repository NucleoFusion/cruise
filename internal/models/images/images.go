package images

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/cruise-org/cruise/internal/messages"
	styledhelp "github.com/cruise-org/cruise/internal/models/help"
	"github.com/cruise-org/cruise/internal/utils"
	"github.com/cruise-org/cruise/pkg/colors"
	"github.com/cruise-org/cruise/pkg/keymap"
	"github.com/cruise-org/cruise/pkg/runtimes"
	"github.com/cruise-org/cruise/pkg/styles"
)

type Images struct {
	Width     int
	Height    int
	List      *ImageList
	Keymap    keymap.ImagesMap
	Vp        viewport.Model
	Help      styledhelp.StyledHelp
	ShowVp    bool
	IsLoading bool
}

func NewImages(w int, h int) *Images {
	vp := viewport.New(w/3, h/2)
	vp.Style = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(colors.Load().FocusedBorder).
		Padding(1).Foreground(colors.Load().Text)

	return &Images{
		Width:     w,
		Height:    h,
		IsLoading: true,
		List:      NewImageList(w-2, h-5-strings.Count(styles.ImagesText, "\n")), // h-5 to account for styled help and title padding
		Keymap:    keymap.NewImagesMap(),
		Help:      styledhelp.NewStyledHelp(keymap.NewImagesMap().Bindings(), w-2),
		Vp:        vp,
		ShowVp:    false,
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
		} else if s.ShowVp {
			if key.Matches(msg, key.NewBinding(key.WithKeys("esc"))) {
				s.ShowVp = false
				return s, nil
			}
			var cmd tea.Cmd
			s.Vp, cmd = s.Vp.Update(msg)
			return s, cmd
		}
		switch {
		case key.Matches(msg, keymap.QuickQuitKey()):
			return s, tea.Quit
		case key.Matches(msg, s.Keymap.Remove):
			curr := s.List.GetCurrentItem()
			err := runtimes.RuntimeSrv.RemoveImage(context.Background(), curr.Runtime, curr.ID)
			if err != nil {
				return s, utils.ReturnError("Images Page", "Error Removing Image", err)
			}

			return s, tea.Batch(s.UpdateImages(), utils.ReturnMsg("Images Page", "Removing Image",
				fmt.Sprintf("Successfully Removed Image w/ ID %s", s.List.GetCurrentItem().ID)))
		case key.Matches(msg, s.Keymap.Pull):
			curr := s.List.GetCurrentItem()
			err := runtimes.RuntimeSrv.PullImage(context.Background(), curr.Runtime, curr.ID)
			if err != nil {
				return s, utils.ReturnError("Images Page", "Error Pulling Image", err)
			}
			return s, tea.Batch(s.UpdateImages(), utils.ReturnMsg("Images Page", "Pulling Image",
				fmt.Sprintf("Successfully Pulled Image w/ ID %s", s.List.GetCurrentItem().ID)))
		case key.Matches(msg, s.Keymap.Push):
			curr := s.List.GetCurrentItem()
			err := runtimes.RuntimeSrv.PushImage(context.Background(), curr.Runtime, curr.ID)
			if err != nil {
				return s, utils.ReturnError("Images Page", "Error Pushing Image", err)
			}
			return s, tea.Batch(s.UpdateImages(), utils.ReturnMsg("Images Page", "Pushing Image",
				fmt.Sprintf("Successfully Pushed Image w/ ID %s", s.List.GetCurrentItem().ID)))
		case key.Matches(msg, s.Keymap.Prune):
			curr := s.List.GetCurrentItem()
			err := runtimes.RuntimeSrv.PruneImages(context.Background(), curr.Runtime)
			if err != nil {
				return s, utils.ReturnError("Images Page", "Error Pruning Image", err)
			}
			return s, tea.Batch(s.UpdateImages(), utils.ReturnMsg("Images Page", "Pruning Image", "Successfully Pruned Images"))
		case key.Matches(msg, s.Keymap.Sync):
			return s, s.UpdateImages()
		case key.Matches(msg, s.Keymap.Layers):
			curr := s.List.GetCurrentItem()
			text, err := runtimes.RuntimeSrv.ImageLayers(context.Background(), curr.Runtime, curr.ID)
			if err != nil {
				return s, utils.ReturnError("Images Page", "Error Querying Image Layers", err)
			}

			s.Vp.SetContent(text)

			s.ShowVp = true
			return s, nil
		}
	}

	var cmd tea.Cmd
	s.List, cmd = s.List.Update(msg)
	return s, cmd
}

func (s *Images) View() string {
	if s.ShowVp {
		return lipgloss.Place(s.Width, s.Height, lipgloss.Center, lipgloss.Center, s.Vp.View())
	}

	return styles.SceneStyle().Render(
		lipgloss.JoinVertical(lipgloss.Center,
			styles.TextStyle().Padding(1, 0).Render(styles.ImagesText), s.GetListText(), s.Help.View()))
}

func (s *Images) GetListText() string {
	if s.IsLoading {
		return lipgloss.Place(s.Width-2, s.Height-4-strings.Count(styles.ImagesText, "\n"),
			lipgloss.Center, lipgloss.Top, "Loading...")
	}

	return lipgloss.NewStyle().Render(s.List.View())
}

func (s *Images) UpdateImages() tea.Cmd {
	return tea.Tick(0, func(_ time.Time) tea.Msg {
		imgs, err := runtimes.RuntimeSrv.Images(context.Background())
		if err != nil {
			return utils.ReturnError("Images Page", "Error Querying Images", err)
		}
		return messages.UpdateImagesMsg{Items: imgs}
	})
}
