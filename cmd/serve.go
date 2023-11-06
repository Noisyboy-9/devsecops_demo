/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/noisyboy-9/golang_api_template/internal/app"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the serving the HTTP server",
	Run:   serveCmdRunner,
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func serveCmdRunner(cmd *cobra.Command, args []string) {
	app.InitApp()
	app.SetupGracefulShutdown()
}
