// SPDX-License-Identifier: Apache-2.0
// Copyright The cruise-org Authors

package volumes

import (
	"context"
	"log"

	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/lipgloss"
	detailrenderer "github.com/cruise-org/cruise/internal/models/detailRenderer"
	"github.com/cruise-org/cruise/pkg/runtimes"
	"github.com/cruise-org/cruise/pkg/types"
)

func (s *Volumes) detailsStatFunc() func() ([]types.StatCard, *types.StatMeta) {
	curr := s.List.GetCurrentItem()
	return func() (stats []types.StatCard, meta *types.StatMeta) {
		return runtimes.RuntimeSrv.VolumeDetails(context.Background(), curr.Runtime, curr.Name)
	}
}

func (s *Volumes) detailsRenderFunc() func(map[string]map[string]string) string {
	return func(m map[string]map[string]string) string {
		vpmap := map[string]viewport.Model{}
		for k, v := range m {
			log.Println("Volume Detail: " + k)

			w, h := s.findSize(k)

			if len(v) == 0 {
				vpmap[k] = detailrenderer.SetVP(w, h, []string{"No Details Found"}, k)
				continue
			}

			arr := make([]string, 0, len(v))
			for key, val := range v {
				arr = append(arr, detailrenderer.FormatLine(key, val, w))
			}

			vpmap[k] = detailrenderer.SetVP(w, h, arr, k)
		}

		return lipgloss.JoinHorizontal(lipgloss.Center,
			vpmap["Volume Details"].View(),
			vpmap["Labels"].View(),
			vpmap["Options"].View(),
		)
	}
}

func (s *Volumes) findSize(title string) (int, int) {
	switch title {
	case "Volume Details":
		return s.Width / 3, s.Height
	case "Labels":
		return s.Width / 3, s.Height
	case "Options":
		return s.Width / 3, s.Height
	default:
		return 0, 0
	}
}
