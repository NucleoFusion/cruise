package dockerhub

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/zalando/go-keyring"
)

func (s *DockerHub) Login() (tea.Cmd, error) {
	secret, err := keyring.Get("cruise", fmt.Sprintf("%s/%s", s.Provider(), s.Domain()))
	if err != nil {
		if err == keyring.ErrNotFound {
			return func() tea.Msg {
				return nil
			}
		}
		return nil, err
	}
}
