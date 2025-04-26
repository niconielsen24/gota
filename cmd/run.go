package cmd

import (
	"gota/internal/task"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var RunCmd = &cobra.Command{
	Use:   "run [TASK]",
	Short: "Run task/s",
	Long:  "Run task/s from the YAML file in the current working directory.",
	Args:  cobra.MinimumNArgs(1),
	Run:   run_run,
}

func run_run(cmd *cobra.Command, args []string) {
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

	tasks_to_run := getTasks(cmd, args, tasks.Tasks)
	if err := runTasks(tasks_to_run); err != nil {
		cmd.PrintErrf("Failed to run task: %s", err.Error())
	}
}

func getTasks(cmd *cobra.Command, args []string, task_map map[string]task.Task) []string {
	tasks_to_run := make([]string, 0, len(args))
	for _, arg := range args {
		t, ok := task_map[arg]
		if !ok {
			cmd.Printf("Task \"%s\" does not exist", arg)
			continue
		}
		tasks_to_run = append(tasks_to_run, t.Run)
	}
	return tasks_to_run
}

func runTasks(tasks []string) error {
	for _, task_str := range tasks {
		if strings.TrimSpace(task_str) == "" {
			continue
		}
		task_parts := strings.Fields(task_str)
		task := exec.Command(task_parts[0], task_parts[1:]...)
		task.Stderr = os.Stderr
		task.Stdout = os.Stdout
		if err := task.Run(); err != nil {
			return err
		}
	}
	return nil
}
