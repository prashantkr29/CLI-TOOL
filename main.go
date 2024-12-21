package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/cobra"
)

// syntax to build a basic command utility
var rootCmd = &cobra.Command{
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

func main() {
	rootCmd.AddCommand(osCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
