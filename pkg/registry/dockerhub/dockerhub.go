package dockerhub

type DockerHub struct {
	RegistryProvider string
	RegistryDomain   string
	Username         string
}

func NewDockerHubProvider(domain, username string) *DockerHub {
	return &DockerHub{
		RegistryProvider: "docker",
		RegistryDomain:   domain,
		Username:         username,
	}
}

func (s *DockerHub) Provider() string { return s.RegistryProvider }
func (s *DockerHub) Domain() string   { return s.RegistryDomain }
