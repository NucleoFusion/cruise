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
		Logs:       NewLogs((w-2)-(w-2)/4, (h-14)-(h-14)/2), // h-15 to account for styled help and title, w-2 for scene padding
		Daemon:     NewDaemon((w-2)/4, (h-14)-(h-14)/2),
		SysRes:     NewSysRes((w-2)-(w-2)/4, (h-14)/2),
		QuickStats: NewQuickStats((w-2)/4, (h-14)/2),
		Help:       styledhelp.NewStyledHelp([]key.Binding{}, w-2),
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
	logo := lipgloss.Place(s.Width-2, 11, // use fixed height for title
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
