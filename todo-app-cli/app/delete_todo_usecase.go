package app

import "fmt"

type DeleteTodoUseCase struct {
	todoRepository TodoRepositoryInterface
}

func NewDeleteTodoUseCase(repositoryInterface TodoRepositoryInterface) *DeleteTodoUseCase {
	return &DeleteTodoUseCase{
		todoRepository: repositoryInterface,
	}
}

func (useCase *DeleteTodoUseCase) Execute(dto *DeleteTodoDto) error {
	err := useCase.todoRepository.deleteTodo(dto.Id)

	if err != nil {
		return fmt.Errorf("could not get todos: %w", err)
	}

	return nil
}
