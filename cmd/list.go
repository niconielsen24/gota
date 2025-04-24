package cmd

import "github.com/spf13/cobra"

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List available tasks",
	Long:  "Lists all defined tasks from the YAML file in the current working directory.",
}
