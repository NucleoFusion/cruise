package msgpopup

import (
	"fmt"

	"github.com/NucleoFusion/cruise/internal/colors"
	"github.com/NucleoFusion/cruise/internal/utils"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type MsgPopup struct {
	Width    int
	Height   int
	Message  string
	Title    string
	Location string
}

func NewMsgPopup(w, h int, msg, title, location string) *MsgPopup {
	return &MsgPopup{
		Width:    w,
		Height:   h,
		Message:  utils.WrapAndLimit(msg, 20, 3),
		Title:    title,
		Location: location,
	}
}

func (s *MsgPopup) Init() tea.Cmd { return nil }

func (s *MsgPopup) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return s, nil
}

func (s *MsgPopup) View() string {
	style := lipgloss.NewStyle()

	text := fmt.Sprintf("%s\n\n%s", style.Foreground(colors.Load().MsgText).Render(s.Title+" | "+s.Location),
		style.Foreground(colors.Load().Text).Render(s.Message))

	return lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(colors.Load().PopupBorder).
		Background(colors.Load().ErrorBg).Padding(1, 3).Render(text)
}
