package volumes

import (
	"fmt"
	"strings"
	"time"

	"github.com/NucleoFusion/cruise/internal/keymap"
	"github.com/NucleoFusion/cruise/internal/messages"
	styledhelp "github.com/NucleoFusion/cruise/internal/models/help"
	"github.com/NucleoFusion/cruise/internal/runtimes/docker"
	"github.com/NucleoFusion/cruise/internal/styles"
	"github.com/NucleoFusion/cruise/internal/utils"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Volumes struct {
	Width      int
	Height     int
	List       *VolumeList
	Details    *VolumeDetail
	Keymap     keymap.VolMap
	Help       styledhelp.StyledHelp
	IsLoading  bool
	ShowDetail bool
}

func NewVolumes(w int, h int) *Volumes {
	return &Volumes{
		Width:      w,
		Height:     h,
		IsLoading:  true,
		ShowDetail: false,
		List:       NewVolumeList(w-2, h-5-strings.Count(styles.VolumesText, "\n")), // h-5 to account for styled help and title padding
		Keymap:     keymap.NewVolMap(),
		Help:       styledhelp.NewStyledHelp(keymap.NewVolMap().Bindings(), w-2),
	}
}

func (s *Volumes) Init() tea.Cmd {
	return s.List.Init()
}

func (s *Volumes) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case messages.VolumesReadyMsg:
		s.IsLoading = false
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
		case key.Matches(msg, s.Keymap.Remove):
			err := docker.RemoveVolumes(s.List.GetCurrentItem().Name)
			if err != nil {
				return s, utils.ReturnError("Volumes Page", "Error Removing Volume", err)
			}
			return s, tea.Batch(s.Refresh(), utils.ReturnMsg("Volumes Page", "Removed Volume",
				fmt.Sprintf("Successfully Removed Volume %s", s.List.GetCurrentItem().Name)))
		case key.Matches(msg, s.Keymap.Prune):
			err := docker.PruneVolumes()
			if err != nil {
				return s, utils.ReturnError("Volumes Page", "Error Pruning Volumes", err)
			}
			return s, tea.Batch(s.Refresh(), utils.ReturnMsg("Volumes Page", "Pruned Volumes",
				"Successfully Pruned Volumes"))
		case key.Matches(msg, s.Keymap.ExitDetails):
			if s.ShowDetail {
				s.ShowDetail = false
				return s, nil
			}
		case key.Matches(msg, s.Keymap.ShowDetails):
			s.ShowDetail = true
			s.Details = NewDetail(s.Width, s.Height, s.List.GetCurrentItem())
			return s, nil
		}
	}

	var cmd tea.Cmd
	s.List, cmd = s.List.Update(msg)
	return s, cmd
}

func (s *Volumes) View() string {
	if s.ShowDetail {
		return styles.SceneStyle().Render(s.Details.View())
	}

	return styles.SceneStyle().Render(
		lipgloss.JoinVertical(lipgloss.Center,
			styles.TextStyle().Padding(1, 0).Render(styles.VolumesText), s.GetListText(), s.Help.View()))
}

func (s *Volumes) GetListText() string {
	if s.IsLoading {
		return lipgloss.Place(s.Width-2, s.Height-4-strings.Count(styles.VolumesText, "\n"),
			lipgloss.Center, lipgloss.Center, styles.TextStyle().Render("Loading..."))
	}

	return lipgloss.NewStyle().Render(s.List.View())
}

func (s *Volumes) Refresh() tea.Cmd {
	return tea.Tick(0, func(_ time.Time) tea.Msg {
		vols, err := docker.GetVolumes()
		if err != nil {
			fmt.Println(err)
			return utils.ReturnError("Volumes Page", "Error Querying Volumes", err)
		}
		return messages.VolumesReadyMsg{Items: vols.Volumes}
	})
}
