package test

import (
	"testing"
	"todo-app-cli/app"
	"todo-app-cli/test"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setUpFunctionalTest() {
	app.Application.Bootstrap()
}

func Test_CreateTodoUseCase_CreatesTodo(t *testing.T) {
	setUpFunctionalTest()

	dto := app.NewCreateTodoDto("name", "test")

	useCase := app.Get[app.CreateTodoUseCase]()
	todo, err := useCase.Execute(dto)

	assert.NoError(t, err)
	assert.Equal(t, todo.Name, "name")
	assert.Equal(t, todo.Description, "test")
	assert.NotNil(t, todo.Id)

	test.AssertCsvContainsTodo(t, todo)
}

func Test_CreateTodoUseCase_CreatesTodoMocked(t *testing.T) {
	setUpFunctionalTest()

	dto := app.NewCreateTodoDto("name", "test")

	expectedTodo := &app.Todo{Id: "generated-uuid-123", Name: "name", Description: "test", Status: "new"}

	mocked := new(MyMockedObject)
	mocked.On("CreateTodo", mock.Anything).Return(expectedTodo, nil).Once()

	app.Application.Container.PartialMock(func() app.TodoRepositoryInterface {
		return mocked
	})

	useCase := app.Get[app.CreateTodoUseCase]()
	todo, err := useCase.Execute(dto)

	mocked.AssertCalled(t, "CreateTodo", mock.MatchedBy(func(todo *app.Todo) bool {
		return todo.Status == app.New
	}))
	assert.NoError(t, err)
	assert.Equal(t, todo.Name, "name")
	assert.Equal(t, todo.Description, "test")
	assert.Equal(t, todo.Status, app.New)
	assert.NotNil(t, todo.Id)
}
