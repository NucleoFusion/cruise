package keymap

import (
	"reflect"

	"github.com/charmbracelet/bubbles/key"
)

type VulnMap struct {
	FocusScanners key.Binding
	FocusList     key.Binding
}

func NewVulnMap() VulnMap {
	return VulnMap{
		FocusScanners: key.NewBinding(
			key.WithKeys("S"),
			key.WithHelp("S", "focus scanners"),
		),
		FocusList: key.NewBinding(
			key.WithKeys("L"),
			key.WithHelp("L", "focus list"),
		),
	}
}

func (m VulnMap) Bindings() []key.Binding {
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
