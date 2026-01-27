package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/cruise-org/cruise/internal/utils"
	"github.com/spf13/viper"
)

var Cfg Config

func SetCfg() error {
	cfg := utils.GetCfgDir()
	p := filepath.Join(cfg, "config.toml")

	if _, err := os.Stat(p); os.IsNotExist(err) {
		if err := os.MkdirAll(cfg, 0o755); err != nil {
			fmt.Printf("failed to create config dir: %s", err.Error())
			os.Exit(1)
		}

		if _, err := os.Create(p); err != nil {
			fmt.Printf("failed to create config file: %s", err.Error())
			os.Exit(1)
		}
	}

	viper.SetConfigFile(p)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	Cfg = Default()

	if err := viper.Unmarshal(&Cfg); err != nil {
		return err
	}

	return nil
}
