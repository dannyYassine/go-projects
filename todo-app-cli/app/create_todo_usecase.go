package app

import "fmt"

type CreateTodoUseCase struct {
	todoRepository TodoRepositoryInterface
}

func NewCreateTodoUseCase(repositoryInterface TodoRepositoryInterface) *CreateTodoUseCase {
	return &CreateTodoUseCase{
		todoRepository: repositoryInterface,
	}
}

func (useCase *CreateTodoUseCase) Execute(dto *CreateTodoDto) (*Todo, error) {
	todo := Todo{Name: dto.Name, Description: dto.Description, Status: dto.Status}

	savedTodo, err := useCase.todoRepository.createTodo(&todo)

	if err != nil {
		return nil, fmt.Errorf("could not create todo: %w", err)
	}

	return savedTodo, nil
}
