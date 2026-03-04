package networks

import (
	"context"
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/cruise-org/cruise/internal/messages"
	detailrenderer "github.com/cruise-org/cruise/internal/models/detailRenderer"
	styledhelp "github.com/cruise-org/cruise/internal/models/help"
	"github.com/cruise-org/cruise/internal/utils"
	"github.com/cruise-org/cruise/pkg/keymap"
	"github.com/cruise-org/cruise/pkg/runtimes"
	"github.com/cruise-org/cruise/pkg/styles"
)

type Networks struct {
	Width      int
	Height     int
	List       *NetworkList
	Details    *detailrenderer.DetailRenderer
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
		List:       NewNetworkList(w-2, h-5-strings.Count(styles.NetworksText, "\n")), // h-5 to account for styled help and title padding
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
	case messages.DetailRendererInit:
		if s.Details == nil {
			return s, nil
		}
		dr, cmd := s.Details.Update(msg)
		if details, ok := dr.(*detailrenderer.DetailRenderer); ok {
			s.Details = details
		}
		return s, cmd

	case messages.DetailRendererContent:
		if s.Details == nil {
			return s, nil
		}

		dr, cmd := s.Details.Update(msg)
		if details, ok := dr.(*detailrenderer.DetailRenderer); ok {
			s.Details = details
		}
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
		} else if s.ShowDetail {
			if key.Matches(msg, s.Keymap.ExitDetails) {
				s.ShowDetail = false
			}
			return s, nil
		}
		switch {
		case key.Matches(msg, keymap.QuickQuitKey()):
			return s, tea.Quit
		case key.Matches(msg, s.Keymap.ShowDetails):
			s.Details = detailrenderer.NewDetailRenderer(s.Width, s.Height, s.detailsStatFunc(), s.detailsRenderFunc())
			s.ShowDetail = true
			return s, s.Details.Init()
		case key.Matches(msg, s.Keymap.ExitDetails):
		case key.Matches(msg, s.Keymap.Remove):
			curr := s.List.GetCurrentItem()
			err := runtimes.RuntimeSrv.RemoveNetwork(context.Background(), curr.Runtime, curr.ID)
			if err != nil {
				return s, utils.ReturnError("Networks Page", "Error Removing Network", err)
			}
			return s, tea.Batch(s.Init(), utils.ReturnMsg("Networks Page", "Removed Network",
				fmt.Sprintf("Successfully Removed Networks w/ ID %s", s.List.GetCurrentItem().ID)))
		case key.Matches(msg, s.Keymap.Prune):
			curr := s.List.GetCurrentItem()
			err := runtimes.RuntimeSrv.PruneNetworks(context.Background(), curr.Runtime)
			if err != nil {
				return s, utils.ReturnError("Networks Page", "Error Pruning Networks", err)
			}
			return s, tea.Batch(s.Init(), utils.ReturnMsg("Networks Page", "Pruned Networks",
				"Successfully Pruned Networks"))
		}
	}

	var cmd tea.Cmd
	s.List, cmd = s.List.Update(msg)
	return s, cmd
}

func (s *Networks) View() string {
	if s.ShowDetail {
		return s.Details.View()
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
