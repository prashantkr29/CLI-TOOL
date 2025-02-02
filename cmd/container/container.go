package container

import (
	"fmt"
	"io"
	"os/exec"
	"strings"

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
		command := exec.Command("curl", "--unix-socket", "/var/run/docker.sock", "http://localhost/containers/json")
		output, err := command.Output()
		if err != nil {
			fmt.Println("Error executing curl command:", err)
			return
		}

		// Use jq to filter and extract the container IDs
		cmdJq := exec.Command("jq", "-r", ".[].Id")
		cmdJq.Stdin = strings.NewReader(string(output))
		containerIds, err := cmdJq.Output()
		if err != nil {
			fmt.Println("Error executing jq command:", err)
			return
		}

		// Print the container IDs
		fmt.Println(string(containerIds))
		ids := strings.Split(string(containerIds), "\n")

		// Iterate over the container IDs
		for _, containerID := range ids {
			if containerID == "" {
				continue // Skip empty container IDs
			}
			fmt.Printf("Fetching CPU stats for container: %s\n", containerID)

			// io.Pipe() to connect the output of curl to jq
			pr, pw := io.Pipe()

			cmdStats := exec.Command("curl", "--unix-socket", "/var/run/docker.sock", fmt.Sprintf("http://localhost/containers/%s/stats?stream=false", containerID))
			cmdStats.Stdout = pw // Set the pipe's writer to the curl command's output

			//Execute jq to process the output of curl
			cmdJqStats := exec.Command("jq", ".cpu_stats.cpu_usage.total_usage, .cpu_stats.system_cpu_usage")
			cmdJqStats.Stdin = pr // Set the pipe's reader to jq's input

			//Start the jqStats and stats commands
			go func() {
				if err := cmdStats.Run(); err != nil {
					fmt.Println("Error running curl:", err)
				}
				pw.Close() // Close the pipe once curl finishes
			}()

			statsOutput, err := cmdJqStats.Output()
			if err != nil {
				fmt.Println("Error fetching CPU stats:", err)
				continue
			}

			// Print CPU stats (total usage and system CPU usage)
			fmt.Printf("CPU Stats for container cpu usage and system cpu usage %s: %s\n", containerID, "\n", string(statsOutput))
		}
	},
}
var ContainerMemory = &cobra.Command{
	Use:   "container-memory",
	Short: "Monitor Container-memory usage",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Monitor memory stats for container")
	},
}
