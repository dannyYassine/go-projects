package app

import "fmt"

type UpdateTodoUseCase struct {
	todoRepository TodoRepositoryInterface
}

func NewUpdateTodoUseCase(repositoryInterface TodoRepositoryInterface) *UpdateTodoUseCase {
	return &UpdateTodoUseCase{
		todoRepository: repositoryInterface,
	}
}

func (useCase *UpdateTodoUseCase) Execute(dto *UpdateTodoDto) (*Todo, error) {
	savedTodo, err := useCase.todoRepository.GetTodo(dto.Id)

	if err != nil {
		return nil, fmt.Errorf("Could not get todo: %w", err)
	}

	if name := dto.Name; name != "" {
		savedTodo.Name = dto.Name
	}
	if description := dto.Description; description != "" {
		savedTodo.Description = dto.Description
	}
	if status := dto.Status; status != "" {
		savedTodo.Status = dto.Status
	}

	newTodo, err := useCase.todoRepository.UpdateTodo(savedTodo)

	if err != nil {
		return nil, fmt.Errorf("could not create todo: %w", err)
	}

	return newTodo, nil
}
