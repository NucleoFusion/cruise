package config

type Config struct {
	Global   Global   `mapstructure:"global"`
	Styles   Styles   `mapstructure:"styles"`
	Keybinds Keybinds `mapstructure:"keybinds"`
}

type Global struct {
	ExportDir string `mapstructure:"export_dir"`
}
