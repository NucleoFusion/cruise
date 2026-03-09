// SPDX-License-Identifier: Apache-2.0
// Copyright The cruise-org Authors

package containers

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/cruise-org/cruise/internal/messages"
	detailrenderer "github.com/cruise-org/cruise/internal/models/detailRenderer"
	"github.com/cruise-org/cruise/pkg/types"
)

type ContainerDetail struct {
	Width       int
	Height      int
	IsLoading   bool
	Curr        types.Container
	StatDetails *detailrenderer.DetailRenderer
	// TODO : ADD Monitoring
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
	return tea.Batch(s.StatDetails.Init())
}

func (s *ContainerDetail) Update(msg tea.Msg) (*ContainerDetail, tea.Cmd) {
	switch msg := msg.(type) {
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
	)
}
