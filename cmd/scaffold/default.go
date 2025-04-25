package scaffold

import (
	"gota/cmd/templates"
	"os"
	"text/template"
)

func CreateDefaultBase(appName string, main_file *os.File) error {
	if err := createDefaultFiles(); err != nil {
		return nil
	}

	return populateDefaultFiles(appName, main_file)
}

func createDefaultFiles() error {
	return os.Mkdir(default_src_dir, dir_permissions)
}

func populateDefaultFiles(appName string, main_file *os.File) error {
	default_main_tpl, _ := templates.FS.ReadFile(template_files["default_main"])
	default_main_content := template.Must(template.New("default_main").Parse(string(default_main_tpl)))
	err := default_main_content.Execute(main_file, map[string]any{
		"Appname": appName,
	})
	return err
}
