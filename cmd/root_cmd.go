package cmd

import (
	"github.com/spf13/cobra"

	"github.com/emadghaffari/go-chat/app"
)

var (
	Runner     CommandLine = &command{}
	configFile             = ""
	debug      bool
	cPath      = "config.yaml"
)

type CommandLine interface {
	RootCmd() *cobra.Command
}

type command struct{}

// rootCmd will run the log streamer
var rootCmd = cobra.Command{
	Use:  "micro",
	Long: "A service that will validate restful transactions and send them to stripe.",
	Run: func(cmd *cobra.Command, args []string) {
		app.Base.StartApplication()
	},
}

// RootCmd will add flags and subcommands to the different commands
func (c *command) RootCmd() *cobra.Command {
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "", "The configuration file")
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "The service debug(true is production - false is dev)")

	// add more commands
	return &rootCmd
}
