package cmd

import (
	"gota/cmd/scaffold"
	"gota/internal/task"

	"github.com/spf13/cobra"
)

var CreateCmd = &cobra.Command{
	Use:   "create TYPE APPNAME",
	Short: "Create a new project structure",
	Long: `Scaffold a new project of the given TYPE with the provided APPNAME.

This sets up the basic folder layout, a main file, and other starter files depending on the type.
Helpful if you want to go from zero to Go in a hurry.`,
	Args: cobra.MinimumNArgs(2),
	Run:  run_create,
}

func run_create(cmd *cobra.Command, args []string) {
	typ, appName := args[0], args[1]

	cmd.Printf("Creating new %v project, name: %v\n", typ, appName)

	output, err := scaffold.Dispatch(typ, appName)
	cmd.Println(output)
	if err != nil {
		cmd.PrintErrf("Error: %s\n", err)
	}
	if err := task.BuildTaskFile(task.DefaultTasks); err != nil {
		cmd.PrintErrf("Error: %s\n", err)
	}
}
