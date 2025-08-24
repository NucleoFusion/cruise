package config

import (
	"path/filepath"

	"github.com/NucleoFusion/cruise/internal/utils"
	"github.com/spf13/viper"
)

var Cfg Config

func SetCfg() error {
	cfg := utils.GetCfgDir()
	viper.SetConfigFile(filepath.Join(cfg, "config.toml"))

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	Cfg = Default()

	if err := viper.Unmarshal(&Cfg); err != nil {
		return err
	}

	return nil
}
