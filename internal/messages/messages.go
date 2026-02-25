package messages

import (
	"encoding/json"
	"io"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/cruise-org/cruise/internal/data"
	"github.com/cruise-org/cruise/pkg/enums"
	"github.com/cruise-org/cruise/pkg/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/events"
)

// TODO: Remove?
type DashboardTick time.Time

func TickDashboard() tea.Cmd {
	return tea.Tick(5*time.Second, func(t time.Time) tea.Msg {
		return DashboardTick(t)
	})
}

type (
	HomeStatContainer struct{ Containers *[]types.Container }
	HomeStatImage     struct{ Images *[]types.Image }
	HomeStatNetwork   struct{ Networks *[]types.Network }
	HomeStatVolume    struct{ Volumes *[]types.Volume }

	HomeLogsTick    struct{}
	HomeLogsMonitor struct{ Monitor *types.Monitor }

	ChangePg struct {
		Pg     enums.PageType
		Exited bool
	}

	CloseDetails struct{}

	PortMapMsg struct {
		Ports []string
		Err   error
	}

	ErrorMsg struct {
		Title string
		Msg   string
		Locn  string
	}

	CloseError struct{}

	MsgPopup struct {
		Title string
		Msg   string
		Locn  string
	}

	CloseMsgPopup struct{}

	StartScanMsg struct {
		Img string
	}

	ScanResponse struct {
		Arr []any
		Err error
	}

	ScannerListMsg struct {
		Found []bool
	}

	NewContainerDetails struct {
		Stats   container.StatsResponseReader
		Decoder *json.Decoder
		Logs    *io.ReadCloser
	}

	ContainerDetailsReady struct {
		Stats   container.StatsResponseReader
		Decoder *json.Decoder
		Logs    *io.ReadCloser
	}

	ContainerDetailsTick struct{}

	ContainerReadyMsg struct {
		Items *[]types.Container
		Err   error
	}

	ImagesReadyMsg struct {
		Map map[string]types.Image
	}

	UpdateImagesMsg struct {
		Items *[]types.Image
	}

	NetworksReadyMsg struct {
		Items *[]types.Network
	}

	VolumesReadyMsg struct {
		Items *[]types.Volume
	}

	UpdateNetworksMsg struct{}

	FzfSelection struct {
		Selection string
		Exited    bool
	}

	SysResReadyMsg struct {
		CPU  *data.CPUInfo
		Mem  *data.MemInfo
		Disk *data.DiskInfo
	}

	NewEvents struct {
		Events []*events.Message
	}

	DaemonReadyMsg struct {
		Err error
	}
)
