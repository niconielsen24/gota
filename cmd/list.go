package cmd

import (
	"gota/internal/task"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List available tasks",
	Long:  "Lists all defined tasks from the YAML file in the current working directory.",
	Args:  cobra.ExactArgs(0),
	Run:   run_list,
}

func run_list(cmd *cobra.Command, args []string) {
	file_bytes, err := os.ReadFile(task.TaskFileStr)
	if err != nil {
		cmd.PrintErrf("Error reading task file: %s\n", err)
		return
	}

	tasks := &task.TaskFile{}
	if err := yaml.Unmarshal(file_bytes, tasks); err != nil {
		cmd.PrintErrf("Error parsing YAML: %s\n", err)
		return
	}

	if len(tasks.Tasks) == 0 {
		cmd.Println("No tasks defined.")
		return
	}

	cmd.Println("\n📋 Available Tasks:")

	for key, val := range tasks.Tasks {
		cmd.Printf("  • %-15s →  %s\n", key, val.Desc)
	}
}
