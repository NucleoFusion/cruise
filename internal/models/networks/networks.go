package networks

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
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Networks struct {
	Width      int
	Height     int
	List       *NetworkList
	Details    *NetworkDetail
	Keymap     keymap.NetMap
	Help       styledhelp.StyledHelp
	IsLoading  bool
	ShowDetail bool
}

func NewNetworks(w int, h int) *Networks {
	return &Networks{
		Width:      w,
		Height:     h,
		IsLoading:  true,
		ShowDetail: false,
		List:       NewNetworkList(w-2, h-5-strings.Count(styles.NetworksText, "\n")), //h-5 to account for styled help and title padding
		Keymap:     keymap.NewNetMap(),
		Help:       styledhelp.NewStyledHelp(keymap.NewNetMap().Bindings(), w-2),
	}
}

func (s *Networks) Init() tea.Cmd {
	return s.List.Init()
}

func (s *Networks) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case messages.NetworksReadyMsg:
		s.IsLoading = false

		var cmd tea.Cmd
		s.List, cmd = s.List.Update(msg)
		return s, cmd
	case messages.UpdateNetworksMsg:
		var cmd tea.Cmd
		s.List, cmd = s.List.Update(msg)
		return s, cmd
	case messages.CloseDetails:
		s.ShowDetail = false
		return s, nil
	case tea.KeyMsg:
		if s.List.Ti.Focused() {
			var cmd tea.Cmd
			s.List, cmd = s.List.Update(msg)
			return s, cmd
		}
		switch {
		case key.Matches(msg, keymap.QuickQuitKey()):
			return s, tea.Quit
		case key.Matches(msg, s.Keymap.ShowDetails):
			s.ShowDetail = true
			s.Details = NewDetail(s.Width, s.Height, s.List.GetCurrentItem())
			return s, nil
		case key.Matches(msg, s.Keymap.ExitDetails):
			if s.ShowDetail {
				s.ShowDetail = false
				return s, nil
			}
		case key.Matches(msg, s.Keymap.Remove):
			err := docker.RemoveNetwork(s.List.GetCurrentItem().ID)
			if err != nil {
				return s, utils.ReturnError("Networks Page", "Error Removing Network", err)
			}
			return s, tea.Batch(s.Refresh(), utils.ReturnMsg("Networks Page", "Removed Network",
				fmt.Sprintf("Successfully Removed Networks w/ ID %s", s.List.GetCurrentItem().ID)))
		case key.Matches(msg, s.Keymap.Prune):
			err := docker.PruneNetworks()
			if err != nil {
				return s, utils.ReturnError("Networks Page", "Error Pruning Networks", err)
			}
			return s, tea.Batch(s.Refresh(), utils.ReturnMsg("Networks Page", "Pruned Networks",
				"Successfully Pruned Networks"))
		}
	}

	var cmd tea.Cmd
	s.List, cmd = s.List.Update(msg)
	return s, cmd
}

func (s *Networks) View() string {
	if s.ShowDetail {
		return styles.SceneStyle().Render(s.Details.View())
	}

	return styles.SceneStyle().Render(
		lipgloss.JoinVertical(lipgloss.Center,
			styles.TextStyle().Padding(1, 0).Render(styles.NetworksText), s.GetListText(), s.Help.View()))
}

func (s *Networks) GetListText() string {
	if s.IsLoading {
		return lipgloss.Place(s.Width-2, s.Height-4-strings.Count(styles.NetworksText, "\n"),
			lipgloss.Center, lipgloss.Top, "Loading...")
	}

	return lipgloss.NewStyle().Render(s.List.View())
}

func (s *Networks) Refresh() tea.Cmd {
	return tea.Tick(3*time.Second, func(_ time.Time) tea.Msg {
		items, err := docker.GetNetworks()
		if err != nil {
			return utils.ReturnError("Networks Page", "Error Querying Networks", err)
		}
		return messages.NetworksReadyMsg{Items: items}
	})
}
