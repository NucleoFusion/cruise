package home

import (
	"github.com/NucleoFusion/cruise/internal/messages"
	"github.com/NucleoFusion/cruise/internal/styles"
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
}

func NewHome(w int, h int) *Home {
	return &Home{
		Width:      w,
		Height:     h,
		Logs:       NewLogs(w, h),
		Daemon:     NewDaemon(w, h),
		SysRes:     NewSysRes(w, h),
		QuickStats: NewQuickStats(w, h),
	}
}

func (s *Home) Init() tea.Cmd {
	return tea.Batch(s.SysRes.Init(), messages.TickDashboard())
}

func (s *Home) Update(msg tea.Msg) (*Home, tea.Cmd) {
	switch msg := msg.(type) {
	case messages.DashboardTick:
		s.Refresh()
		return s, messages.TickDashboard()
	case messages.SysResReadyMsg:
		var cmd tea.Cmd
		s.SysRes, cmd = s.SysRes.Update(msg)
		return s, cmd
	case tea.KeyMsg:
		switch msg.String() {
		}
	}
	return s, nil
}

func (s *Home) View() string {
	logo := lipgloss.Place(s.Width-2, s.Height/3-2, // Accounting for the border
		lipgloss.Center, lipgloss.Center, styles.TitleStyle().Render(styles.LogoText))
	sysres := s.SysRes.View()
	daemon := s.Daemon.View()
	stats := s.QuickStats.View()
	logs := s.Logs.View()

	view := lipgloss.JoinVertical(lipgloss.Left, logo,
		lipgloss.JoinHorizontal(lipgloss.Center, sysres, stats),
		lipgloss.JoinHorizontal(lipgloss.Left, daemon, logs),
	)

	return view
}

func (s *Home) Refresh() {
	s.SysRes.Refresh()
}
