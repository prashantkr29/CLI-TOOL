package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var containerCmd = &cobra.Command{
	Use:   "container",
	Short: "Monitor container-related stats",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Monitoring container stats...")
	},
}
