package detailrenderer

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/cruise-org/cruise/pkg/types"
)

type DetailRenderer struct {
	Width     int
	Height    int
	Stats     *[]types.StatCard
	Meta      *types.StatMeta
	VPMap     *map[string]viewport.Model
	IsLoading bool
}

func NewDetailRenderer(w, h int, stats *[]types.StatCard, meta *types.StatMeta) *DetailRenderer {
	return &DetailRenderer{
		Width:     w,
		Height:    h,
		Stats:     stats,
		Meta:      meta,
		IsLoading: true,
	}
}

func (s *DetailRenderer) Init() tea.Cmd {
	return s.initRenderer()
}

func (s *DetailRenderer) Update(tea.Msg) (tea.Model, tea.Cmd) {
	return s, nil
}

func (s *DetailRenderer) View() string {
	return "Details Here"
}
