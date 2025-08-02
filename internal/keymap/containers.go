package keymap

import (
	"reflect"

	"github.com/charmbracelet/bubbles/key"
)

type ContainersMap struct {
	Start   key.Binding
	Exec    key.Binding
	Restart key.Binding
	Stop    key.Binding
	Remove  key.Binding
	Pause   key.Binding
}

func NewContainersMap() ContainersMap {
	return ContainersMap{
		Start: key.NewBinding(
			key.WithKeys("s"),
			key.WithHelp("s", "start"),
		),
		Stop: key.NewBinding(
			key.WithKeys("t"),
			key.WithHelp("t", "stop"),
		),
		Remove: key.NewBinding(
			key.WithKeys("d"),
			key.WithHelp("d", "remove"),
		),
		Pause: key.NewBinding(
			key.WithKeys("p"),
			key.WithHelp("p", "pause"),
		),
		Restart: key.NewBinding(
			key.WithKeys("r"),
			key.WithHelp("r", "restart"),
		),
		Exec: key.NewBinding(
			key.WithKeys("e"),
			key.WithHelp("e", "exec -it"),
		),
	}
}

func (m ContainersMap) Bindings() []key.Binding {
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
