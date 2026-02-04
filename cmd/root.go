package cmd

import (
	"fmt"
	"os"

	"github.com/cruise-org/cruise/pkg/config"
	"github.com/cruise-org/cruise/pkg/runtimes"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cruise",
	Short: "Cruise is a TUI for Docker",
	Long:  `Cruise is a powerful terminal UI for managing your docker containers, composes and much more.`,
	Run: func(cmd *cobra.Command, args []string) {
		runCmd.Run(cmd, args)
	},
}

func Execute() {
	if err := config.SetCfg(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := runtimes.InitializeService(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
