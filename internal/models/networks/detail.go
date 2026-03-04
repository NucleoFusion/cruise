// SPDX-License-Identifier: Apache-2.0
// Copyright The cruise-org Authors

package networks

import (
	"context"

	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/lipgloss"
	detailrenderer "github.com/cruise-org/cruise/internal/models/detailRenderer"
	"github.com/cruise-org/cruise/pkg/runtimes"
	"github.com/cruise-org/cruise/pkg/types"
)

func (s *Networks) detailsStatFunc() func() ([]types.StatCard, *types.StatMeta) {
	curr := s.List.GetCurrentItem()
	return func() (stats []types.StatCard, meta *types.StatMeta) {
		return runtimes.RuntimeSrv.NetworkDetails(context.Background(), curr.Runtime, curr.ID)
	}
}

func (s *Networks) detailsRenderFunc() func(map[string]map[string]string) string {
	return func(m map[string]map[string]string) string {
		vpmap := map[string]viewport.Model{}
		for k, v := range m {
			w, h := s.findSize(k)

			arr := make([]string, 0, len(v))
			for key, val := range v {
				arr = append(arr, detailrenderer.FormatLine(key, val, w))
			}

			vpmap[k] = detailrenderer.SetVP(w, h, arr, k)
		}

		return lipgloss.JoinHorizontal(lipgloss.Center,
			lipgloss.JoinVertical(lipgloss.Center, vpmap["Network Details"].View(), vpmap["IPAM"].View()),
			vpmap["Labels"].View(),
			vpmap["Options"].View(),
		)
	}
}

func (s *Networks) findSize(title string) (int, int) {
	switch title {
	case "IPAM":
		return s.Width / 3, s.Height / 2
	case "Network Details":
		return s.Width / 3, s.Height / 2
	case "Labels":
		return s.Width / 3, s.Height
	case "Options":
		return s.Width / 3, s.Height
	default:
		return 0, 0
	}
}
