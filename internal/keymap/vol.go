package keymap

import (
	"reflect"

	"github.com/charmbracelet/bubbles/key"
)

type VolMap struct {
	Remove      key.Binding
	Prune       key.Binding
	ShowDetails key.Binding
	ExitDetails key.Binding
}

func NewVolMap() VolMap {
	return VolMap{
		Remove: key.NewBinding(
			key.WithKeys("r"),
			key.WithHelp("r", "remove"),
		),
		Prune: key.NewBinding(
			key.WithKeys("p"),
			key.WithHelp("p", "prune"),
		),
		ShowDetails: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "show details"),
		),
		ExitDetails: key.NewBinding(
			key.WithKeys("esc"),
			key.WithHelp("esc", "exit details"),
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
