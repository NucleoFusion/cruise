package colors

import "github.com/charmbracelet/lipgloss"

type ColorPalette struct {
	// Background layers
	Crust  lipgloss.Color // darkest
	Mantle lipgloss.Color
	Base   lipgloss.Color // regular background

	// Surface & overlays
	Surface0 lipgloss.Color
	Surface1 lipgloss.Color
	Surface2 lipgloss.Color
	Overlay0 lipgloss.Color
	Overlay1 lipgloss.Color
	Overlay2 lipgloss.Color

	// Text layers
	Subtext0 lipgloss.Color
	Subtext1 lipgloss.Color
	Text     lipgloss.Color

	// Accent/Funky colors
	Rosewater lipgloss.Color
	Flamingo  lipgloss.Color
	Pink      lipgloss.Color
	Mauve     lipgloss.Color
	Red       lipgloss.Color
	Maroon    lipgloss.Color
	Peach     lipgloss.Color
	Yellow    lipgloss.Color
	Green     lipgloss.Color
	Teal      lipgloss.Color
	Sky       lipgloss.Color
	Sapphire  lipgloss.Color
	Blue      lipgloss.Color
	Lavender  lipgloss.Color
}

func Load() ColorPalette {
	return ColorPalette{
		Crust:  lipgloss.Color("#11111b"),
		Mantle: lipgloss.Color("#181825"),
		Base:   lipgloss.Color("#1e1e2e"),

		Surface0: lipgloss.Color("#313244"),
		Surface1: lipgloss.Color("#45475a"),
		Surface2: lipgloss.Color("#585b70"),
		Overlay0: lipgloss.Color("#6c7086"),
		Overlay1: lipgloss.Color("#7f849c"),
		Overlay2: lipgloss.Color("#9399b2"),

		Subtext0: lipgloss.Color("#a6adc8"),
		Subtext1: lipgloss.Color("#bac2de"),
		Text:     lipgloss.Color("#cdd6f4"),

		Rosewater: lipgloss.Color("#f5e0dc"),
		Flamingo:  lipgloss.Color("#f2cdcd"),
		Pink:      lipgloss.Color("#f5c2e7"),
		Mauve:     lipgloss.Color("#cba6f7"),
		Red:       lipgloss.Color("#f38ba8"),
		Maroon:    lipgloss.Color("#eba0ac"),
		Peach:     lipgloss.Color("#f9e2af"),
		Yellow:    lipgloss.Color("#f5e0dc"),
		Green:     lipgloss.Color("#a6e3a1"),
		Teal:      lipgloss.Color("#94e2d5"),
		Sky:       lipgloss.Color("#74c7ec"),
		Sapphire:  lipgloss.Color("#89b4fa"),
		Blue:      lipgloss.Color("#8aadf4"),
		Lavender:  lipgloss.Color("#b4befe"),
	}
}
