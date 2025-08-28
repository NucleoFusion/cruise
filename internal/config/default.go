package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func Default() Config {
	expDir := GetDefExportDir()

	if _, err := os.Stat(expDir); os.IsNotExist(err) {
		if err := os.MkdirAll(expDir, 0o755); err != nil {
			fmt.Printf("failed to create config dir: %s", err.Error())
			os.Exit(1)
		}
	}

	return Config{
		Global: Global{
			ExportDir: expDir,
			Shell:     "bash", // TODO: Refactor for shell
		},
		Keybinds: Keybinds{
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
				Exit:  "esc",
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
			HelpKeyText:      "#1e1e2e", // TODO: Docs change
			HelpDescText:     "#6c7086",
			ErrorText:        "#f38ba8",
			ErrorBg:          "#11111b",
			PopupBorder:      "#74c7ec",
			PlaceholderText:  "#585b70",
			MsgText:          "#74c7ec",
		},
	}
}

func GetDefExportDir() string {
	switch runtime.GOOS {
	case "linux", "darwin":
		home, _ := os.UserHomeDir()
		return filepath.Join(home, ".cruise")
	case "windows":
		home, _ := os.UserHomeDir()
		return filepath.Join(home, ".cruise", "export")
	default:
		cfg, _ := os.UserConfigDir()
		return cfg
	}
}
