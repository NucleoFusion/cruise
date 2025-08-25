package config

func Default() Config {
	return Config{
		Global: Global{
			ExportDir: "/home/nucleofusion/.cruise/export", // Absolute Path
		},
		Keybinds: Keybinds{ // TODO: Refactor keymap to use this
			Global: GlobalKeybinds{
				PageFinder:    "tab",
				ListUp:        "up",
				ListDown:      "down",
				FocusSearch:   "/",
				UnfocusSearch: "esc",
			},
			Fzf: FzfKeybinds{
				Up:    "up",
				Down:  "down",
				Enter: "enter",
				Exit:  "exit",
			},
			Container: ContainersKeybinds{
				Start:       "s",
				Stop:        "t",
				Remove:      "d",
				Restart:     "r",
				Pause:       "p",
				Unpause:     "u",
				Exec:        "e",
				ShowDetails: "enter",
				ExitDetails: "esc",
				PortMap:     "m",
			},
			Images: ImagesKeybinds{
				Remove: "r",
				Prune:  "d",
				Push:   "p",
				Pull:   "u",
				Build:  "b",
				Layers: "l",
			},
			Network: NetworkKeybinds{
				ExitDetails: "esc",
				ShowDetails: "enter",
				Prune:       "p",
				Remove:      "r",
			},
			Volumes: VolumeKeybinds{
				ExitDetails: "esc",
				ShowDetails: "enter",
				Prune:       "p",
				Remove:      "r",
			},
			Vulnerability: VulnerabilityKeybinds{
				FocusScanners: "S",
				FocusList:     "L",
			},
			Monitoring: MonitorKeybinds{
				Search:     "/",
				ExitSearch: "esc",
			},
		},
		Styles: Styles{
			Text:             "#cdd6f4",
			SubtitleText:     "#74c7ec",
			SubtitleBg:       "#313244",
			MenuSelectedBg:   "#b4befe",
			MenuSelectedText: "#1e1e2e",
			FocusedBorder:    "#b4befe",
			UnfocusedBorder:  "#45475a",
			HelpKeyBg:        "#313244",
			HelpKeyText:      "#cdd6f4",
			HelpDescText:     "#6c7086",
			ErrorText:        "#f38ba8",
			ErrorBg:          "#11111b",
			PopupBorder:      "#74c7ec",
			PlaceholderText:  "#585b70",
			MsgText:          "#74c7ec",
		}, // TODO: Fill styles and refactor colors package and its uses
	}
}
