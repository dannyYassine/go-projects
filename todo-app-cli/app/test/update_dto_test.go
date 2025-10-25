package test

import (
	"testing"
	"todo-app-cli/app"
	"todo-app-cli/test"

	"github.com/stretchr/testify/assert"
)

func TestNewUpdateTodoDto(t *testing.T) {
	dto := app.NewUpdateTodoDto("1", "test", "test", "new")

	test.Equal(t, dto.Id, "1")
	assert.Equal(t, dto.Name, "test")
	test.Equal(t, dto.Description, "test")
	test.Equal(t, dto.Status, app.TodoStatus("new"))
}
