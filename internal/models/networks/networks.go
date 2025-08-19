package networks

import (
	"strings"
	"time"

	"github.com/NucleoFusion/cruise/internal/docker"
	"github.com/NucleoFusion/cruise/internal/keymap"
	"github.com/NucleoFusion/cruise/internal/messages"
	"github.com/NucleoFusion/cruise/internal/styles"
	"github.com/NucleoFusion/cruise/internal/utils"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Networks struct {
	Width   int
	Height  int
	List    *NetworkList
	Details *NetworkDetail
	// Keymap    keymap.NetworksMap
	Help       help.Model
	IsLoading  bool
	ShowDetail bool
}

func NewNetworks(w int, h int) *Networks {
	return &Networks{
		Width:      w,
		Height:     h,
		IsLoading:  true,
		ShowDetail: false,
		List:       NewNetworkList(w-4, h-7-strings.Count(styles.NetworksText, "\n")),
		// Keymap:    keymap.NewNetworksMap(),
		Help: help.New(),
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
		switch msg.String() {
		case "enter":
			s.ShowDetail = true
			s.Details = NewDetail(s.Width, s.Height, s.List.GetCurrentItem())
			return s, nil
		case "esc":
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

func (s *Networks) View() string {
	if s.ShowDetail {
		return s.Details.View()
	}

	return lipgloss.JoinVertical(lipgloss.Center,
		styles.TextStyle().Render(styles.NetworksText), s.GetListText(), s.Help.View(keymap.NewDynamic([]key.Binding{})))
}

func (s *Networks) GetListText() string {
	if s.IsLoading {
		return lipgloss.Place(s.Width-2, s.Height-4-strings.Count(styles.NetworksText, "\n"),
			lipgloss.Center, lipgloss.Top, "Loading...")
	}

	return lipgloss.NewStyle().Padding(1).Render(s.List.View())
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
