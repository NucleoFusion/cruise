package styles

import (
	"github.com/NucleoFusion/cruise/internal/colors"
	"github.com/charmbracelet/lipgloss"
)

var LogoText = `
   ______              _           
  / ____/_____ __  __ (_)_____ ___ 
 / /    / ___// / / // // ___// _ \
/ /___ / /   / /_/ // /(__  )/  __/
\____//_/    \__,_//_//____/ \___/ 
`

func TextStyle() lipgloss.Style {
	return lipgloss.NewStyle().Foreground(colors.Load().Text)
}
