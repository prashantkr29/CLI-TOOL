package cmd

import (
	"github.com/spf13/cobra"
)

func Execute() {
	var rootCmd = &cobra.Command{
		Use:   "go-ebpf-monitor",
		Short: "Monitor system and container stats using eBPF",
	}
	rootCmd.AddCommand(container.containerCmd, host.hostCmd)
	rootCmd.Execute()
}
