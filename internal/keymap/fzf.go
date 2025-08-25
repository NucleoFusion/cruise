package keymap

import (
	"reflect"

	"github.com/NucleoFusion/cruise/internal/config"
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
	m := config.Cfg.Keybinds.Fzf
	return FuzzyMap{
		Up: key.NewBinding(
			key.WithKeys(m.Up),
			key.WithHelp(m.Up, "up"),
		),
		Down: key.NewBinding(
			key.WithKeys(m.Down),
			key.WithHelp(m.Down, "down"),
		),
		Enter: key.NewBinding(
			key.WithKeys(m.Enter),
			key.WithHelp(m.Enter, "enter"),
		),
		Exit: key.NewBinding(
			key.WithKeys(m.Exit),
			key.WithHelp(m.Exit, "exit"),
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
