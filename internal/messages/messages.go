package messages

import (
	"encoding/json"
	"io"
	"time"

	"github.com/NucleoFusion/cruise/internal/data"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/events"
	"github.com/docker/docker/api/types/image"
)

type DashboardTick time.Time

func TickDashboard() tea.Cmd {
	return tea.Tick(5*time.Second, func(t time.Time) tea.Msg {
		return DashboardTick(t)
	})
}

type (
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

	ContainerReadyMsg struct {
		Items []container.Summary
		Err   error
	}

	ImagesReadyMsg struct {
		Items []image.Summary
	}

	UpdateImagesMsg struct{}

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
