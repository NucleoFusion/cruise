package keymap

import (
	"reflect"

	"github.com/NucleoFusion/cruise/internal/config"
	"github.com/charmbracelet/bubbles/key"
)

type ContainersMap struct {
	Start       key.Binding
	Exec        key.Binding
	Restart     key.Binding
	Stop        key.Binding
	Remove      key.Binding
	Pause       key.Binding
	Unpause     key.Binding
	PortMap     key.Binding
	ShowDetails key.Binding
	ExitDetails key.Binding
}

func NewContainersMap() ContainersMap {
	m := config.Cfg.Keybinds.Container
	return ContainersMap{
		Start: key.NewBinding(
			key.WithKeys(m.Start),
			key.WithHelp(m.Start, "start"),
		),
		Stop: key.NewBinding(
			key.WithKeys(m.Stop),
			key.WithHelp(m.Stop, "stop"),
		),
		Remove: key.NewBinding(
			key.WithKeys(m.Remove),
			key.WithHelp(m.Remove, "remove"),
		),
		Pause: key.NewBinding(
			key.WithKeys(m.Pause),
			key.WithHelp(m.Pause, "pause"),
		),
		Unpause: key.NewBinding(
			key.WithKeys(m.Unpause),
			key.WithHelp(m.Unpause, "unpause"),
		),
		Restart: key.NewBinding(
			key.WithKeys(m.Restart),
			key.WithHelp(m.Restart, "restart"),
		),
		Exec: key.NewBinding(
			key.WithKeys(m.Exec),
			key.WithHelp(m.Exec, "exec -it"),
		),
		PortMap: key.NewBinding(
			key.WithKeys(m.PortMap),
			key.WithHelp(m.PortMap, "port map"),
		),
		ShowDetails: key.NewBinding(
			key.WithKeys(m.ShowDetails),
			key.WithHelp(m.ShowDetails, "show detail"),
		),
		ExitDetails: key.NewBinding(
			key.WithKeys(m.ExitDetails),
			key.WithHelp(m.ExitDetails, "exit detail"),
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
