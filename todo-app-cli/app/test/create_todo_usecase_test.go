package test

import (
	"fmt"
	"testing"
	"todo-app-cli/app"
	"todo-app-cli/test"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setUpFunctionalTest(_ *testing.T) (*app.Application, func(t *testing.T)) {
	application := app.NewApplication()
	application.Bootstrap()

	return application, func(t *testing.T) {
		application.Shutdown()
	}
}

func setUpIntegrationTest(_ *testing.T) (*app.Application, func(t *testing.T)) {
	application := app.NewApplication()
	application.Bootstrap()

	return application, func(t *testing.T) {
		application.Shutdown()
	}
}

func Test_CreateTodoUseCase_CreatesTodo(t *testing.T) {
	application, suite := setUpFunctionalTest(t)
	defer suite(t)

	dto := app.NewCreateTodoDto("name", "test")

	useCase := app.Get[app.CreateTodoUseCase](application)
	todo, err := useCase.Execute(dto)

	assert.NoError(t, err)
	assert.Equal(t, todo.Name, "name")
	assert.Equal(t, todo.Description, "test")
	assert.NotNil(t, todo.Id)

	test.AssertCsvContainsTodo(t, todo)
}

func Test_CreateTodoUseCase_CreatesTodoMocked(t *testing.T) {
	application, suite := setUpFunctionalTest(t)
	defer suite(t)

	dto := app.NewCreateTodoDto("name", "test")

	expectedTodo := &app.Todo{Id: "generated-uuid-123", Name: "name", Description: "test", Status: "new"}

	mocked := new(MyMockedObject)
	mocked.On("CreateTodo", mock.Anything).Return(expectedTodo, nil).Once()

	application.Container.PartialMock(func() app.TodoRepositoryInterface {
		return mocked
	})

	useCase := app.Get[app.CreateTodoUseCase](application)
	todo, err := useCase.Execute(dto)

	mocked.AssertCalled(t, "CreateTodo", mock.MatchedBy(func(todo *app.Todo) bool {
		return todo.Status == app.New && todo.Id == ""
	}))
	assert.NoError(t, err)
	assert.Equal(t, todo.Id, expectedTodo.Id)
	assert.Equal(t, todo.Name, "name")
	assert.Equal(t, todo.Description, "test")
	assert.Equal(t, todo.Status, app.New)
	assert.NotNil(t, todo.Id)
}

func Test_CreateTodoUseCase_HandlesError(t *testing.T) {
	application, suite := setUpIntegrationTest(t)
	defer suite(t)

	dto := app.NewCreateTodoDto("name", "test")

	mocked := new(MyMockedObject)
	mocked.On("CreateTodo", mock.Anything).Return((*app.Todo)(nil), fmt.Errorf("cant create todo")).Once()

	application.Container.PartialMock(func() app.TodoRepositoryInterface {
		return mocked
	})

	useCase := app.Get[app.CreateTodoUseCase](application)
	todo, err := useCase.Execute(dto)

	assert.Nil(t, todo)
	assert.Error(t, err)

}
