package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "learn-to-boil",
	Short: "...",
	Long:  `...`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}
