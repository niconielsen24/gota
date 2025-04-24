package cmd

import (
	"gota/cmd/templates"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"

	"github.com/spf13/cobra"
)

// Owner has read write and exec permisions
// other users only have read and exec
const dir_permissions = 0755

var dispatch = map[string]func(appName string, main_file *os.File) error{
	"cli":     createCliBase,
	"web-api": func(appName string, main_file *os.File) error { return nil },
}

var CreateCmd = &cobra.Command{
	Use:   "create TYPE APPNAME",
	Short: "Create a new project structure",
	Long: `Scaffold a new project of the given TYPE with the provided APPNAME.

This sets up the basic folder layout, a main file, and other starter files depending on the type.
Helpful if you want to go from zero to Go in a hurry.`,
	Args: cobra.MinimumNArgs(2),
	Run:  run,
}

func run(cmd *cobra.Command, args []string) {
	typ := args[0]
	appName := args[1]

	project_f, exists := dispatch[typ]
	if !exists {
		cmd.PrintErrf("Project type: %v, does not exist\n", typ)
		return
	}

	cmd.Printf("Creating new %v project, name: %v\n", typ, appName)

	output, main_file, err := createProjectBase(appName)
	cmd.Println(output)
	defer main_file.Close()

	if err != nil {
		cmd.PrintErrf("Error running command: %s\n", err.Error())
		return
	}

	if err := project_f(appName, main_file); err != nil {
		cmd.PrintErrf("Error running command: %s\n", err.Error())
		return
	}
}

// Creates the base common to every project, i.e. root dir, main file,
// and initializes the go mod file.
func createProjectBase(appName string) (string, *os.File, error) {
	os.Mkdir(appName, dir_permissions)
	cwd, _ := os.Getwd()
	_ = os.Chdir(filepath.Join(cwd, appName))

	output, err := exec.Command("go", "mod", "init", appName).CombinedOutput()
	if err != nil {
		return string(output), nil, err
	}

	main_file, err := os.Create(appName + ".go")
	if err != nil {
		return err.Error(), nil, err
	}

	return string(output), main_file, nil
}

func createCliBase(appName string, main_file *os.File) error {
	root_cmd_file, err := createCliFiles()
	if err != nil {
		return err
	}
	if err := populateCliFiles(appName, main_file, root_cmd_file); err != nil {
		return err
	}
	if err := getCliPackages(); err != nil {
		return err
	}

	return nil
}

func createCliFiles() (*os.File, error) {
	_ = os.Mkdir("cmd", dir_permissions)
	os.Mkdir("cmd", dir_permissions)
	cwd, _ := os.Getwd()
	root_cmd_file, _ := os.Create(filepath.Join(cwd, "cmd", "root.go"))

	return root_cmd_file, nil
}

func populateCliFiles(appName string, main_file, root_cmd_file *os.File) error {
	main_tpl, _ := templates.FS.ReadFile("main.go.tpl")
	main_content := template.Must(template.New("main").Parse(string(main_tpl)))
	err := main_content.Execute(main_file, map[string]any{
		"Appname": appName,
	})
	if err != nil {
		return err
	}

	root_cmd_tpl, _ := templates.FS.ReadFile("root_cmd.go.tpl")
	root_cmd_content := template.Must(template.New("root_cmd").Parse(string(root_cmd_tpl)))
	err = root_cmd_content.Execute(root_cmd_file, map[string]any{
		"Appname": appName,
	})
	return err

}

func getCliPackages() error {
	cmd := exec.Command("go", "get", "github.com/spf13/cobra")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
