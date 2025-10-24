package app

import "fmt"

type CreateTodoUseCase struct {
	todoRepository TodoRepositoryInterface
}

func NewCreateTodoUseCase() *CreateTodoUseCase {
	return &CreateTodoUseCase{
		todoRepository: NewTodoCsvRepository(),
	}
}

func (useCase *CreateTodoUseCase) Execute(dto CreateTodoDto) (*Todo, error) {
	todo := &Todo{Name: dto.Name}

	savedTodo, err := useCase.todoRepository.createTodo(todo)

	if err != nil {
		return nil, fmt.Errorf("could not create todo: %w", err)
	}

	return savedTodo, nil
}
