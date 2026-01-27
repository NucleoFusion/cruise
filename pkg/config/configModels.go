package config

type Config struct {
	Global   Global   `mapstructure:"global" toml:"global"`
	Styles   Styles   `mapstructure:"styles" toml:"styles"`
	Keybinds Keybinds `mapstructure:"keybinds" toml:"keybinds"`
}

type Global struct {
	ExportDir string   `mapstructure:"export_dir" toml:"export_dir"`
	Term      string   `mapstructure:"term" toml:"term"`
	Runtimes  []string `mapstructure:"runtimes" toml:"runtimes"`
}
