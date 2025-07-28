package keymap

import (
	"reflect"

	"github.com/charmbracelet/bubbles/key"
)

type FuzzyMap struct {
	StartWriting key.Binding
	Up           key.Binding
	Down         key.Binding
	Enter        key.Binding
	Exit         key.Binding
}

func NewFuzzyMap() FuzzyMap {
	return FuzzyMap{
		Up: key.NewBinding(
			key.WithKeys("up"),
			key.WithHelp("up", "up"),
		),
		Down: key.NewBinding(
			key.WithKeys("down"),
			key.WithHelp("down", "down"),
		),
		Enter: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "enter"),
		),
		Exit: key.NewBinding(
			key.WithKeys("esc"),
			key.WithHelp("esc", "exit"),
		),
	}
}

func (m FuzzyMap) Bindings() []key.Binding {
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
