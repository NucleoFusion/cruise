package containers

import (
	"fmt"
	"os/exec"
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

type Containers struct {
	Width     int
	Height    int
	List      *ContainerList
	Keymap    keymap.ContainersMap
	Help      help.Model
	IsLoading bool
}

func NewContainers(w int, h int) *Containers {
	return &Containers{
		Width:     w,
		Height:    h,
		IsLoading: true,
		List:      NewContainerList(w-4, h-7-strings.Count(styles.ContainersText, "\n")),
		Keymap:    keymap.NewContainersMap(),
		Help:      help.New(),
	}
}

func (s *Containers) Init() tea.Cmd {
	return s.List.Init()
}

func (s *Containers) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case messages.NewContainerDetails:
		var cmd tea.Cmd
		s.List, cmd = s.List.Update(msg)
		return s, cmd
	case messages.ContainerReadyMsg:
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
		switch {
		case key.Matches(msg, s.Keymap.Start):
			err := docker.StartContainer(s.List.GetCurrentItem().ID)
			if err != nil {
				return s, func() tea.Msg {
					return messages.ErrorMsg{
						Title: "Error Starting Container",
						Msg:   err.Error(),
						Locn:  "Containers Page",
					}
				}
			}
			return s, nil
		case key.Matches(msg, s.Keymap.Pause):
			err := docker.PauseContainer(s.List.GetCurrentItem().ID)
			if err != nil {
				return s, func() tea.Msg {
					return messages.ErrorMsg{
						Title: "Error Pausing Container",
						Msg:   err.Error(),
						Locn:  "Containers Page",
					}
				}
			}
			return s, nil
		case key.Matches(msg, s.Keymap.Unpause):
			err := docker.UnpauseContainer(s.List.GetCurrentItem().ID)
			if err != nil {
				return s, func() tea.Msg {
					return messages.ErrorMsg{
						Title: "Error Unpausing Container",
						Msg:   err.Error(),
						Locn:  "Containers Page",
					}
				}
			}
			return s, nil
		case key.Matches(msg, s.Keymap.Remove):
			err := docker.RemoveContainer(s.List.GetCurrentItem().ID)
			if err != nil {
				return s, func() tea.Msg {
					return messages.ErrorMsg{
						Title: "Error Removing Container",
						Msg:   err.Error(),
						Locn:  "Containers Page",
					}
				}
			}
			return s, nil
		case key.Matches(msg, s.Keymap.Restart):
			err := docker.RestartContainer(s.List.GetCurrentItem().ID)
			if err != nil {
				return s, func() tea.Msg {
					return messages.ErrorMsg{
						Title: "Error Restarting Container",
						Msg:   err.Error(),
						Locn:  "Containers Page",
					}
				}
			}
			return s, nil
		case key.Matches(msg, s.Keymap.Stop):
			err := docker.StopContainer(s.List.GetCurrentItem().ID)
			if err != nil {
				return s, func() tea.Msg {
					return messages.ErrorMsg{
						Title: "Error Stopping Container",
						Msg:   err.Error(),
						Locn:  "Containers Page",
					}
				}
			}
			return s, nil
		case key.Matches(msg, s.Keymap.Exec):
			cmd := exec.Command("ghostty", "-e", fmt.Sprintf("docker exec -it %s %s", s.List.GetCurrentItem().ID, "sh"))
			cmd.Stdin = nil
			cmd.Stdout = nil
			cmd.Stderr = nil

			err := cmd.Start()
			if err != nil {
				return s, func() tea.Msg {
					return messages.ErrorMsg{
						Title: "Error Execing into Container",
						Msg:   err.Error(),
						Locn:  "Containers Page",
					}
				}
			}

			return s, nil
		}
	}

	var cmd tea.Cmd
	s.List, cmd = s.List.Update(msg)
	return s, cmd
}

func (s *Containers) View() string {
	return lipgloss.JoinVertical(lipgloss.Center,
		styles.TextStyle().Render(styles.ContainersText), s.GetListText(), s.Help.View(keymap.NewDynamic(s.Keymap.Bindings())))
}

func (s *Containers) GetListText() string {
	if s.IsLoading {
		return lipgloss.Place(s.Width-2, s.Height-4-strings.Count(styles.ContainersText, "\n"),
			lipgloss.Center, lipgloss.Top, "Loading...")
	}

	return lipgloss.NewStyle().Padding(1).Render(s.List.View())
}

func (s *Containers) Refresh() tea.Cmd {
	return tea.Tick(3*time.Second, func(_ time.Time) tea.Msg {
		items, err := docker.GetContainers()
		return messages.ContainerReadyMsg{
			Items: items,
			Err:   err,
		}
	})
}
