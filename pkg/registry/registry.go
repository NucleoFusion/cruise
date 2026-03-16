package registry

import (
	"github.com/cruise-org/cruise/pkg/types"
	"github.com/zalando/go-keyring"
)

type Registry interface {
	Provider() string // "Harbor", "DockerHub"
	Domain() string   // "harbor.local", "docker.io"

	Login(username, password string) error
	Logout() error

	// Listing
	ListImages() ([]types.RegistryImage, error)
	// TODO: Add these
	// ImageDetails() ([]types.RegistryImage, error)

	// Operations
	// Pull(image types.RegistryImage) error
	// Push(image types.RegistryImage) error
	// Retag(image types.RegistryImage, newTag string) error
}

func IsLoggedIn(user string) bool {
	_, err := keyring.Get("cruise", user)
	return err == nil
}
