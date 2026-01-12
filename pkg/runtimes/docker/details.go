package dockerruntime

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

type ContainerDetails struct {
	ID string
}

func (s *ContainerDetails) Title() string { return "Container Details" }

func (s *ContainerDetails) Stats(ctx context.Context, cli *client.Client) (*map[string]string, error) {
	insp, err := cli.ContainerInspect(ctx, s.ID)
	if err != nil {
		return nil, err
	}

	startedAt, err := time.Parse(time.RFC3339Nano, insp.State.StartedAt)
	if err != nil {
		log.Fatal(err)
	}

	uptime := time.Since(startedAt).String()

	return &map[string]string{
		"ID":            s.ID,
		"Name":          insp.Name,
		"Entrypoint":    strings.Join(insp.Config.Entrypoint, " "),
		"Command":       strings.Join(insp.Config.Cmd, " "),
		"Image":         insp.Image,
		"Status":        string(insp.State.Status),
		"RestartPolicy": string(insp.HostConfig.RestartPolicy.Name),
		"Uptime":        uptime,
	}, nil
}

type Resource struct {
	ID string
}

func (s *Resource) Title() string { return "Resource" }

func (s *Resource) Stats(ctx context.Context, cli *client.Client) (*map[string]string, error) {
	res, err := cli.ContainerStats(ctx, s.ID, false)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var stats container.StatsResponse
	json.Unmarshal(data, &stats)

	r, w := GetBlkio(stats, res.OSType)

	return &map[string]string{
		"CPU":       fmt.Sprintf("%d", stats.CPUStats.CPUUsage.TotalUsage),
		"Memory":    fmt.Sprintf("%d", stats.MemoryStats.Usage),
		"Processes": fmt.Sprintf("%d", stats.NumProcs),
		"Blkio":     fmt.Sprintf("%d Rx / %d Wx", r, w),
	}, nil
}

type Networks struct {
	ID string
}

func (s *Networks) Title() string { return "Networks" }

func (s *Networks) Stats(ctx context.Context, cli *client.Client) (*map[string]string, error) {
	res, err := cli.ContainerInspect(ctx, s.ID)
	if err != nil {
		return nil, err
	}

	nets := []string{}
	for k := range res.NetworkSettings.Networks {
		nets = append(nets, k)
	}

	return &map[string]string{
		"IP":       res.NetworkSettings.IPAddress,
		"Mac":      res.NetworkSettings.MacAddress,
		"Networks": strings.Join(nets, "\n"),
	}, nil
}
