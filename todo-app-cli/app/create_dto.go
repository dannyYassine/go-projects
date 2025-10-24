package app

type CreateTodoDto struct {
	Name        string
	Description string
	Status      TodoStatus
}

func NewCreateTodoDto(name string) *CreateTodoDto {
	return &CreateTodoDto{Name: name, Status: New}
}
