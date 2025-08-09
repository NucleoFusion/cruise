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
	"github.com/charmbracelet/bubbles/viewport"
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
	Vp        viewport.Model
	Help      help.Model
	ShowVp    bool
	IsLoading bool
}

func NewImages(w int, h int) *Images {
	vp := viewport.New(w*2/3, h/2)
	vp.Style = styles.PageStyle().Padding(1, 2)

	return &Images{
		Width:     w,
		Height:    h,
		IsLoading: true,
		List:      NewImageList(w-4, h-7-strings.Count(styles.ImagesText, "\n")),
		Keymap:    keymap.NewImagesMap(),
		Help:      help.New(),
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
				return s, func() tea.Msg {
					return messages.ErrorMsg{
						Msg:   err.Error(),
						Locn:  "Images Page",
						Title: "Error Pulling Image",
					}
				}
			}
			return s, nil
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
				return s, func() tea.Msg {
					return messages.ErrorMsg{
						Msg:   err.Error(),
						Locn:  "Images Page",
						Title: "Error Pushing Image",
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
				return s, func() tea.Msg {
					return messages.ErrorMsg{
						Msg:   err.Error(),
						Locn:  "Images Page",
						Title: "Error Querying Image Layers",
					}
				}
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
