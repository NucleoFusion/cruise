package styledhelp

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/lipgloss"
	"github.com/cruise-org/cruise/pkg/colors"
	"github.com/cruise-org/cruise/pkg/keymap"
	"github.com/cruise-org/cruise/pkg/styles"
)

type StyledHelp struct {
	Width      int
	styledhelp help.Model
	kmap       help.KeyMap
	vp         viewport.Model
}

func NewStyledHelp(b []key.Binding, w int) StyledHelp {
	h := help.New()
	h.Styles.FullKey = h.Styles.FullKey.Margin(0, 0).Padding(0, 0)
	h.Styles.FullDesc = h.Styles.FullDesc.Margin(0, 0).Padding(0, 0)
	h.Styles.ShortKey = lipgloss.NewStyle().Background(colors.Load().HelpKeyBg).Foreground(colors.Load().HelpKeyText)
	h.Styles.ShortDesc = lipgloss.NewStyle().Foreground(colors.Load().HelpDescText)

	newk := make([]key.Binding, 0, len(b))
	newk = append(newk, key.NewBinding(key.WithKeys("tab"), key.WithHelp(" tab ", "switch pages")), keymap.QuickQuitKey())
	for _, bind := range b {
		newk = append(newk, padBinding(bind))
	}

	vp := viewport.New(w, 3)
	vp.Style = styles.SubpageStyle()

	return StyledHelp{
		Width:      w - 4,
		styledhelp: h,
		kmap:       keymap.NewDynamic(newk),
		vp:         vp,
	}
}

func (s *StyledHelp) View() string {
	s.vp.SetContent(lipgloss.PlaceHorizontal(s.Width-2, lipgloss.Center, strings.Trim(s.styledhelp.View(s.kmap), "\n")))
	return s.vp.View()
}

func padBinding(b key.Binding) key.Binding {
	k := fmt.Sprintf(" %s ", b.Help().Key) // pad to fixed width
	return key.NewBinding(
		key.WithKeys(b.Keys()...),
		key.WithHelp(k, b.Help().Desc),
	)
}
