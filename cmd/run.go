package cmd

import (
	"fmt"

	"github.com/NucleoFusion/cruise/internal/models/root"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the Toney TUI",
	Run: func(cmd *cobra.Command, args []string) {
		p := tea.NewProgram(root.NewRoot(), tea.WithAltScreen())
		if _, err := p.Run(); err != nil {
			fmt.Println("Alas, error")
			fmt.Println(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
