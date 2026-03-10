package containers

import (
	"context"
	"sort"

	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/lipgloss"
	detailrenderer "github.com/cruise-org/cruise/internal/models/detailRenderer"
	"github.com/cruise-org/cruise/pkg/runtimes"
	"github.com/cruise-org/cruise/pkg/types"
)

func (s *Containers) detailsStatFunc() func() ([]types.StatCard, *types.StatMeta) {
	curr := s.List.GetCurrentItem()
	return func() (stats []types.StatCard, meta *types.StatMeta) {
		return runtimes.RuntimeSrv.ContainerDetails(context.Background(), curr.Runtime, curr.Name)
	}
}

func (s *Containers) detailsRenderFunc() func(map[string]map[string]string) string {
	return func(m map[string]map[string]string) string {
		vpmap := map[string]viewport.Model{}
		for k, v := range m {
			w, h := s.findSize(k)

			if len(v) == 0 {
				vpmap[k] = detailrenderer.SetVP(w, h, []string{"No Details Found"}, k)
				continue
			}

			arr := make([]string, 0, len(v))
			for key, val := range v {
				arr = append(arr, detailrenderer.FormatLine(key, val, w))
			}

			sort.Strings(arr)
			vpmap[k] = detailrenderer.SetVP(w, h, arr, k)
		}

		return lipgloss.JoinHorizontal(lipgloss.Center,
			vpmap["Container Details"].View(),
			vpmap["Resources"].View(),
			vpmap["Networks"].View(),
			vpmap["Volumes"].View(),
		)
	}
}

func (s *Containers) findSize(title string) (int, int) {
	switch title {
	case "Container Details":
		return s.Width / 4, s.Height / 2
	case "Resources":
		return s.Width / 4, s.Height / 2
	case "Networks":
		return s.Width / 4, s.Height / 2
	case "Volumes":
		return s.Width / 4, s.Height / 2
	default:
		return 0, 0
	}
}
