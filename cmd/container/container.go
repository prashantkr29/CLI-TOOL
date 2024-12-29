package container

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ContainerCmd = &cobra.Command{
	Use:   "container",
	Short: "Monitor container-related stats",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Monitoring container stats...")
	},
}

var ContainerCpu = &cobra.Command{
	Use:   "container-cpu",
	Short: "Monitor Container-cpu usage",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Monitor cpu stats for container")
	},
}
var ContainerMemory = &cobra.Command{
	Use:   "container-memory",
	Short: "Monitor Container-memory usage",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Monitor memory stats for container")
	},
}
