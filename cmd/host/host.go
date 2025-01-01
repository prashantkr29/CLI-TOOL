package host

import (
	"fmt"
	"time"

	"sort"

	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/process"
	"github.com/spf13/cobra"
)

func monitorDiskUsage() {
	fmt.Printf("%-20s %-10s %-10s %-10s %-10s\n", "Filesystem", "Total", "Used", "Free", "Usage%")
	fmt.Println("--------------------------------------------------------------")

	partitions, err := disk.Partitions(false)
	if err != nil {
		fmt.Printf("Error fetching disk partitions: %v\n", err)
		return
	}

	for _, partition := range partitions {
		usage, err := disk.Usage(partition.Mountpoint)
		if err != nil {
			fmt.Printf("Error fetching disk usage for %s: %v\n", partition.Mountpoint, err)
			continue
		}

		fmt.Printf("%-20s %-10.2f %-10.2f %-10.2f %-10.2f%%\n",
			partition.Device,
			float64(usage.Total)/(1024*1024*1024),
			float64(usage.Used)/(1024*1024*1024),
			float64(usage.Free)/(1024*1024*1024),
			usage.UsedPercent)
	}
}

func monitorDiskIO() {
	fmt.Printf("%-10s %-10s %-10s %-10s %-10s\n", "Device", "Read/s", "Write/s", "ReadBytes/s", "WriteBytes/s")
	fmt.Println("--------------------------------------------------------")

	ioStats, err := disk.IOCounters()
	if err != nil {
		fmt.Printf("Error fetching disk IO stats: %v\n", err)
		return
	}

	for device, stats := range ioStats {
		fmt.Printf("%-10s %-10d %-10d %-10d %-10d\n",
			device,
			stats.ReadCount,
			stats.WriteCount,
			stats.ReadBytes,
			stats.WriteBytes)
	}
}

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
			times, err := cpu.Times(false)
			if err != nil {
				fmt.Printf("Error fetching CPU times: %v\n", err)
				return
			}

			// Get CPU usage percentage
			percentages, err := cpu.Percent(time.Second, true)
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
var HostProcess = &cobra.Command{
	Use:   "Host-process",
	Short: "Tracks the process and its stats",
	Run: func(cmd *cobra.Command, args []string) {
		for {
			// Get all processes on the system
			procs, err := process.Processes()
			if err != nil {
				fmt.Printf("Error fetching processes: %v\n", err)
				return
			}

			// Clear the screen

			fmt.Print("\033[H\033[J")

			fmt.Printf("PROCESS MONITORING \n")
			fmt.Printf("-------------------------------------------------------------\n")
			fmt.Printf("%-10s %-25s %-10s %-10s %-10s\n", "PID", "Name", "CPU(%)", "Mem(%)", "Status")
			fmt.Printf("-------------------------------------------------------------\n")

			var processStats []struct {
				pid    int32
				name   string
				cpu    float64
				mem    float64
				status string
			}

			// Collect stats for each process
			for _, p := range procs {

				// Get process individual details
				cpuPercent, _ := p.CPUPercent()
				memPercent, _ := p.MemoryPercent()
				name, _ := p.Name()
				status, _ := p.Status()

				processStats = append(processStats, struct {
					pid    int32
					name   string
					cpu    float64
					mem    float64
					status string
				}{
					pid:    p.Pid,
					name:   name,
					cpu:    cpuPercent,
					mem:    float64(memPercent),
					status: status[0],
				})
			}

			// Sort the processes (currently by CPU usage)
			sort.SliceStable(processStats, func(i, j int) bool {
				return processStats[i].pid < processStats[j].pid
			})

			// Print process stats
			for _, stat := range processStats {
				fmt.Printf("%-10d %-25s %-10.2f %-10.2f %-10s\n", stat.pid, stat.name, stat.cpu, stat.mem, stat.status)
			}

			fmt.Printf("-------------------------------------------------------------\n")
			time.Sleep(2 * time.Second)
		}
	},
}

var HostDisk = &cobra.Command{
	Use:   "Host-disk",
	Short: "Monitor disk stats",
	Run: func(cmd *cobra.Command, args []string) {
		for {
			fmt.Print("\033[H\033[2J")
			fmt.Println("Disk Usage:")
			monitorDiskUsage()

			fmt.Println("\nDisk I/O:")
			monitorDiskIO()

			// Refresh every 2 seconds
			time.Sleep(2 * time.Second)
		}
	},
}
