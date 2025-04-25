package task

import (
	"os"
	"strings"
	"testing"
)

func TestBuildTaskFile(t *testing.T) {
	// Clean up before test
	_ = os.Remove(".task.yaml")

	// Define test input
	tasks := []struct {
		Alias string
		Task  Task
	}{
		{
			Alias: "build",
			Task:  Task{Desc: "Build the project", Run: "go build ."},
		},
		{
			Alias: "test",
			Task:  Task{Desc: "Run tests", Run: "go test ./..."},
		},
	}

	// Run the function
	err := BuildTaskFile(tasks)
	if err != nil {
		t.Fatalf("BuildTaskFile failed: %v", err)
	}

	// Check file existence
	_, err = os.Stat(".task.yaml")
	if err != nil {
		t.Fatalf(".task.yaml was not created: %v", err)
	}

	// Read file content
	content, err := os.ReadFile(".task.yaml")
	if err != nil {
		t.Fatalf("Failed to read .task.yaml: %v", err)
	}

	// Basic content checks
	strContent := string(content)
	for _, task := range tasks {
		if !strings.Contains(strContent, task.Alias) ||
			!strings.Contains(strContent, task.Task.Desc) ||
			!strings.Contains(strContent, task.Task.Run) {
			t.Errorf("Generated content missing task: %s\nContent:\n%s", task.Alias, strContent)
		}
	}

	// Clean up after test
	_ = os.Remove(".task.yaml")
}

