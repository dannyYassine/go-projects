package app

type Todo struct {
	Id          string
	Name        string
	Description string
	Status      TodoStatus
}

func NewTodo() *Todo {
	return &Todo{Status: InProgress}
}
