package images

import (
	"fmt"
	"strings"
	"time"

	"github.com/NucleoFusion/cruise/internal/docker"
	"github.com/NucleoFusion/cruise/internal/keymap"
	"github.com/NucleoFusion/cruise/internal/messages"
	styledhelp "github.com/NucleoFusion/cruise/internal/models/help"
	"github.com/NucleoFusion/cruise/internal/styles"
	"github.com/NucleoFusion/cruise/internal/utils"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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
	vp := viewport.New(w*2/3, h/2-2)
	vp.Style = styles.PageStyle().Padding(1, 2)

	return &Images{
		Width:     w,
		Height:    h,
		IsLoading: true,
		List:      NewImageList(w-4, h-8-strings.Count(styles.ImagesText, "\n")),
		Keymap:    keymap.NewImagesMap(),
		Help:      styledhelp.NewStyledHelp(keymap.NewImagesMap().Bindings(), w),
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
			err := docker.RemoveImage(s.List.GetCurrentItem().ID)
			if err != nil {
				return s, utils.ReturnError("Images Page", "Error Removing Image", err)
			}
			return s, tea.Batch(s.UpdateImages(), utils.ReturnMsg("Images Page", "Removing Image",
				fmt.Sprintf("Successfully Removed Image w/ ID %s", s.List.GetCurrentItem().ID)))
		case key.Matches(msg, s.Keymap.Pull):
			curr := s.List.GetCurrentItem()
			img := curr.ID // Accurately get the image name
			if len(curr.RepoTags) > 0 && curr.RepoTags[0] != "<none>:<none>" {
				img = curr.RepoTags[0]
			} else if len(curr.RepoDigests) > 0 {
				img = curr.RepoDigests[0]
			}

			err := docker.PullImage(img)
			if err != nil {
				return s, utils.ReturnError("Images Page", "Error Pulling Image", err)
			}
			return s, tea.Batch(s.UpdateImages(), utils.ReturnMsg("Images Page", "Pulling Image",
				fmt.Sprintf("Successfully Pulled Image w/ ID %s", s.List.GetCurrentItem().ID)))
		case key.Matches(msg, s.Keymap.Push):
			curr := s.List.GetCurrentItem()
			img := curr.ID // Accurately get the image name
			if len(curr.RepoTags) > 0 && curr.RepoTags[0] != "<none>:<none>" {
				img = curr.RepoTags[0]
			} else if len(curr.RepoDigests) > 0 {
				img = curr.RepoDigests[0]
			}

			err := docker.PushImage(img)
			if err != nil {
				return s, utils.ReturnError("Images Page", "Error Pushing Image", err)
			}
			return s, tea.Batch(s.UpdateImages(), utils.ReturnMsg("Images Page", "Pushing Image",
				fmt.Sprintf("Successfully Pushed Image w/ ID %s", s.List.GetCurrentItem().ID)))
		case key.Matches(msg, s.Keymap.Prune):
			err := docker.PruneImages()
			if err != nil {
				return s, utils.ReturnError("Images Page", "Error Pruning Image", err)
			}
			return s, tea.Batch(s.UpdateImages(), utils.ReturnMsg("Images Page", "Pruning Image", "Successfully Pruned Images"))
		case key.Matches(msg, s.Keymap.Sync):
			return s, s.UpdateImages()
		case key.Matches(msg, s.Keymap.Layers):
			curr := s.List.GetCurrentItem()
			img := curr.ID // Accurately get the image name
			if len(curr.RepoTags) > 0 && curr.RepoTags[0] != "<none>:<none>" {
				img = curr.RepoTags[0]
			} else if len(curr.RepoDigests) > 0 {
				img = curr.RepoDigests[0]
			}

			text, err := docker.ImageHistory(img)
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

	return lipgloss.JoinVertical(lipgloss.Center,
		styles.TextStyle().Render(styles.ImagesText), s.GetListText(), s.Help.View())
}

func (s *Images) GetListText() string {
	if s.IsLoading {
		return lipgloss.Place(s.Width-2, s.Height-4-strings.Count(styles.ImagesText, "\n"),
			lipgloss.Center, lipgloss.Top, "Loading...")
	}

	return lipgloss.NewStyle().PaddingLeft(1).Render(s.List.View())
}

func (s *Images) UpdateImages() tea.Cmd {
	return tea.Tick(0, func(_ time.Time) tea.Msg {
		images, err := docker.GetImages()
		if err != nil {
			return utils.ReturnError("Images Page", "Error Querying Images", err)
		}
		return messages.UpdateImagesMsg{Items: images}
	})
}
