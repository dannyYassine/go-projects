package app

type CreateTodoDto struct {
	Name        string
	Description string
	Status      TodoStatus
}

func NewCreateTodoDto(name string, description string) *CreateTodoDto {
	return &CreateTodoDto{Name: name, Description: description, Status: New}
}
