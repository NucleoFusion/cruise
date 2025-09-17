package cmd

import (
	"fmt"
	"log"
	"os"
	"runtime/debug"

	"github.com/NucleoFusion/cruise/internal/config"
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

	defer func() {
		if r := recover(); r != nil {
			log.Printf("PANIC: %v", r)
			log.Printf("STACK:\n%s", debug.Stack())
		}
	}()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
