package config

type Keybinds struct {
	Global        GlobalKeybinds        `mapstructure:"global"`
	Container     ContainersKeybinds    `mapstructure:"container"`
	Images        ImagesKeybinds        `mapstructure:"images"`
	Fzf           FzfKeybinds           `mapstructure:"fzf"`
	Monitoring    MonitorKeybinds       `mapstructure:"monitoring"`
	Network       NetworkKeybinds       `mapstructure:"network"`
	Volumes       VolumeKeybinds        `mapstructure:"volume"`
	Vulnerability VulnerabilityKeybinds `mapstructure:"vulnerability"`
}

type GlobalKeybinds struct {
	PageFinder    string `mapstructure:"page_finder"`
	ListUp        string `mapstructure:"list_up"`
	ListDown      string `mapstructure:"list_down"`
	FocusSearch   string `mapstructure:"focus_search"`
	UnfocusSearch string `mapstructure:"unfocus_search"`
}

type ContainersKeybinds struct {
	Start       string `mapstructure:"start"`
	Exec        string `mapstructure:"exec"`
	Restart     string `mapstructure:"restart"`
	Stop        string `mapstructure:"stop"`
	Remove      string `mapstructure:"remove"`
	Pause       string `mapstructure:"pause"`
	Unpause     string `mapstructure:"unpause"`
	PortMap     string `mapstructure:"port_map"`
	ShowDetails string `mapstructure:"show_details"`
	ExitDetails string `mapstructure:"exit_details"`
}

type FzfKeybinds struct {
	Up    string `mapstructure:"up"`
	Down  string `mapstructure:"down"`
	Enter string `mapstructure:"enter"`
	Exit  string `mapstructure:"exit"`
}

type ImagesKeybinds struct {
	Remove string `mapstructure:"remove"`
	Prune  string `mapstructure:"prune"`
	Push   string `mapstructure:"push"`
	Pull   string `mapstructure:"pull"`
	Build  string `mapstructure:"build"`
	Layers string `mapstructure:"layers"`
	Exit   string `mapstructure:"exit"`
}

type MonitorKeybinds struct {
	Search     string `mapstructure:"search"`
	ExitSearch string `mapstructure:"exit_search"`
}

type NetworkKeybinds struct {
	Remove      string `mapstructure:"remove"`
	Prune       string `mapstructure:"prune"`
	ShowDetails string `mapstructure:"show_details"`
	ExitDetails string `mapstructure:"exit_details"`
}

type VolumeKeybinds struct {
	Remove      string `mapstructure:"remove"`
	Prune       string `mapstructure:"prune"`
	ShowDetails string `mapstructure:"show_details"`
	ExitDetails string `mapstructure:"exit_details"`
}

type VulnerabilityKeybinds struct {
	FocusScanners string `mapstructure:"focus_scanners"`
	FocusList     string `mapstructure:"focus_list"`
}
