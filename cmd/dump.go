package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/NucleoFusion/cruise/internal/config"
	"github.com/NucleoFusion/cruise/internal/utils"
	"github.com/pelletier/go-toml/v2"
	"github.com/spf13/cobra"
)

var dumpPath string

var dumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "Dump the config file at the given filepath.",
	Run: func(cmd *cobra.Command, args []string) {
		path, _ := cmd.Flags().GetString("path")
		fmt.Println("Dumping config to:", path)

		if filepath.Ext(path) == "" {
			path = filepath.Join(path, "dump.toml")
		}

		f, err := os.Create(path)
		if err != nil {
			fmt.Println("Error Creating File: ", err.Error())
			return
		}
		defer f.Close()

		encoder := toml.NewEncoder(f)
		err = encoder.Encode(config.Default())
		if err != nil {
			fmt.Println("Failed to write to file: ", err.Error())
			return
		}

		fmt.Println("Successfully Dumped file to " + path)
	},
}

func init() {
	dumpCmd.Flags().StringVarP(&dumpPath, "path", "p", filepath.Join(utils.GetCfgDir(), "dump.toml"), "filepath to dump the config")
	rootCmd.AddCommand(dumpCmd)
}
