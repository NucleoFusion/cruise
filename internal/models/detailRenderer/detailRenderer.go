package detailrenderer

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/cruise-org/cruise/internal/messages"
	"github.com/cruise-org/cruise/pkg/types"
)

type DetailRenderer struct {
	Width      int
	Height     int
	StatsFunc  func() ([]types.StatCard, *types.StatMeta)
	RenderFunc func(map[string]map[string]string) string
	Stats      *[]types.StatCard
	Meta       *types.StatMeta
	VPMap      *map[string]map[string]string
	IsLoading  bool
}

func NewDetailRenderer(w, h int,
	statsf func() ([]types.StatCard, *types.StatMeta),
	renderf func(map[string]map[string]string) string,
) *DetailRenderer {
	return &DetailRenderer{
		Width:      w,
		Height:     h,
		StatsFunc:  statsf,
		RenderFunc: renderf,
		IsLoading:  true,
	}
}

func (s *DetailRenderer) Init() tea.Cmd {
	// For async execution
	return func() tea.Msg {
		stats, meta := s.StatsFunc()
		return messages.DetailRendererInit{
			Stats: &stats,
			Meta:  meta,
		}
	}
}

func (s *DetailRenderer) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case messages.DetailRendererInit:
		s.Stats = msg.Stats
		s.Meta = msg.Meta
		log.Printf("Stats: %v", msg.Stats)
		log.Printf("MetaMap: %v", msg.Meta.SpanMap)
		return s, s.initRenderer()
	case messages.DetailRendererContent:
		s.VPMap = msg.VPMap
		log.Printf("VPMap: %v", msg.VPMap)
		s.IsLoading = false
		return s, nil
	}
	return s, nil
}

func (s *DetailRenderer) View() string {
	if s.IsLoading {
		return loadingView(s.Width, s.Height).View()
	}

	return s.RenderFunc(*s.VPMap)
}
