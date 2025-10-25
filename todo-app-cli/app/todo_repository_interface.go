package app

type TodoRepositoryInterface interface {
	CreateTodo(todo *Todo) (*Todo, error)
	UpdateTodo(todo *Todo) (*Todo, error)
	GetTodo(id string) (*Todo, error)
	GetAllTodos() (*[]Todo, error)
	DeleteTodo(id string) error
}
