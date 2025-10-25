package test

import (
	"testing"
	"todo-app-cli/app"
	"todo-app-cli/test"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateTodoUseCaseCreatesTodo(t *testing.T) {
	app.Application.Register().Boot()

	dto := app.NewCreateTodoDto("name", "test")

	useCase := app.Get[app.CreateTodoUseCase]()
	todo, err := useCase.Execute(dto)

	assert.NoError(t, err)
	assert.Equal(t, todo.Name, "name")
	assert.Equal(t, todo.Description, "test")
	assert.NotNil(t, todo.Id)

	test.AssertCsvContainsTodo(t, todo)
}

func TestCreateTodoUseCaseCreatesTodoMocked(t *testing.T) {
	app.Application.Register()
	mocked := new(MyMockedObject)
	dto := app.NewCreateTodoDto("name", "test")

	expectedTodo := &app.Todo{Id: "generated-uuid-123", Name: "name", Description: "test", Status: "new"}

	mocked.On("CreateTodo", mock.Anything).Return(expectedTodo, nil).Once()
	app.Application.Container.PartialMock(func() app.TodoRepositoryInterface {
		return mocked
	})
	app.Application.Boot()

	useCase := app.Get[app.CreateTodoUseCase]()
	todo, err := useCase.Execute(dto)

	assert.NoError(t, err)
	assert.Equal(t, todo.Name, "name")
	assert.Equal(t, todo.Description, "test")
	assert.NotNil(t, todo.Id)
}

type MyMockedObject struct {
	app.TodoCsvRepository
	mock.Mock
}

func (m *MyMockedObject) CreateTodo(todo *app.Todo) (*app.Todo, error) {
	args := m.Called(todo)
	return args.Get(0).(*app.Todo), args.Error(1)
}
