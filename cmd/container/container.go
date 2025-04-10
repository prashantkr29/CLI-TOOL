package container

import (
	"context"
	"fmt"
	"log"
	"os/exec"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/spf13/cobra"
)

var ContainerCmd = &cobra.Command{
	Use:   "container",
	Short: "Monitor container-related stats",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Monitoring container stats...")
	},
}

var ContainerList = &cobra.Command{
	Use:   "container-list",
	Short: "List out the container running on the host",
	Run: func(cmd *cobra.Command, args []string) {
		cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
		if err != nil {
			log.Fatalf("Error creating Docker client: %v", err)
		}

		// Get a list of running containers
		containers, err := cli.ContainerList(context.Background(), container.ListOptions{})
		if err != nil {
			log.Fatalf("Error listing containers: %v", err)
		}

		// Print container IDs
		fmt.Println("Running Containers:")
		for _, container := range containers {
			fmt.Println(container.ID[:12], container.Image, container.State)
		}
	},
}
var Containerstats = &cobra.Command{
	Use:   "container-stats",
	Short: "Monitor Container stats usage",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Memory Usage for Containers")
		result := exec.Command("docker", "stats", "--no-stream")

		output, err := result.CombinedOutput()
		if err != nil {
			fmt.Println("error fetching the stats")
			return
		}
		fmt.Println(string(output))
	},
}

var ContainerProcess = &cobra.Command{
	Use:   "container-process",
	Short: "Fetch Container Processes",
	Run: func(cmd *cobra.Command, args []string) {
		cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
		if err != nil {
			log.Fatalf("Error creating Docker client: %v", err)
		}

		// Get a list of running containers
		containers, err := cli.ContainerList(context.Background(), container.ListOptions{})
		if err != nil {
			log.Fatalf("Error listing containers: %v", err)
		}

		// Print container IDs
		fmt.Println("Running Containers:")
		for _, container := range containers {
			processList, err := cli.ContainerTop(context.Background(), container.ID[:12], []string{"-aux"})
			if err != nil {
				log.Fatalf("Error retrieving processes: %v", err)
			}
			fmt.Println(container.ID[:12], container.Image, container.State)

			for _, process := range processList.Processes {
				fmt.Println(process)
			}
		}
	},
}
