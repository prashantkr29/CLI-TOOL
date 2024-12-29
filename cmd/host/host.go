package host

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
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
		for {
			times, err := cpu.Times(false) // false = aggregated CPU stats
			if err != nil {
				fmt.Printf("Error fetching CPU times: %v\n", err)
				return
			}

			// Get CPU usage percentage
			percentages, err := cpu.Percent(time.Second, true) // true = per-core stats
			if err != nil {
				fmt.Printf("Error fetching CPU usage: %v\n", err)
				return
			}

			// Clear the screen (for a dynamic top-like effect)
			fmt.Print("\033[H\033[2J")

			// Display CPU details
			fmt.Printf("CPU USAGE (similar to 'top' command)\n")
			fmt.Printf("-------------------------------------------------------------\n")

			// Aggregate CPU Times (First CPU times struct is aggregated for all CPUs)
			if len(times) > 0 {
				fmt.Printf("User Time:      %.2f%%\n", times[0].User)
				fmt.Printf("System Time:    %.2f%%\n", times[0].System)
				fmt.Printf("Idle Time:      %.2f%%\n", times[0].Idle)
				fmt.Printf("I/O Wait Time:  %.2f%%\n", times[0].Iowait)
			}

			fmt.Printf("-------------------------------------------------------------\n")

			// Display Per-CPU Utilization
			for i, usage := range percentages {
				fmt.Printf("CPU %d Usage:   %.2f%%\n", i, usage)
			}

			fmt.Printf("-------------------------------------------------------------\n")

			// Sleep before the next update
			time.Sleep(2 * time.Second)
		}
	},
}
var HostMemory = &cobra.Command{
	Use:   "Host-memory",
	Short: "Monitor Host-memory usage",
	Run: func(cmd *cobra.Command, args []string) {
		for {
			memory, err := mem.VirtualMemory()
			if err != nil {
				fmt.Printf("Error fetching memory stats: %v\n", err)
				return
			}

			// Display memory usage stats
			// Clear the screen (for a dynamic top-like effect)
			fmt.Print("\033[H\033[2J")

			// Display formatted memory stats
			fmt.Printf("MEMORY USAGE \n")
			fmt.Printf("-------------------------------------------------------------\n")
			fmt.Printf("Total:     %10.2f GB\n", float64(memory.Total)/(1024*1024*1024))
			fmt.Printf("Used:      %10.2f GB   (%6.2f%%)\n", float64(memory.Used)/(1024*1024*1024), memory.UsedPercent)
			fmt.Printf("Free:      %10.2f GB\n", float64(memory.Free)/(1024*1024*1024))
			fmt.Printf("Available: %10.2f GB\n", float64(memory.Available)/(1024*1024*1024))
			fmt.Printf("-------------------------------------------------------------\n")

			// Sleep for a while before updating
			time.Sleep(2 * time.Second)
		}
	},
}
