package test

import (
	"todo-app-cli/app"

	"github.com/stretchr/testify/mock"
)

type MyMockedObject struct {
	app.TodoCsvRepository
	mock.Mock
}

func (m *MyMockedObject) CreateTodo(todo *app.Todo) (*app.Todo, error) {
	args := m.Called(todo)
	return args.Get(0).(*app.Todo), args.Error(1)
}
