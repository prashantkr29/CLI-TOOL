package host

import (
	"fmt"

	"github.com/spf13/cobra"
)

var HostCmd = &cobra.Command{
	Use:   "host",
	Short: "Monitor host-related stats",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Monitoring host stats...")

	},
}