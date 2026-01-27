// SPDX-License-Identifier: Apache-2.0
// Copyright The cruise-org Authors

package styles

import "github.com/charmbracelet/lipgloss"

func DropdownBorder() lipgloss.Border {
	b := lipgloss.RoundedBorder()

	b.Top = " "
	b.TopLeft = b.Left   // "|"
	b.TopRight = b.Right // "|"

	return b
}
