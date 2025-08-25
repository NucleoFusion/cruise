package styledhelp

import (
	"fmt"
	"strings"

	"github.com/NucleoFusion/cruise/internal/colors"
	"github.com/NucleoFusion/cruise/internal/keymap"
	"github.com/NucleoFusion/cruise/internal/styles"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/lipgloss"
)

type StyledHelp struct {
	Width      int
	styledhelp help.Model
	kmap       help.KeyMap
}

func NewStyledHelp(b []key.Binding, w int) StyledHelp {
	h := help.New()
	h.Styles.FullKey = h.Styles.FullKey.Margin(0, 0).Padding(0, 0)
	h.Styles.FullDesc = h.Styles.FullDesc.Margin(0, 0).Padding(0, 0)
	h.Styles.ShortKey = lipgloss.NewStyle().Background(colors.Load().HelpKeyText).Foreground(colors.Load().Text)
	h.Styles.ShortDesc = lipgloss.NewStyle().Foreground(colors.Load().HelpDescText)

	newk := make([]key.Binding, 0, len(b))
	newk = append(newk, key.NewBinding(key.WithKeys("tab"), key.WithHelp(" tab ", "switch pages")))
	for _, bind := range b {
		newk = append(newk, padBinding(bind))
	}

	return StyledHelp{
		Width:      w - 4,
		styledhelp: h,
		kmap:       keymap.NewDynamic(newk),
	}
}

func (s *StyledHelp) View() string {
	return styles.SubpageStyle().Render(lipgloss.Place(s.Width, 1, lipgloss.Center, lipgloss.Center,
		strings.Trim(s.styledhelp.View(s.kmap), "\n")))
}

func padBinding(b key.Binding) key.Binding {
	k := fmt.Sprintf(" %s ", b.Help().Key) // pad to fixed width
	return key.NewBinding(
		key.WithKeys(b.Keys()...),
		key.WithHelp(k, b.Help().Desc),
	)
}
