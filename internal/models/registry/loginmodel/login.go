package loginmodel

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/cruise-org/cruise/internal/messages"
	"github.com/cruise-org/cruise/internal/utils"
	"github.com/cruise-org/cruise/pkg/colors"
	"github.com/cruise-org/cruise/pkg/registry"
	"github.com/cruise-org/cruise/pkg/styles"
)

type LoginModel struct {
	Width    int
	Registry registry.Registry
	PassTi   textinput.Model
}

func NewLoginModel(w int, r registry.Registry) *LoginModel {
	ti := textinput.New()
	ti.EchoMode = textinput.EchoPassword
	ti.CharLimit = w - 4 - 11 // Taking into account border and 'Password : '
	ti.Focus()

	return &LoginModel{
		Width:    w,
		Registry: r,
		PassTi:   ti,
	}
}

func (s *LoginModel) Init() tea.Cmd { return nil }

func (s *LoginModel) Update(msg tea.Msg) (*LoginModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEscape:
			return s, func() tea.Msg {
				return messages.CloseLoginMessage{}
			}
		default:
			ti, cmd := s.PassTi.Update(msg)
			s.PassTi = ti
			return s, cmd
		}
	}
	return s, nil
}

func (s *LoginModel) View() string {
	view := lipgloss.JoinVertical(lipgloss.Left,
		fmt.Sprintf("Username : %s", s.Registry.Username()),
		fmt.Sprintf("Provider : %s", s.Registry.Provider()),
		fmt.Sprintf("Domain   : %s", utils.Shorten(s.Registry.Domain(), 30)),
		lipgloss.JoinHorizontal(lipgloss.Left,
			"Password : ", s.PassTi.View(),
		),
	)
	style := lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(colors.Load().FocusedBorder).
		Padding(1, 2)

	return lipgloss.JoinVertical(lipgloss.Center,
		styles.TitleStyle().Render("Login"),
		style.Render(lipgloss.PlaceHorizontal(s.Width, lipgloss.Left, view)),
	)
}
