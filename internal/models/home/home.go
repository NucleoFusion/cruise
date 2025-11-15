package home

import (
	"github.com/NucleoFusion/cruise/internal/keymap"
	"github.com/NucleoFusion/cruise/internal/messages"
	styledhelp "github.com/NucleoFusion/cruise/internal/models/help"
	"github.com/NucleoFusion/cruise/internal/styles"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Home struct {
	Width      int
	Height     int
	Logs       *Logs
	Daemon     *Daemon
	SysRes     *SysRes
	QuickStats *QuickStats
	Help       styledhelp.StyledHelp
}

func NewHome(w int, h int) *Home {
	return &Home{
		Width:      w,
		Height:     h,
		Logs:       NewLogs(w, h),
		Daemon:     NewDaemon(w, h),
		SysRes:     NewSysRes(w, h),
		QuickStats: NewQuickStats(w, h),
		Help:       styledhelp.NewStyledHelp([]key.Binding{}, w),
	}
}

func (s *Home) Init() tea.Cmd {
	return tea.Batch(s.SysRes.Init(), s.Logs.Init(), messages.TickDashboard())
}

func (s *Home) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case messages.DashboardTick:
		cmd := s.Refresh()
		return s, tea.Batch(cmd, messages.TickDashboard())
	case messages.SysResReadyMsg:
		var cmd tea.Cmd
		s.SysRes, cmd = s.SysRes.Update(msg)
		return s, cmd
	case messages.NewEvents:
		var cmd tea.Cmd
		s.Logs, cmd = s.Logs.Update(msg)
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
	logo := lipgloss.Place(s.Width-2, s.Height/3-2, // Accounting for the border
		lipgloss.Center, lipgloss.Center, styles.TextStyle().Render(styles.LogoText))
	sysres := s.SysRes.View()
	daemon := s.Daemon.View()
	stats := s.QuickStats.View()
	logs := s.Logs.View()

	view := lipgloss.JoinVertical(lipgloss.Center, logo,
		lipgloss.JoinHorizontal(lipgloss.Center, sysres, stats),
		lipgloss.JoinHorizontal(lipgloss.Center, daemon, logs),
		s.Help.View(),
	)

	return styles.SceneStyle().Render(view)
}

func (s *Home) Refresh() tea.Cmd {
	return tea.Batch(s.SysRes.Refresh())
}
