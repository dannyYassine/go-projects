package app

type UpdateTodoDto struct {
	Id          string
	Name        string
	Description string
	Status      TodoStatus
}

func NewUpdateTodoDto(id string, name string, description string, status string) *UpdateTodoDto {
	return &UpdateTodoDto{Id: id, Name: name, Description: description, Status: NewTodoStatus(status)}
}
