package app

import (
	"fmt"
	"strings"

	"github.com/rodaine/table"
)

type ConsoleRenderer struct{}

func NewConsoleRenderer() *ConsoleRenderer {
	return &ConsoleRenderer{}
}

func (cr *ConsoleRenderer) PrintTodo(todo *Todo) {
	if todo == nil {
		return
	}

	cr.print(&[]Todo{*todo})
}

func (cr *ConsoleRenderer) PrintTodos(todos *[]Todo) {
	cr.print(todos)
}

func (cr *ConsoleRenderer) print(todos *[]Todo) {
	table.DefaultHeaderFormatter = func(format string, vals ...interface{}) string {
		return strings.ToUpper(fmt.Sprintf(format, vals...))
	}

	tbl := table.New("ID", "Name", "Description", "Status")
	tbl.WithHeaderSeparatorRow(1)
	for _, todo := range *todos {
		tbl.AddRow(todo.Id, todo.Name, todo.Description, todo.Status)
	}
	tbl.WithHeaderSeparatorRow(1)

	tbl.Print()
}
