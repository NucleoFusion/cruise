package home

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/cruise-org/cruise/internal/messages"
	styledhelp "github.com/cruise-org/cruise/internal/models/help"
	"github.com/cruise-org/cruise/pkg/keymap"
	"github.com/cruise-org/cruise/pkg/styles"
)

// - Container Events (Coded / Nice)
// Enabled Runtimes (Coded / Nice)

type Home struct {
	Width  int
	Height int
	SysRes *SysRes
	Stats  *QuickStats
	Logs   *Logs
	Help   styledhelp.StyledHelp
}

func NewHome(w int, h int) *Home {
	return &Home{
		Width:  w,
		Height: h,
		SysRes: NewSysRes(2*w/3, (h-11)/2-2),
		Stats:  NewQuickStats(w/3, (h-11)/2-2),
		Logs:   NewLogs(2*w/3, (h-11)/2),
		Help:   styledhelp.NewStyledHelp([]key.Binding{}, w-2),
	}
}

func (s *Home) Init() tea.Cmd {
	return tea.Batch(s.SysRes.Init(), s.Stats.Init(), s.Logs.Init())
}

func (s *Home) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case messages.SysResReadyMsg:
		sr, cmd := s.SysRes.Update(msg)
		s.SysRes = sr
		return s, cmd
	case messages.HomeStatContainer:
		st, cmd := s.Stats.Update(msg)
		s.Stats = st
		return s, cmd
	case messages.HomeStatVolume:
		st, cmd := s.Stats.Update(msg)
		s.Stats = st
		return s, cmd
	case messages.HomeStatImage:
		st, cmd := s.Stats.Update(msg)
		s.Stats = st
		return s, cmd
	case messages.HomeStatNetwork:
		st, cmd := s.Stats.Update(msg)
		s.Stats = st
		return s, cmd
	case messages.HomeLogsMonitor:
		st, cmd := s.Logs.Update(msg)
		s.Logs = st
		return s, cmd
	case messages.HomeLogsTick:
		st, cmd := s.Logs.Update(msg)
		s.Logs = st
		return s, cmd

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keymap.QuickQuitKey()):
			return s, tea.Quit
		}
	}
	return s, nil
}

func (s *Home) View() string {
	logo := lipgloss.Place(s.Width-2, 11, // use fixed height for title
		lipgloss.Center, lipgloss.Center, styles.TextStyle().Render(styles.LogoText))

	view := lipgloss.JoinVertical(lipgloss.Center, logo,
		lipgloss.JoinHorizontal(lipgloss.Center, s.SysRes.View(), s.Stats.View()),
		s.Logs.View(),
		s.Help.View(),
	)
	return styles.SceneStyle().Render(view)
}
