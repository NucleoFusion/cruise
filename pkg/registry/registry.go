package registry

import "github.com/cruise-org/cruise/pkg/types"

type Registry interface {
	Provider() string // "Harbor", "DockerHub"
	Domain() string   // "harbor.local", "docker.io"

	Login(username, password string) error
	Logout() error
	Ping() error // for validate existing token

	// Listing
	ListImages() ([]types.RegistryImage, error)

	// Operations
	Pull(image types.RegistryImage) error
	Push(image types.RegistryImage) error
	Retag(image types.RegistryImage, newTag string) error
}
