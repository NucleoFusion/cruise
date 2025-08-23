package docker

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/NucleoFusion/cruise/internal/utils"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/events"
)

type Details struct {
	IsLoading   bool
	CPU         float64
	Mem         float64
	Size        float64
	Logs        []string
	EventStream chan *events.Message
}

func GetPorts() ([]string, error) {
	conts, err := GetContainers()
	if err != nil {
		return nil, err
	}

	arr := make([]string, 0)
	for _, cnt := range conts {
		inp, err := cli.ContainerInspect(context.Background(), cnt.ID)
		if err != nil {
			return nil, err
		}

		for port, bindings := range inp.NetworkSettings.Ports {
			for _, b := range bindings {
				arr = append(arr, fmt.Sprintf("%s --> %s:%s | %s", port, b.HostIP, b.HostPort, cnt.Names[0]))
			}
		}
	}

	return arr, nil
}

func InspectContainer(id string) (container.InspectResponse, error) {
	return cli.ContainerInspect(context.Background(), id)
}

func GetNumContainers() int {
	containers, err := cli.ContainerList(context.Background(), container.ListOptions{All: true})
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return -1
	}

	return len(containers)
}

func GetContainers() ([]container.Summary, error) {
	return cli.ContainerList(context.Background(), container.ListOptions{All: true})
}

// Gives a realtime updater
func GetContainerStats(id string) (container.StatsResponseReader, error) {
	return cli.ContainerStats(context.Background(), id, true)
}

// Stream of logs
func GetContainerLogs(ctx context.Context, id string, tail int) (io.ReadCloser, error) {
	return cli.ContainerLogs(ctx, id, container.LogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Tail:       fmt.Sprintf("%d", tail),
		Follow:     true,
		Timestamps: false,
		Details:    false,
	})
}

func ContainerFormattedSummary(item container.Summary, width int) string {
	name := "-"
	if len(item.Names) > 0 {
		name = strings.TrimPrefix(item.Names[0], "/")
	}

	format := strings.Repeat(fmt.Sprintf("%%-%ds ", width), 9)

	return fmt.Sprintf(
		format,
		utils.ShortID(item.ID),
		utils.Shorten(name, width),
		utils.Shorten(item.Image, width),
		utils.CreatedAgo(item.Created),
		utils.Shorten(utils.FormatPorts(item.Ports), width),
		item.State,
		utils.Shorten(item.Status, width),
		utils.Shorten(utils.FormatMounts(item.Mounts), width),
		utils.FormatSize(item.SizeRootFs),
	)
}

func ContainerHeaders(width int) string {
	format := strings.Repeat(fmt.Sprintf("%%-%ds ", width), 9)

	return fmt.Sprintf(
		format,
		"ID",
		"Name",
		"Image",
		"Created",
		"Ports",
		"State",
		"Status",
		"Mounts",
		"Size",
	)
}

func GetBlkio(stats container.StatsResponse, platform string) (uint64, uint64) {
	switch platform {
	case "windows":
		return stats.StorageStats.ReadSizeBytes, stats.StorageStats.WriteSizeBytes
	case "linux":
		var read, write uint64
		for _, entry := range stats.BlkioStats.IoServiceBytesRecursive {
			switch entry.Op {
			case "Read":
				read += entry.Value
			case "Write":
				write += entry.Value
			}
		}
		return read, write
	}
	return 0, 0
}

func StartContainer(ID string) error {
	err := cli.ContainerStart(context.Background(), ID, container.StartOptions{})
	if err != nil {
		return err
	}

	return nil
}

func RestartContainer(ID string) error {
	err := cli.ContainerRestart(context.Background(), ID, container.StopOptions{})
	if err != nil {
		return err
	}

	return nil
}

func RemoveContainer(ID string) error {
	err := cli.ContainerRemove(context.Background(), ID, container.RemoveOptions{})
	if err != nil {
		return err
	}

	return nil
}

func PauseContainer(ID string) error {
	err := cli.ContainerPause(context.Background(), ID)
	if err != nil {
		return err
	}

	return nil
}

func UnpauseContainer(ID string) error {
	err := cli.ContainerUnpause(context.Background(), ID)
	if err != nil {
		return err
	}

	return nil
}

func StopContainer(ID string) error {
	err := cli.ContainerStop(context.Background(), ID, container.StopOptions{})
	if err != nil {
		return err
	}

	return nil
}
