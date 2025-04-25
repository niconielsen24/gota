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
		cmd.PrintErrf("Error: %s\n", err)
	}

	tasks := &task.TaskFile{}
	if err := yaml.Unmarshal(file_bytes, tasks); err != nil {
		cmd.PrintErrf("Error: %s\n", err)
	}

	cmd.Println("Tasks :")
	for key, val := range tasks.Tasks {
		cmd.Printf(" |> %s:\n\t%s\n", key, val.Desc)
	}
}
