package keymap

import (
	"reflect"

	"github.com/charmbracelet/bubbles/key"
	"github.com/cruise-org/cruise/pkg/config"
)

type VulnMap struct {
	FocusScanners key.Binding
	FocusList     key.Binding
	Export        key.Binding
}

func NewVulnMap() VulnMap {
	m := config.Cfg.Keybinds.Vulnerability
	return VulnMap{
		FocusScanners: key.NewBinding(
			key.WithKeys(m.FocusScanners),
			key.WithHelp(m.FocusScanners, "focus scanners"),
		),
		FocusList: key.NewBinding(
			key.WithKeys(m.FocusList),
			key.WithHelp(m.FocusList, "focus list"),
		),
		Export: key.NewBinding(
			key.WithKeys(m.Export),
			key.WithHelp(m.Export, "export"),
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
