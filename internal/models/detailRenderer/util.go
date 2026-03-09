package detailrenderer

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/cruise-org/cruise/internal/messages"
	"github.com/cruise-org/cruise/internal/utils"
	"github.com/cruise-org/cruise/pkg/colors"
	"github.com/cruise-org/cruise/pkg/styles"
)

var (
	KeyStyle = func() lipgloss.Style {
		return lipgloss.NewStyle().Background(colors.Load().HelpKeyBg).Foreground(colors.Load().HelpKeyText)
	}
	ValStyle = func() lipgloss.Style {
		return lipgloss.NewStyle().Foreground(colors.Load().HelpDescText)
	}
)

func SetVP(w, h int, content []string, title string) viewport.Model {
	text := lipgloss.JoinVertical(lipgloss.Left,
		lipgloss.PlaceHorizontal(w-4, lipgloss.Center, styles.TitleStyle().Render(" "+title+" ")),
		"\n\n",
		lipgloss.PlaceVertical(h, lipgloss.Top, strings.Join(content, "\n\n")),
	)

	vp := NewVP(w-1, h)
	vp.SetContent(text)

	return vp
}

func FormatLine(key, val string, w int) string {
	line := fmt.Sprintf("%s %s",
		KeyStyle().Render(fmt.Sprintf(" %s: ", utils.Shorten(key, w/3-3))),
		ValStyle().Render(utils.Shorten(val, 2*w/3-3)),
	)

	return line
}

func (s *DetailRenderer) initRenderer() tea.Cmd {
	return func() tea.Msg {
		vpmap := make(map[string]map[string]string)

		for _, v := range *s.Stats {
			statKey := v.Title()
			log.Println("[RENDERER INIT] " + statKey)

			m, err := v.Stats(context.Background())
			if err != nil {
				*m = map[string]string{"Error:": fmt.Sprintf("Error: %v", err)}
			}
			vpmap[statKey] = *m
		}

		return messages.DetailRendererContent{VPMap: &vpmap}
	}
}

func loadingView(w, h int) *viewport.Model {
	vp := NewVP(w, h)
	vp.SetContent(lipgloss.Place(w, h, lipgloss.Center, lipgloss.Center, "Loading..."))
	return &vp
}

func NewVP(w, h int) viewport.Model {
	style := styles.PageStyle().Padding(1, 2)

	vp := viewport.New(w, h)
	vp.Style = style

	return vp
}
