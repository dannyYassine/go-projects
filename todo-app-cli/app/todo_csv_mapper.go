package app

type TodoCsvMapper struct {
}

func NewTodoCsvMapper() *TodoCsvMapper {
	return &TodoCsvMapper{}
}

func (m *TodoCsvMapper) toTodo(record []string) *Todo {
	return &Todo{Id: record[0], Name: record[1], Description: record[2], Status: TodoStatus(record[3])}
}

func (m *TodoCsvMapper) toRecord(todo *Todo) []string {
	return []string{todo.Id, todo.Name, todo.Description, string(todo.Status)}
}
