package app

type TodoRepositoryInterface interface {
	createTodo(todo *Todo) (*Todo, error)
	updateTodo(todo *Todo) (*Todo, error)
	getTodo(id string) (*Todo, error)
	getAllTodos() (*[]Todo, error)
	deleteTodo(id string) error
}
