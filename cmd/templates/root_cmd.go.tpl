package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
  // add subcommands to root command ...
}

var rootCmd = &cobra.Command{
	Use:   "{{.Appname}}",
  {{if .ShortDesc}} Short: "{{.ShortDesc}}",{{end}}
  {{if .LongDesc}} Long: "{{.LongDesc}}",{{end}}
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to {{.Appname}}!")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
