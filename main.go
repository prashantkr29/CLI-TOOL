package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/prashantkr29/CLI-TOOL/cmd/container"
	"github.com/prashantkr29/CLI-TOOL/cmd/host"
	"github.com/spf13/cobra"
)

// syntax to build a basic command utility
var Cmd = &cobra.Command{
	Use:   "mycli",
	Short: "root command for terminal",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("welcome to the Terminal")
	},
}
var osCmd = &cobra.Command{
	Use:   "os-check",
	Short: "list the operating system details",
	Run: func(cmd *cobra.Command, args []string) {
		osName := runtime.GOOS
		arch := runtime.GOARCH
		hostname, err := os.Hostname()
		if err == nil {
			hostname = "unknown"
		}
		fmt.Println(osName)
		fmt.Println(arch)
		fmt.Println(hostname)
	},
}

var RootCmd = &cobra.Command{
	Use:   "monitor",
	Short: "Monitor system and container stats using eBPF",
	Run: func(RootCmd *cobra.Command, args []string) {
		fmt.Print("run commands for host and container")
	},
}

func main() {
	RootCmd.AddCommand(container.ContainerCmd, host.HostCmd)
	RootCmd.AddCommand(container.ContainerCpu, host.HostCpu)
	RootCmd.AddCommand(container.ContainerMemory, host.HostMemory, host.HostProcess, host.HostDisk)
	Cmd.AddCommand(osCmd, RootCmd)
	if err := Cmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
