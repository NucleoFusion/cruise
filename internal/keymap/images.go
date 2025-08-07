package keymap

import (
	"reflect"

	"github.com/charmbracelet/bubbles/key"
)

type ImagesMap struct {
	Remove key.Binding
	Prune  key.Binding
	Layers key.Binding
}

func NewImagesMap() ImagesMap {
	return ImagesMap{
		Remove: key.NewBinding(
			key.WithKeys("r"),
			key.WithHelp("r", "remove"),
		),
		Prune: key.NewBinding(
			key.WithKeys("p"),
			key.WithHelp("p", "prune"),
		),
		Layers: key.NewBinding(
			key.WithKeys("l"),
			key.WithHelp("l", "layers"),
		),
	}
}

func (m ImagesMap) Bindings() []key.Binding {
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
