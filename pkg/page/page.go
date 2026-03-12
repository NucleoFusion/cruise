// SPDX-License-Identifier: Apache-2.0
// Copyright The cruise-org Authors

package page

import tea "github.com/charmbracelet/bubbletea"

type Page interface {
	Init() tea.Cmd
	Update(msg tea.Msg) (Page, tea.Cmd)
	View() string

	// Called when page needs to clean up routines and more
	// mostly on page switch
	Cleanup()
}
