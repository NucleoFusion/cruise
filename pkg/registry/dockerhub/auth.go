package dockerhub

import (
	"fmt"

	"github.com/zalando/go-keyring"
)

func (s *DockerHub) Login(pass string) error {
	token, err := s.ping(pass)
	if err != nil {
		return err
	}

	return keyring.Set("cruise", fmt.Sprintf("%s/%s", s.Provider(), s.Domain()), token)
}

func (s *DockerHub) Logout() error {
	return keyring.Delete("cruise", fmt.Sprintf("%s/%s", s.Provider(), s.Domain()))
}
