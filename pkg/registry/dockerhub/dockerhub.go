package dockerhub

import (
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

type DockerHub struct {
	RegistryProvider string
	RegistryDomain   string
	RegistryUsername string
}

func NewDockerHubProvider(domain, username string) *DockerHub {
	return &DockerHub{
		RegistryProvider: "docker",
		RegistryDomain:   domain,
		RegistryUsername: username,
	}
}

func (s *DockerHub) Username() string { return s.RegistryUsername }
func (s *DockerHub) Provider() string { return s.RegistryProvider }
func (s *DockerHub) Domain() string   { return s.RegistryDomain }

func (s *DockerHub) ping(pass string) (string, error) {
	client := resty.New()

	var result struct {
		Token string `json:"token"`
	}

	resp, err := client.R().
		SetBody(map[string]string{
			"username": s.Username(),
			"password": pass,
		}).
		SetResult(&result).
		Post("https://hub.docker.com/v2/users/login")
	if err != nil {
		return "", err
	}
	if resp.StatusCode() != http.StatusOK {
		return "", fmt.Errorf("invalid credentials")
	}

	return result.Token, nil
}
