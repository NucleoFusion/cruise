package colors

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/cruise-org/cruise/pkg/config"
)

type ColorPalette struct {
	Text             lipgloss.Color
	SubtitleText     lipgloss.Color
	SubtitleBg       lipgloss.Color
	UnfocusedBorder  lipgloss.Color
	FocusedBorder    lipgloss.Color
	HelpKeyBg        lipgloss.Color
	HelpKeyText      lipgloss.Color
	HelpDescText     lipgloss.Color
	MenuSelectedBg   lipgloss.Color
	MenuSelectedText lipgloss.Color
	ErrorText        lipgloss.Color
	ErrorBg          lipgloss.Color
	PopupBorder      lipgloss.Color
	PlaceholderText  lipgloss.Color
	MsgText          lipgloss.Color
}

func Load() ColorPalette {
	s := config.Cfg.Styles
	return ColorPalette{
		UnfocusedBorder:  lipgloss.Color(s.UnfocusedBorder),
		FocusedBorder:    lipgloss.Color(s.FocusedBorder),
		SubtitleText:     lipgloss.Color(s.SubtitleText),
		SubtitleBg:       lipgloss.Color(s.SubtitleBg),
		HelpKeyBg:        lipgloss.Color(s.HelpKeyBg),
		HelpKeyText:      lipgloss.Color(s.HelpKeyText),
		HelpDescText:     lipgloss.Color(s.HelpDescText),
		MenuSelectedBg:   lipgloss.Color(s.MenuSelectedBg),
		MenuSelectedText: lipgloss.Color(s.MenuSelectedText),
		ErrorText:        lipgloss.Color(s.ErrorText),
		ErrorBg:          lipgloss.Color(s.ErrorBg),
		PopupBorder:      lipgloss.Color(s.PopupBorder),
		Text:             lipgloss.Color(s.Text),
		PlaceholderText:  lipgloss.Color(s.PlaceholderText),
		MsgText:          lipgloss.Color(s.MsgText),
	}
}
