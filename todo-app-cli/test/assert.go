package test

import (
	"log"
	"os"
	"testing"
	"todo-app-cli/app"

	"github.com/stretchr/testify/assert"
)

func Equal(t *testing.T, a, b interface{}) {
	if a != b {
		t.Errorf(`Assert failed: %q not equal to %q`, a, b)
	}
}

func AssertCsvContainsTodo(t *testing.T, todo *app.Todo) {
	fileContent := loadFile()
	assert.Contains(t, fileContent, todo.Id)
}

func AssertCsvNotContainsTodo(t *testing.T, todo *app.Todo) {
	fileContent := loadFile()
	assert.NotContains(t, fileContent, todo.Id)
}

func loadFile() string {
	filePath := "../todo.csv"

	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	return string(content)
}
