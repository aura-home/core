package cmd

import (
	aura "github.com/aura-home/core/internal/app"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starting Aura Home",
	Run: func(cmd *cobra.Command, args []string) {
		configPath := cmd.Flag("config").Value.String()
		app := aura.NewApp(configPath)
		app.MostStart()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().StringP("config", "c", "./config/debug.yaml", "config file")
}
