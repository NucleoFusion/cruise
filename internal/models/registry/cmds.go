package registrymodel

import (
	"fmt"
	"sync"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/cruise-org/cruise/internal/messages"
	"github.com/cruise-org/cruise/pkg/config"
	"github.com/cruise-org/cruise/pkg/registry"
)

func (s *RegistryModel) parseRegistries() tea.Cmd {
	return func() tea.Msg {
		cfgRegistries := config.Cfg.Global.Registry
		registries := make([]registry.Registry, 0)
		for _, v := range cfgRegistries {
			if v.Ignore == true {
				continue
			}

			r, err := registry.GetRegistry(&v)
			if err != nil {
				continue
			}

			registries = append(registries, r)
		}

		return messages.ParsedRegistries{Registries: registries}
	}
}

func (s *RegistryModel) authenticateRegistries() tea.Cmd {
	return func() tea.Msg {
		var wg sync.WaitGroup
		ch := make(chan messages.RegistryLoginMessage)

		for _, v := range s.Registries {
			wg.Add(1)
			go func() {
				defer wg.Done()
				if !registry.IsLoggedIn(fmt.Sprintf("%s/%s", v.Provider(), v.Domain())) {
					ch <- messages.RegistryLoginMessage{Registry: v}
				}
			}()
		}

		wg.Wait()
		close(ch)

		return messages.PendingRegistryLogin{Ch: ch}
	}
}
