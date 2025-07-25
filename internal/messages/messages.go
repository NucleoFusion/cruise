package messages

import (
	"time"

	"github.com/NucleoFusion/cruise/internal/data"
	tea "github.com/charmbracelet/bubbletea"
)

type DashboardTick time.Time

func TickDashboard() tea.Cmd {
	return tea.Tick(5*time.Second, func(t time.Time) tea.Msg {
		return DashboardTick(t)
	})
}

type (
	SysResReadyMsg struct {
		CPU  *data.CPUInfo
		Mem  *data.MemInfo
		Disk *data.DiskInfo
	}
	DaemonReadyMsg struct {
		Err error
	}
)
