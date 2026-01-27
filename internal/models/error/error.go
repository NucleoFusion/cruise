package errorpopup

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/cruise-org/cruise/internal/utils"
	"github.com/cruise-org/cruise/pkg/colors"
)

type ErrorPopup struct {
	Width    int
	Height   int
	Message  string
	Title    string
	Location string
}

func NewErrorPopup(w, h int, msg, title, location string) *ErrorPopup {
	return &ErrorPopup{
		Width:    w,
		Height:   h,
		Message:  utils.WrapAndLimit(msg, 20, 3),
		Title:    title,
		Location: location,
	}
}

func (s *ErrorPopup) Init() tea.Cmd { return nil }

func (s *ErrorPopup) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return s, nil
}

func (s *ErrorPopup) View() string {
	style := lipgloss.NewStyle()

	text := fmt.Sprintf("%s\n\n%s", style.Foreground(colors.Load().ErrorText).Render(s.Title+" | "+s.Location),
		style.Foreground(colors.Load().Text).Render(s.Message))

	return lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(colors.Load().PopupBorder).
		Background(colors.Load().ErrorBg).Padding(1, 3).Render(text)
}
