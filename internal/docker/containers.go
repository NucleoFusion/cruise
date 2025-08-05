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

func GetNumContainers() int {
	containers, err := cli.ContainerList(context.Background(), container.ListOptions{All: true})
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return -1
	}

	return len(containers)
}

func GetContainers() ([]container.Summary, error) {
	containers, err := cli.ContainerList(context.Background(), container.ListOptions{All: true})
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return containers, err
	}

	return containers, nil
}

// Gives a realtime updater
func GetContainerStats(id string) (container.StatsResponseReader, error) {
	return cli.ContainerStats(context.Background(), id, true)
}

// Stream of logs
func GetContainerLogs(ctx context.Context, id string) (io.ReadCloser, error) {
	return cli.ContainerLogs(ctx, id, container.LogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Tail:       "4",
		Follow:     true,
		Timestamps: false,
		Details:    false,
	})
}

func FormattedSummary(item container.Summary, width int) string {
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

func SummaryHeaders(width int) string {
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
