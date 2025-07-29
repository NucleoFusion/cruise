package messages

import (
	"time"

	"github.com/NucleoFusion/cruise/internal/data"
	"github.com/NucleoFusion/cruise/internal/docker"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/events"
)

type DashboardTick time.Time

func TickDashboard() tea.Cmd {
	return tea.Tick(5*time.Second, func(t time.Time) tea.Msg {
		return DashboardTick(t)
	})
}

type (
	ContainerReadyMsg struct {
		Items []container.Summary
		Err   error
	}

	FzfSelection struct {
		Selection string
		Exited    bool
	}

	SysResReadyMsg struct {
		CPU  *data.CPUInfo
		Mem  *data.MemInfo
		Disk *data.DiskInfo
	}

	LogsReadyMsg struct {
		Info *docker.EventsInfo
	}

	NewEvents struct {
		Events []*events.Message
	}

	DaemonReadyMsg struct {
		Err error
	}
)
