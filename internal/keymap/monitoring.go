package keymap

import (
	"reflect"

	"github.com/NucleoFusion/cruise/internal/config"
	"github.com/charmbracelet/bubbles/key"
)

type MonitorMap struct {
	Search     key.Binding
	ExitSearch key.Binding
}

func NewMonitorMap() MonitorMap {
	m := config.Cfg.Keybinds.Monitoring
	return MonitorMap{
		Search: key.NewBinding(
			key.WithKeys(m.Search),
			key.WithHelp(m.Search, "search"),
		),
		ExitSearch: key.NewBinding(
			key.WithKeys(m.ExitSearch),
			key.WithHelp(m.ExitSearch, "exit search"),
		),
	}
}

func (m MonitorMap) Bindings() []key.Binding {
	var bindings []key.Binding

	v := reflect.ValueOf(m)
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if binding, ok := field.Interface().(key.Binding); ok {
			bindings = append(bindings, binding)
		}
	}

	return bindings
}
