package registrymodel

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/cruise-org/cruise/internal/messages"
	"github.com/cruise-org/cruise/pkg/page"
	"github.com/cruise-org/cruise/pkg/registry"
)

type RegistryModel struct {
	Width  int
	Height int
	State  int // Defines current phase in initialization, check View() for exact
	// Data
	Registries []registry.Registry
}

func NewRegistryModel(w, h int) *RegistryModel {
	return &RegistryModel{
		Width:  w,
		Height: h,
		State:  1,
	}
}

func (s *RegistryModel) Cleanup() {}

func (s *RegistryModel) Init() tea.Cmd { return s.parseRegistries() }

func (s *RegistryModel) Update(msg tea.Msg) (page.Page, tea.Cmd) {
	switch msg := msg.(type) {
	case messages.ParsedRegistries:
		s.Registries = msg.Registries
		s.State = 2
		return s, s.authenticateRegistries()
		// TODO: Handle Rest of the Auth
	}
	return s, nil
}

func (s *RegistryModel) View() string {
	switch s.State {
	case 1:
		return "Parsing Registries..."
	case 2:
		return "Checking Authentication Status..."
	case 3:
		return "Logging In..."
	}

	return "Registry Model"
}
