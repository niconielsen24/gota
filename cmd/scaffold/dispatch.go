package scaffold

import (
	"fmt"
	"os"
)

type Creator func(appName string, main_file *os.File) error

var dispatch = map[string]Creator{
	"cli":     CreateCliBase,
	"default": CreateDefaultBase,
}

func Dispatch(typ, appName string) (string, error) {
	project_f, ok := dispatch[typ]
	if !ok {
		return "", fmt.Errorf("project type %q does not exist", typ)
	}

	output, mainFile, err := CreateProjectBase(appName)
	if err != nil {
		return output, err
	}
	defer mainFile.Close()

	if err := project_f(appName, mainFile); err != nil {
		return output, err
	}

	return output, nil
}
