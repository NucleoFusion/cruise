package containers

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
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
	dvp := viewport.New(w/4, h/2)
	dvp.Style = styles.PageStyle().Padding(1, 2)
	dvp.SetContent(getDashboardView(cnt, insp, w))

	// Resource VP
	rvp := viewport.New(w/4, h/2)
	rvp.Style = styles.PageStyle().Padding(1, 2)
	rvp.SetContent(getResourceView(insp, container.StatsResponse{}, w))

	// Network VP
	nvp := viewport.New(w/4, h/2)
	nvp.Style = styles.PageStyle().Padding(1, 2)
	// ivp.SetContent(getIPAMView(ntw, ipamOpts, w))

	// Volumes VP
	vvp := viewport.New(w/4, h/2)
	vvp.Style = styles.PageStyle().Padding(1, 2)
	// ovp.SetContent(getOptionsView(ntw, opts, w))

	// Logs VP
	lvp := viewport.New(w-2, h/2)
	lvp.Style = styles.PageStyle().Padding(1, 2)
	// ovp.SetContent(getOptionsView(ntw, opts, w))

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

			logs, err := docker.GetContainerLogs(context.Background(), s.Container.ID)
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
		select {
		case line := <-s.LogStreamer.lines:
			if len(line) > s.Width/2 {
				line = line[:s.Width/2-3] + "..."
			}
			s.LogItems = append(s.LogItems, line)
			if len(s.LogItems) > 4 {
				s.LogItems = s.LogItems[len(s.LogItems)-4:]
			}
		default:
			break
		}
	}

	switch msg := msg.(type) {
	case messages.ContainerDetailsTick:
		s.UpdateVP()
		return s, tea.Tick(2*time.Second, func(_ time.Time) tea.Msg { return messages.ContainerDetailsTick{} })
	case messages.ContainerDetailsReady:
		s.IsLoading = false
		s.StatsReader = msg.Stats
		s.Decoder = msg.Decoder

		s.Logs = msg.Logs
		s.LogItems = make([]string, 0)
		s.StartLogStream()

		var m container.StatsResponse
		s.Decoder.Decode(&m)
		s.Stats = m
		s.UpdateVP()

		return s, nil
	}
	s.UpdateVP()
	return s, nil
}

func (s *ContainerDetail) View() string {
	return lipgloss.JoinVertical(lipgloss.Center, lipgloss.JoinHorizontal(lipgloss.Center, s.DashVp.View(), s.ResourceVp.View(), s.NetVp.View(), s.VolVp.View()),
		s.LogsVp.View())
}

func getDashboardView(cnt container.Summary, insp container.InspectResponse, w int) string {
	name := cnt.ID
	if len(cnt.Names) != 0 {
		name = cnt.Names[0]
	}
	created := time.Unix(cnt.Created, 0)
	text := fmt.Sprintf("%s %s \n\n%s %s \n\n%s %s \n\n%s %s \n\n%s %s\n\n%s %s\n\n%s %s",
		styles.DetailKeyStyle().Render(" ID: "), styles.TextStyle().Render(utils.Shorten(cnt.ID, w/4-10)),
		styles.DetailKeyStyle().Render(" Name: "), styles.TextStyle().Render(utils.Shorten(name, w/4-10)),
		styles.DetailKeyStyle().Render(" Command: "), styles.TextStyle().Render(utils.Shorten(cnt.Command, w/4-10)),
		styles.DetailKeyStyle().Render(" Image: "), styles.TextStyle().Render(cnt.Image),
		styles.DetailKeyStyle().Render(" Status: "), styles.TextStyle().Render(utils.Shorten(cnt.Status, w/4-10)),
		styles.DetailKeyStyle().Render(" Restart Policy: "), styles.TextStyle().Render(utils.Shorten(string(insp.HostConfig.RestartPolicy.Name), w/4-10)),
		styles.DetailKeyStyle().Render(" Uptime: "), styles.TextStyle().Render(utils.Shorten(utils.FormatDuration(time.Since(created)), w/4-10)))

	return lipgloss.JoinVertical(lipgloss.Left, lipgloss.PlaceHorizontal(w/4-6, lipgloss.Center, styles.TitleStyle().Render(" Container Details ")), "\n\n", text)
}

func getResourceView(cnt container.InspectResponse, stats container.StatsResponse, w int) string {
	if stats.ID == "" {
		return lipgloss.JoinVertical(lipgloss.Center, lipgloss.PlaceHorizontal(w/4-6, lipgloss.Center, styles.TitleStyle().Render(" Resources ")), "\n\n", "Loading...")
	}

	ps := stats.PidsStats.Current
	if ps == 0 {
		ps = uint64(stats.NumProcs)
	}

	rx, wx := docker.GetBlkio(stats, cnt.Platform)

	text := fmt.Sprintf("%s %s \n\n%s %s \n\n%s %s \n\n%s %s ",
		styles.DetailKeyStyle().Render(" CPU: "), styles.TextStyle().Render(utils.Shorten(fmt.Sprintf("%dms", stats.CPUStats.CPUUsage.TotalUsage/1000000), w/4-10)),
		styles.DetailKeyStyle().Render(" Memory: "), styles.TextStyle().Render(utils.Shorten(fmt.Sprintf("%dMb", stats.MemoryStats.Usage/(1024*1024)), w/4-10)),
		styles.DetailKeyStyle().Render(" Processes: "), styles.TextStyle().Render(utils.Shorten(fmt.Sprintf("%d", ps), w/4-10)),
		styles.DetailKeyStyle().Render(" Blkio: "), styles.TextStyle().Render(fmt.Sprintf("%dRx / %dWx", rx, wx)))

	return lipgloss.JoinVertical(lipgloss.Left, lipgloss.PlaceHorizontal(w/4-6, lipgloss.Center, styles.TitleStyle().Render(" Resource ")), "\n\n", text)
}

func (s *ContainerDetail) UpdateVP() {
	s.ResourceVp.SetContent(getResourceView(s.Insp, s.Stats, s.Width))
}

func (s *ContainerDetail) StartLogStream() {
	// Cancel previous log stream
	if s.LogStreamer != nil {
		s.LogStreamer.cancel()
	}

	ctx, cancel := context.WithCancel(context.Background())
	lines := make(chan string, 100)

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
