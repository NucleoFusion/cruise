package keymap

import (
	"reflect"

	"github.com/charmbracelet/bubbles/key"
)

type MonitorMap struct {
	Search     key.Binding
	ExitSearch key.Binding
}

func NewMonitorMap() MonitorMap {
	return MonitorMap{
		Search: key.NewBinding(
			key.WithKeys("/"),
			key.WithHelp("/", "search"),
		),
		ExitSearch: key.NewBinding(
			key.WithKeys("esc"),
			key.WithHelp("esc", "exit search"),
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
