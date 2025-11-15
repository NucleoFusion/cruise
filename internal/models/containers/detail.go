package containers

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/NucleoFusion/cruise/internal/docker"
	"github.com/NucleoFusion/cruise/internal/messages"
	"github.com/NucleoFusion/cruise/internal/styles"
	"github.com/NucleoFusion/cruise/internal/utils"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/pkg/stdcopy"
)

type ContainerDetail struct {
	Width       int
	Height      int
	IsLoading   bool
	Container   container.Summary
	Insp        container.InspectResponse
	DashVp      viewport.Model
	LogsVp      viewport.Model
	ResourceVp  viewport.Model
	NetVp       viewport.Model
	VolVp       viewport.Model
	Decoder     *json.Decoder
	Logs        *io.ReadCloser
	LogStreamer *LogStreamer
	LogItems    []string
	StatsReader container.StatsResponseReader
	Stats       container.StatsResponse
}

func NewDetail(w int, h int, cnt container.Summary) *ContainerDetail {
	insp, _ := docker.InspectContainer(cnt.ID)

	// Dash VP
	dvp := viewport.New(w-(w-2)*3/4, h/2)
	dvp.Style = styles.PageStyle().Padding(1, 2)
	dvp.SetContent(getDashboardView(cnt, insp, w-2-(w-2)*3/4-4))

	// Resource VP
	rvp := viewport.New((w-2)/4, h/2)
	rvp.Style = styles.PageStyle().Padding(1, 2)
	rvp.SetContent(getResourceView(insp, container.StatsResponse{}, (w-2)/4-4))

	// Network VP
	nvp := viewport.New((w-2)/4, h/2)
	nvp.Style = styles.PageStyle().Padding(1, 2)
	nvp.SetContent(getNetworksView(insp, (w-2)/4-4))

	// Volumes VP
	vvp := viewport.New((w-2)/4, h/2)
	vvp.Style = styles.PageStyle().Padding(1, 2)
	vvp.SetContent(getVolumeView(insp, (w-2)/4-4))

	// Logs VP
	lvp := viewport.New(w-2, h-h/2)
	lvp.Style = styles.PageStyle().Padding(1, 2)
	lvp.SetContent(getLogsView(make([]string, 0), w-4, h-h/2-2))

	return &ContainerDetail{
		Width:      w,
		Height:     h,
		Container:  cnt,
		Insp:       insp,
		DashVp:     dvp,
		NetVp:      nvp,
		LogsVp:     lvp,
		VolVp:      vvp,
		ResourceVp: rvp,
		LogItems:   make([]string, 0),
		IsLoading:  true,
	}
}

func (s *ContainerDetail) Init() tea.Cmd {
	return tea.Batch(func() tea.Msg { return messages.ContainerDetailsTick{} },
		tea.Tick(0, func(_ time.Time) tea.Msg {
			stats, err := docker.GetContainerStats(s.Container.ID)
			if err != nil {
				return utils.ReturnError("Containers Page", "Error Querying Stats", err)
			}
			decoder := json.NewDecoder(stats.Body)

			logs, err := docker.GetContainerLogs(context.Background(), s.Container.ID, 20)
			if err != nil {
				return utils.ReturnError("Containers Page", "Error Querying Logs", err)
			}

			return messages.ContainerDetailsReady{
				Stats:   stats,
				Decoder: decoder,
				Logs:    &logs,
			}
		}))
}

func (s *ContainerDetail) Update(msg tea.Msg) (*ContainerDetail, tea.Cmd) {
	if s.LogStreamer != nil {
	drain: // named scope to return to
		for {
			select {
			case line := <-s.LogStreamer.lines:
				s.LogItems = append(s.LogItems, utils.Shorten(line, s.Width))
				if s.LogsVp.YOffset+s.LogsVp.Height < len(s.LogItems)-2 {
					s.LogsVp.YOffset += 1
				}
			default:
				break drain // returning to named scope
			}
		}
	}

	switch msg := msg.(type) {
	case messages.ContainerDetailsTick:
		if !s.IsLoading {
			var st container.StatsResponse
			s.Decoder.Decode(&st)
			s.Stats = st
			s.UpdateVP()
		}
		return s, tea.Tick(2*time.Second, func(_ time.Time) tea.Msg { return messages.ContainerDetailsTick{} })
	case messages.ContainerDetailsReady:
		s.StatsReader = msg.Stats
		s.Decoder = msg.Decoder
		s.IsLoading = false

		s.Logs = msg.Logs
		s.LogItems = make([]string, 0)
		s.StartLogStream()

		var st container.StatsResponse
		s.Decoder.Decode(&st)
		s.Stats = st

		s.UpdateVP()

		return s, nil
	}
	s.UpdateVP()
	return s, nil
}

func (s *ContainerDetail) View() string {
	return lipgloss.JoinVertical(lipgloss.Center,
		lipgloss.JoinHorizontal(lipgloss.Center, s.DashVp.View(), s.ResourceVp.View(), s.NetVp.View(), s.VolVp.View()),
		s.LogsVp.View())
}

func getDashboardView(cnt container.Summary, insp container.InspectResponse, w int) string {
	name := cnt.ID
	if len(cnt.Names) != 0 {
		name = cnt.Names[0]
	}
	created := time.Unix(cnt.Created, 0)
	text := fmt.Sprintf("%s %s \n\n%s %s \n\n%s %s \n\n%s %s \n\n%s %s\n\n%s %s\n\n%s %s",
		styles.DetailKeyStyle().Render(" ID: "), styles.TextStyle().Render(utils.Shorten(cnt.ID, w-5)),
		styles.DetailKeyStyle().Render(" Name: "), styles.TextStyle().Render(utils.Shorten(name, w-5)),
		styles.DetailKeyStyle().Render(" Command: "), styles.TextStyle().Render(utils.Shorten(cnt.Command, w-5)),
		styles.DetailKeyStyle().Render(" Image: "), styles.TextStyle().Render(cnt.Image),
		styles.DetailKeyStyle().Render(" Status: "), styles.TextStyle().Render(utils.Shorten(cnt.Status, w-5)),
		styles.DetailKeyStyle().Render(" Restart Policy: "), styles.TextStyle().Render(utils.Shorten(string(insp.HostConfig.RestartPolicy.Name), w-5)),
		styles.DetailKeyStyle().Render(" Uptime: "), styles.TextStyle().Render(utils.Shorten(utils.FormatDuration(time.Since(created)), w-5)))

	return lipgloss.JoinVertical(lipgloss.Left, lipgloss.PlaceHorizontal(w, lipgloss.Center, styles.TitleStyle().Render(" Container Details ")), "\n\n", text)
}

func getResourceView(cnt container.InspectResponse, stats container.StatsResponse, w int) string {
	if stats.ID == "" {
		return lipgloss.JoinVertical(lipgloss.Center, lipgloss.PlaceHorizontal(w, lipgloss.Center, styles.TitleStyle().Render(" Resources ")), "\n\n", "Loading...")
	}

	ps := stats.PidsStats.Current
	if ps == 0 {
		ps = uint64(stats.NumProcs)
	}

	rx, wx := docker.GetBlkio(stats, cnt.Platform)

	text := fmt.Sprintf("%s %s \n\n%s %s \n\n%s %s \n\n%s %s ",
		styles.DetailKeyStyle().Render(" CPU: "), styles.TextStyle().Render(utils.Shorten(fmt.Sprintf("%dms", stats.CPUStats.CPUUsage.TotalUsage/500000), w-5)),
		styles.DetailKeyStyle().Render(" Memory: "), styles.TextStyle().Render(utils.Shorten(fmt.Sprintf("%dMb", stats.MemoryStats.Usage/(524*524)), w-5)),
		styles.DetailKeyStyle().Render(" Processes: "), styles.TextStyle().Render(utils.Shorten(fmt.Sprintf("%d", ps), w-5)),
		styles.DetailKeyStyle().Render(" Blkio: "), styles.TextStyle().Render(utils.Shorten(fmt.Sprintf("%dRx / %dWx", rx, wx), w-5)))

	return lipgloss.JoinVertical(lipgloss.Left, lipgloss.PlaceHorizontal(w, lipgloss.Center, styles.TitleStyle().Render(" Resource ")), "\n\n", text)
}

func getNetworksView(insp container.InspectResponse, w int) string {
	ports := "\n"
	for k, v := range insp.NetworkSettings.Ports {
		host := make([]string, 0)
		for _, pb := range v {
			host = append(host, fmt.Sprintf("%s:%s", pb.HostIP, pb.HostPort))
		}
		ports += fmt.Sprintf("\n%s <-> %s", string(k), strings.Join(host, " ~ "))
	}

	nets := "\n"
	for k, v := range insp.NetworkSettings.Networks {
		nets += fmt.Sprintf("\n%s ~ %s", k, v.IPAddress)
	}
	text := fmt.Sprintf("%s %s \n\n%s %s \n\n%s %s ",
		styles.DetailKeyStyle().Render(" Name: "), styles.TextStyle().Render(utils.Shorten(insp.HostConfig.NetworkMode.NetworkName(), w-5)),
		styles.DetailKeyStyle().Render(" Networks: "), styles.TextStyle().Render(nets),
		styles.DetailKeyStyle().Render(" Ports:- "), styles.TextStyle().Render(ports))

	return lipgloss.JoinVertical(lipgloss.Left, lipgloss.PlaceHorizontal(w, lipgloss.Center, styles.TitleStyle().Render(" Networks ")), "\n\n", text)
}

func getVolumeView(cnt container.InspectResponse, w int) string {
	if len(cnt.Mounts) == 0 {
		return lipgloss.JoinVertical(lipgloss.Center, lipgloss.PlaceHorizontal(w, lipgloss.Center, styles.TitleStyle().Render(" Volume ")), "\n\n", "No Mounts Found")
	}
	text := fmt.Sprintf("%s %s \n\n%s %s \n\n%s %s \n\n%s %s \n\n%s %s ",
		styles.DetailKeyStyle().Render(" Name: "), styles.TextStyle().Render(utils.Shorten(cnt.Mounts[0].Name, w-5)),
		styles.DetailKeyStyle().Render(" Dest: "), styles.TextStyle().Render(utils.Shorten(cnt.Mounts[0].Destination, w-5)),
		styles.DetailKeyStyle().Render(" Source: "), styles.TextStyle().Render(utils.Shorten(cnt.Mounts[0].Source, w-5)),
		styles.DetailKeyStyle().Render(" Type: "), styles.TextStyle().Render(utils.Shorten(string(cnt.Mounts[0].Type), w-5)),
		styles.DetailKeyStyle().Render(" Mode: "), styles.TextStyle().Render(cnt.Mounts[0].Mode))

	return lipgloss.JoinVertical(lipgloss.Left, lipgloss.PlaceHorizontal(w, lipgloss.Center, styles.TitleStyle().Render(" Volume ")), "\n\n", text)
}

func getLogsView(items []string, w int, h int) string {
	if len(items) == 0 {
		return lipgloss.JoinVertical(lipgloss.Center, lipgloss.PlaceHorizontal(w, lipgloss.Center, styles.TitleStyle().Render(" Logs ")), "\n\n", "No Logs yet...")
	}

	its := items
	if len(items) > h-7 { // Accounted for padding and header
		its = items[len(items)-(h-7):]
	}

	return lipgloss.JoinVertical(lipgloss.Left, lipgloss.PlaceHorizontal(w, lipgloss.Center, styles.TitleStyle().Render(" Logs ")), "\n\n", strings.Join(its, "\n"))
}

func (s *ContainerDetail) UpdateVP() {
	if s.IsLoading {
		return
	}
	s.ResourceVp.SetContent(getResourceView(s.Insp, s.Stats, (s.Width-2)/4-4))
	s.LogsVp.SetContent(getLogsView(s.LogItems, s.Width-4, s.LogsVp.Height))
}

func (s *ContainerDetail) StartLogStream() {
	// Cancel previous log stream
	if s.LogStreamer != nil {
		s.LogStreamer.cancel()
	}

	ctx, cancel := context.WithCancel(context.Background())
	lines := make(chan string, 50)

	s.LogStreamer = &LogStreamer{
		ctx:    ctx,
		cancel: cancel,
		lines:  lines,
	}

	// Something about demuxing the logs
	stdoutReader, stdoutWriter := io.Pipe()
	stderrReader, stderrWriter := io.Pipe()

	go func() {
		defer stdoutWriter.Close()
		defer stderrWriter.Close()
		_, err := stdcopy.StdCopy(stdoutWriter, stderrWriter, *s.Logs)
		if err != nil {
			fmt.Println("StdCopy Error:", err)
		}
	}()

	go func() {
		logs := s.Logs
		defer (*logs).Close()

		scanner := bufio.NewScanner(stdoutReader)

		for {
			select {
			case <-ctx.Done():
				return
			default:
				if scanner.Scan() {
					line := scanner.Text()
					if line != "" {
						lines <- line
					}
				}
			}
		}
	}()

	go func() {
		logs := s.Logs
		defer (*logs).Close()

		scanner := bufio.NewScanner(stderrReader)

		for {
			select {
			case <-ctx.Done():
				return
			default:
				if scanner.Scan() {
					line := scanner.Text()
					if line != "" {
						lines <- line
					}
				}
			}
		}
	}()
}
