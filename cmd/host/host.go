package host

import (
	"fmt"

	"github.com/spf13/cobra"
)

var HostCmd = &cobra.Command{
	Use:   "Host",
	Short: "Monitor host-related stats",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Monitoring host stats...")

	},
}
var HostCpu = &cobra.Command{
	Use:   "Host-cpu",
	Short: "Monitor Host-cpu usage",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Monitor cpu stats for Host")
	},
}
var HostMemory = &cobra.Command{
	Use:   "Host-memory",
	Short: "Monitor Host-memory usage",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Monitor memory stats for Host")
	},
}
