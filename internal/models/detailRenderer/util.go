package detailrenderer

import (
	"context"
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/cruise-org/cruise/internal/messages"
	"github.com/cruise-org/cruise/internal/utils"
	"github.com/cruise-org/cruise/pkg/colors"
	"github.com/cruise-org/cruise/pkg/styles"
	"github.com/cruise-org/cruise/pkg/types"
)

var (
	keyStyle = lipgloss.NewStyle().Background(colors.Load().HelpKeyBg).Foreground(colors.Load().HelpKeyText)
	valStyle = lipgloss.NewStyle().Foreground(colors.Load().HelpDescText)
)

func (s *DetailRenderer) initRenderer() tea.Cmd {
	return func() tea.Msg {
		vpmap := make(map[int]viewport.Model)
		wCol, hRow := sizeCalculator(s.Width, s.Height, s.Meta)

		for _, v := range *s.Stats {
			statKey := v.Title()
			cardMeta := (*s.Meta.SpanMap)[statKey]

			vp := NewVP(cardMeta.Columns*wCol, cardMeta.Rows*hRow)
			vp.SetContent(statView(wCol*cardMeta.Columns, v))

			vpmap[cardMeta.Index] = vp
		}

		return messages.DetailRendererInitialized{VPs: &vpmap}
	}
}

func statView(w int, stat types.StatCard) string {
	var text string

	m, err := stat.Stats(context.Background())
	if err != nil {
		text = utils.Shorten("Error: "+err.Error(), w)
	} else {
		text = convertMapToString(w, m)
	}

	return lipgloss.JoinVertical(lipgloss.Left,
		lipgloss.PlaceHorizontal(w, lipgloss.Center, styles.TitleStyle().Render(" Resource ")), "\n\n", text)
}

// Calculates the size of viewport from metadata for one row / column
func sizeCalculator(w, h int, meta *types.StatMeta) (int, int) {
	return w / meta.TotalColumns, h / meta.TotalRows
}

func convertMapToString(w int, m *map[string]string) string {
	arr := make([]string, 0)
	for k, v := range *m {
		content := fmt.Sprintf("%s %s", keyStyle.Render(fmt.Sprintf(" %s :", k)), valStyle.Render(v))

		arr = append(arr, utils.Shorten(content, w-2))
	}

	return strings.Join(arr, "\n")
}

func NewVP(w, h int) viewport.Model {
	vp := viewport.New(w, h)
	vp.Style = styles.PageStyle().Padding(1, 2)

	return vp
}
