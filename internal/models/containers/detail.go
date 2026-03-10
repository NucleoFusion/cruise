// SPDX-License-Identifier: Apache-2.0
// Copyright The cruise-org Authors

package containers

import (
	"context"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/cruise-org/cruise/internal/messages"
	detailrenderer "github.com/cruise-org/cruise/internal/models/detailRenderer"
	"github.com/cruise-org/cruise/pkg/runtimes"
	"github.com/cruise-org/cruise/pkg/styles"
	"github.com/cruise-org/cruise/pkg/types"
)

type ContainerDetail struct {
	Width       int
	Height      int
	IsLoading   bool
	Curr        types.Container
	StatDetails *detailrenderer.DetailRenderer
	// TODO : ADD Monitoring
	Monitor *types.Monitor
	Events  []types.Log
}

var tickCmd = func() tea.Cmd {
	return tea.Tick(time.Second*3, func(_ time.Time) tea.Msg {
		return messages.ContainerDetailsTick{}
	})
}

func NewDetail(w int, h int, curr types.Container,
	detailsStatFunc func() func() ([]types.StatCard, *types.StatMeta), // Passing a function that returns a function
	detailsRenderFunc func() func(map[string]map[string]string) string, // Passing a function that returns a function
) *ContainerDetail {
	detailsRenderer := detailrenderer.NewDetailRenderer(w, h, detailsStatFunc(), detailsRenderFunc())

	return &ContainerDetail{
		Width:       w,
		Height:      h,
		Curr:        curr,
		StatDetails: detailsRenderer,
		IsLoading:   true,
	}
}

func (s *ContainerDetail) Init() tea.Cmd {
	return tea.Batch(s.StatDetails.Init(), tickCmd(), func() tea.Msg {
		m, err := runtimes.RuntimeSrv.ContainerLogs(context.Background(), s.Curr.Runtime, s.Curr.ID)
		if err != nil {
			return messages.ErrorMsg{Title: "Could not get container logs", Locn: "Container Details", Msg: err.Error()}
		}

		return messages.ContainerDetailsMonitorReady{Monitor: m}
	})
}

func (s *ContainerDetail) Update(msg tea.Msg) (*ContainerDetail, tea.Cmd) {
	switch msg := msg.(type) {
	case messages.ContainerDetailsMonitorReady:
		s.Monitor = msg.Monitor
		return s, nil

	case messages.ContainerDetailsTick:
		if s.IsLoading {
			return s, tickCmd()
		}

		buf := make([]types.Log, 0)
		for {
			select {
			case v := <-s.Monitor.Incoming:
				buf = append(buf, v)
			default:
				goto done
			}
		}

	done:
		for _, v := range buf {
			s.Events = append(s.Events, v)
			if len(s.Events) == s.Height-5 {
				s.Events = s.Events[1:]
			}
		}

		return s, tickCmd()

	case messages.DetailRendererInit:
		if s.StatDetails == nil {
			return s, nil
		}
		dr, cmd := s.StatDetails.Update(msg)
		if details, ok := dr.(*detailrenderer.DetailRenderer); ok {
			s.StatDetails = details
		}
		return s, cmd

	case messages.DetailRendererContent:
		if s.StatDetails == nil {
			return s, nil
		}

		dr, cmd := s.StatDetails.Update(msg)
		if details, ok := dr.(*detailrenderer.DetailRenderer); ok {
			s.StatDetails = details
		}
		s.IsLoading = false
		return s, cmd
	}
	return s, nil
}

func (s *ContainerDetail) View() string {
	if s.IsLoading || s.StatDetails == nil {
		return "Loading..."
	}

	return lipgloss.JoinVertical(lipgloss.Center,
		s.StatDetails.View(),
		// TODO: Log View
		s.LogView(),
	)
}

func (s *ContainerDetail) LogView() string {
	return styles.SubpageStyle().PaddingTop(1).PaddingLeft(4).Render(lipgloss.JoinVertical(lipgloss.Center,
		styles.TitleStyle().Render("Event Logs"),
		lipgloss.NewStyle().
			PaddingTop(1).
			Width(s.Width-8).     //-6 from padding(4) and border(2)
			Height(s.Height/2-4). //-4 from title(1) border(2) and padding(1)
			Align(lipgloss.Left, lipgloss.Bottom).
			Render(s.FormattedView())))
}

func (s *ContainerDetail) FormattedView() string {
	if s.IsLoading {
		return "Loading Logs....\n"
	}

	if len(s.Events) == 0 {
		return "No Logs yet, waiting....\n"
	}

	text := ""
	events := s.Events
	for _, msg := range events {
		text += runtimes.FormatLog(msg) + "\n"
	}

	return text
}
