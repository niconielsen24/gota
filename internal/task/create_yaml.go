package task

import (
	"os"

	"gopkg.in/yaml.v3"
)

const TaskFileStr = ".task.yaml"

type TaskFile struct {
	Tasks map[string]Task `yaml:"tasks"`
}

type Task struct {
	Desc string `yaml:"desc"`
	Run  string `yaml:"run"`
}

// Default project tasks
var DefaultTasks = []struct {
	Alias string
	Task  Task
}{
	{"build", Task{Desc: "Compile the Go project", Run: "go build -v ./..."}},
	{"run", Task{Desc: "Run the main Go application", Run: "go run ."}},
	{"test", Task{Desc: "Run all Go tests", Run: "go test -v ./..."}},
	{"tidy", Task{Desc: "Clean up go.mod and go.sum", Run: "go mod tidy"}},
	{"fmt", Task{Desc: "Format all Go files", Run: "go fmt ./..."}},
	{"lint", Task{Desc: "Lint Go code with go vet", Run: "go vet ./..."}},
	{"clean", Task{Desc: "Remove build artifacts", Run: "go clean"}},
	{"deps", Task{Desc: "Download project dependencies", Run: "go mod download"}},
	{"coverage", Task{Desc: "Run tests with coverage", Run: "go test -cover ./..."}},
	{"install", Task{Desc: "Build and install to GOPATH/bin", Run: "go install"}},
}

func BuildTaskFile(tasks []struct {
	Alias string
	Task  Task
}) error {
	t := map[string]Task{}
	for _, val := range tasks {
		t[val.Alias] = val.Task
	}

	bytes_out, err := yaml.Marshal(TaskFile{Tasks: t})
	if err != nil {
		return err
	}

	task_file, err := os.Create(TaskFileStr)
	if err != nil {
		return err
	}

	if _, err := task_file.Write(bytes_out); err != nil {
		return err
	}
	return nil
}
