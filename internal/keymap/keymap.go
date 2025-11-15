package keymap

import (
	"fmt"

	"github.com/NucleoFusion/cruise/internal/config"
	"github.com/charmbracelet/bubbles/key"
)

type DynamicMap struct {
	keys []key.Binding
}

func QuickQuitKey() key.Binding {
	q := config.Cfg.Keybinds.Global.QuickQuit
	return key.NewBinding(key.WithKeys(q),
		key.WithHelp(fmt.Sprintf(" %s ", q), "quit"))
}

func NewDynamic(keys []key.Binding) *DynamicMap {
	return &DynamicMap{
		keys: keys,
	}
}

func (d DynamicMap) ShortHelp() []key.Binding {
	return d.keys
}

func (d DynamicMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{d.keys}
}
