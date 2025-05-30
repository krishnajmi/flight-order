package cmd

import (
	"log/slog"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "flight-cli",
	Short: "all application for flight",
	Long:  `flight-order collection command`,
}

func init() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))
	cobra.OnInitialize(initServer)
}

func Run() {
	rootCmd.Execute()
}
