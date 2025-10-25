package app

import "fmt"

type ListTodosUseCase struct {
	todoRepository TodoRepositoryInterface
}

func NewListTodosUseCase(repositoryInterface TodoRepositoryInterface) *ListTodosUseCase {
	return &ListTodosUseCase{
		todoRepository: repositoryInterface,
	}
}

func (useCase *ListTodosUseCase) Execute() (*[]Todo, error) {
	todos, err := useCase.todoRepository.getAllTodos()

	if err != nil {
		return nil, fmt.Errorf("could not get todos: %w", err)
	}

	return todos, nil
}
