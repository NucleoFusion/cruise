package registrymodel

import tea "github.com/charmbracelet/bubbletea"

type RegistryModel struct {
	Width     int
	Height    int
	IsLoading bool
}

func NewRegistryModel(w, h int) *RegistryModel {
	return &RegistryModel{
		Width:  w,
		Height: h,
	}
}

func (s *RegistryModel) Init() tea.Cmd { return nil }

func (s *RegistryModel) Update(_ tea.Msg) (*RegistryModel, tea.Cmd) { return s, nil }

func (s *RegistryModel) View() string {
	return "Registry Model"
}
