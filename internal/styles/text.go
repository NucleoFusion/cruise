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

var ContainersText = `
   ______               __          _                         
  / ____/____   ____   / /_ ____ _ (_)____   ___   _____ _____
 / /    / __ \ / __ \ / __// __ '// // __ \ / _ \ / ___// ___/
/ /___ / /_/ // / / // /_ / /_/ // // / / //  __// /   (__  ) 
\____/ \____//_/ /_/ \__/ \__,_//_//_/ /_/ \___//_/   /____/  `

var ImagesText = `
    ____                               
   /  _/___ ___  ____ _____ ____  _____
   / // __ '__ \/ __ '/ __ '/ _ \/ ___/
 _/ // / / / / / /_/ / /_/ /  __(__  ) 
/___/_/ /_/ /_/\__,_/\__, /\___/____/  
                    /____/             `

func TextStyle() lipgloss.Style {
	return lipgloss.NewStyle().Foreground(colors.Load().Text)
}
