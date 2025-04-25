package scaffold

import (
	"gota/cmd/templates"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"
)

func CreateCliBase(appName string, main_file *os.File) error {
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
	_ = os.Mkdir(cli_cmd_dir, dir_permissions)
	cwd, _ := os.Getwd()
	root_cmd_file, _ := os.Create(filepath.Join(cwd, cli_cmd_dir, "root.go"))

	return root_cmd_file, nil
}

func populateCliFiles(appName string, main_file, root_cmd_file *os.File) error {
	main_tpl, _ := templates.FS.ReadFile(template_files["cli_main"])
	main_content := template.Must(template.New("cli_main").Parse(string(main_tpl)))
	err := main_content.Execute(main_file, map[string]any{
		"Appname": appName,
	})
	if err != nil {
		return err
	}

	root_cmd_tpl, _ := templates.FS.ReadFile(template_files["cli_root_cmd"])
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
