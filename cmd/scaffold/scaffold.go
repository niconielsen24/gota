package scaffold

import (
	"os"
	"os/exec"
	"path/filepath"
)

// Directories
const cli_cmd_dir = "cmd"
const default_src_dir = "src"

// Files tpl
var template_files = map[string]string{
	"cli_main":     "cli_main.go.tpl",
	"default_main": "default_main.go.tpl",
	"cli_root_cmd": "root_cmd.go.tpl",
}

// Owner has read write and exec permisions
// other users only have read and exec
const dir_permissions = 0755

// Creates the base common to every project, i.e. root dir, main file,
// and initializes the go mod file.
func CreateProjectBase(appName string) (string, *os.File, error) {
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
