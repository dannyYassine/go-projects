package app

type DeleteTodoDto struct {
	Id string
}

func NewDeleteTodoDto(id string) *DeleteTodoDto {
	return &DeleteTodoDto{Id: id}
}
