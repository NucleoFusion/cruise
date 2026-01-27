package keymap

import (
	"reflect"

	"github.com/charmbracelet/bubbles/key"
	"github.com/cruise-org/cruise/pkg/config"
)

type VolMap struct {
	Remove      key.Binding
	Prune       key.Binding
	ShowDetails key.Binding
	ExitDetails key.Binding
}

func NewVolMap() VolMap {
	m := config.Cfg.Keybinds.Volumes
	return VolMap{
		Remove: key.NewBinding(
			key.WithKeys(m.Remove),
			key.WithHelp(m.Remove, "remove"),
		),
		Prune: key.NewBinding(
			key.WithKeys(m.Prune),
			key.WithHelp(m.Prune, "prune"),
		),
		ShowDetails: key.NewBinding(
			key.WithKeys(m.ShowDetails),
			key.WithHelp(m.ShowDetails, "show details"),
		),
		ExitDetails: key.NewBinding(
			key.WithKeys(m.ExitDetails),
			key.WithHelp(m.ExitDetails, "exit details"),
		),
	}
}

func (m VolMap) Bindings() []key.Binding {
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
