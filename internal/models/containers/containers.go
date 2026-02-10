package containers

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
	"github.com/cruise-org/cruise/pkg/config"
	"github.com/cruise-org/cruise/pkg/keymap"
	"github.com/cruise-org/cruise/pkg/runtimes"
	"github.com/cruise-org/cruise/pkg/styles"
)

type Containers struct {
	Width       int
	Height      int
	List        *ContainerList
	Details     *ContainerDetail
	Vp          viewport.Model
	Keymap      keymap.ContainersMap
	Help        styledhelp.StyledHelp
	IsLoading   bool
	ShowPortmap bool
	ShowDetail  bool
}

func NewContainers(w int, h int) *Containers {
	vp := viewport.New(w/3, h/2)
	vp.Style = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(colors.Load().FocusedBorder).
		Padding(1).Foreground(colors.Load().Text)

	return &Containers{
		Width:     w,
		Height:    h,
		IsLoading: true,
		List:      NewContainerList(w-2, h-5-strings.Count(styles.ContainersText, "\n")), // h-5 to account for styled help and title padding
		Keymap:    keymap.NewContainersMap(),
		Help:      styledhelp.NewStyledHelp(keymap.NewContainersMap().Bindings(), w-2),
		Vp:        vp,
	}
}

func (s *Containers) Init() tea.Cmd {
	return s.List.Init()
}

func (s *Containers) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case messages.ContainerDetailsReady:
		if s.ShowDetail {
			var cmd tea.Cmd
			s.Details, cmd = s.Details.Update(msg)
			return s, cmd
		}
	case messages.ContainerDetailsTick:
		if s.ShowDetail {
			var cmd tea.Cmd
			s.Details, cmd = s.Details.Update(msg)
			return s, cmd
		}
	case messages.NewContainerDetails:
		var cmd tea.Cmd
		s.List, cmd = s.List.Update(msg)
		return s, cmd
	case messages.ContainerReadyMsg:
		s.IsLoading = false

		var cmd tea.Cmd
		s.List, cmd = s.List.Update(msg)
		return s, cmd
	case messages.PortMapMsg:
		if msg.Err != nil {
			return s, utils.ReturnError("Containers Page", "Error Getting Ports", msg.Err)
		}

		s.Vp.SetContent(strings.Join(msg.Ports, "\n"))
		if len(msg.Ports) == 0 {
			s.Vp.SetContent("No Port Mappings found")
		}

		return s, nil
	case tea.KeyMsg:
		if s.List.Ti.Focused() {
			var cmd tea.Cmd
			s.List, cmd = s.List.Update(msg)
			return s, cmd
		} else if s.ShowPortmap {
			if msg.String() == config.Cfg.Keybinds.Container.ExitDetails {
				s.ShowPortmap = false
				return s, nil
			}
			var cmd tea.Cmd
			s.Vp, cmd = s.Vp.Update(msg)
			return s, cmd
		}
		switch {
		case key.Matches(msg, keymap.QuickQuitKey()):
			return s, tea.Quit
		case key.Matches(msg, s.Keymap.Start):
			curr := s.List.GetCurrentItem()
			err := runtimes.RuntimeSrv.StartContainer(context.Background(), curr.Runtime, curr.ID)
			if err != nil {
				return s, utils.ReturnError("Containers Page", "Error Starting Contianer", err)
			}
			return s, utils.ReturnMsg("Container Page", "Started Container",
				fmt.Sprintf("Successfully Started Container w/ ID %s", s.List.GetCurrentItem().ID))

		case key.Matches(msg, s.Keymap.Pause):
			curr := s.List.GetCurrentItem()
			err := runtimes.RuntimeSrv.PauseContainer(context.Background(), curr.Runtime, curr.ID)
			if err != nil {
				return s, utils.ReturnError("Containers Page", "Error Pausing Contianer", err)
			}

			return s, utils.ReturnMsg("Container Page", "Pausing Container",
				fmt.Sprintf("Successfully Pausing Container w/ ID %s", s.List.GetCurrentItem().ID))

		case key.Matches(msg, s.Keymap.Unpause):
			curr := s.List.GetCurrentItem()
			err := runtimes.RuntimeSrv.UnpauseContainer(context.Background(), curr.Runtime, curr.ID)
			if err != nil {
				return s, utils.ReturnError("Containers Page", "Error Unpausing Contianer", err)
			}
			return s, utils.ReturnMsg("Container Page", "Unpausing Container",
				fmt.Sprintf("Successfully Unpausing Container w/ ID %s", s.List.GetCurrentItem().ID))

		case key.Matches(msg, s.Keymap.Remove):
			curr := s.List.GetCurrentItem()
			err := runtimes.RuntimeSrv.RemoveContainer(context.Background(), curr.Runtime, curr.ID)
			if err != nil {
				return s, utils.ReturnError("Containers Page", "Error Removing Contianer", err)
			}
			return s, utils.ReturnMsg("Container Page", "Removing Container",
				fmt.Sprintf("Successfully Removing Container w/ ID %s", s.List.GetCurrentItem().ID))

		case key.Matches(msg, s.Keymap.Restart):
			curr := s.List.GetCurrentItem()
			err := runtimes.RuntimeSrv.RestartContainer(context.Background(), curr.Runtime, curr.ID)
			if err != nil {
				return s, utils.ReturnError("Containers Page", "Error Restarting Contianer", err)
			}
			return s, utils.ReturnMsg("Container Page", "Restarting Container",
				fmt.Sprintf("Successfully Restarting Container w/ ID %s", s.List.GetCurrentItem().ID))

		case key.Matches(msg, s.Keymap.Stop):
			curr := s.List.GetCurrentItem()
			err := runtimes.RuntimeSrv.StopContainer(context.Background(), curr.Runtime, curr.ID)
			if err != nil {
				return s, utils.ReturnError("Containers Page", "Error Stopping Contianer", err)
			}

			return s, utils.ReturnMsg("Container Page", "Stopping Container",
				fmt.Sprintf("Successfully Stopped Container w/ ID %s", s.List.GetCurrentItem().ID))

		case key.Matches(msg, s.Keymap.Exec):
			curr := s.List.GetCurrentItem()
			cmd := runtimes.RuntimeSrv.ExecContainer(context.Background(), curr.Runtime, curr.ID)
			cmd.Stdin = nil
			cmd.Stdout = nil
			cmd.Stderr = nil

			err := cmd.Start()
			if err != nil {
				return s, utils.ReturnError("Containers Page", "Error Execing into Container", err)
			}
			return s, nil
		case key.Matches(msg, s.Keymap.PortMap):
			s.ShowPortmap = true
			s.Vp.SetContent("Loading...")

			return s, tea.Tick(0, func(_ time.Time) tea.Msg {
				curr := s.List.GetCurrentItem()
				ports, err := runtimes.RuntimeSrv.PortsMap(context.Background(), curr.Runtime, curr.ID)
				return messages.PortMapMsg{Ports: ports, Err: err}
			})

		// TODO: Details
		case key.Matches(msg, s.Keymap.ShowDetails):
			s.ShowDetail = true
			// s.Details = NewDetail(s.Width, s.Height, s.List.GetCurrentItem())
			return s, s.Details.Init()
		case key.Matches(msg, s.Keymap.ExitDetails):
			if s.ShowDetail {
				s.ShowDetail = false
				return s, nil
			}
		}
	}

	var cmd tea.Cmd
	s.List, cmd = s.List.Update(msg)
	return s, cmd
}

func (s *Containers) View() string {
	if s.ShowPortmap {
		return lipgloss.Place(s.Width, s.Height, lipgloss.Center, lipgloss.Center, s.Vp.View())
	}

	if s.ShowDetail {
		return styles.SceneStyle().Render(s.Details.View())
	}

	return styles.SceneStyle().Render(
		lipgloss.JoinVertical(lipgloss.Center,
			styles.TextStyle().Padding(1, 0).Render(styles.ContainersText), s.GetListText(), s.Help.View()))
}

func (s *Containers) GetListText() string {
	if s.IsLoading {
		return lipgloss.Place(s.Width-2, s.Height-4-strings.Count(styles.ContainersText, "\n"),
			lipgloss.Center, lipgloss.Top, "Loading...")
	}

	return lipgloss.NewStyle().Render(s.List.View())
}
