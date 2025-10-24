package app

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

type TodoCsvRepository struct {
	mapper *TodoCsvMapper
}

func NewTodoCsvRepository() *TodoCsvRepository {
	return &TodoCsvRepository{
		mapper: NewTodoCsvMapper(),
	}
}

func (t TodoCsvRepository) createTodo(todo *Todo) (*Todo, error) {
	filePath := filepath.Join("app", "todo.csv")

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("failed to create file: %v", err)
	}

	defer func() {
		file.Close()
	}()

	w := csv.NewWriter(file)

	id := uuid.NewString()
	todo.Id = id

	todoCsv := t.mapper.toRecord(todo)

	if err := w.Write(todoCsv); err != nil {
		todo.Id = ""
		return nil, err
	}

	w.Flush()

	if err := w.Error(); err != nil {
		todo.Id = ""
		return nil, err
	}

	return todo, nil
}

func (t TodoCsvRepository) updateTodo(todo *Todo) (*Todo, error) {
	innerTodo, err := t.getTodo(todo.Id)

	if err != nil {
		return nil, err
	}

	innerTodo.Name = todo.Name
	innerTodo.Description = todo.Description
	innerTodo.Status = todo.Status

	return innerTodo, nil
}

func (t TodoCsvRepository) getTodo(id string) (*Todo, error) {
	file, err := os.Open("todo.csv")
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	_, err = reader.Read()
	if err != nil && err != io.EOF {
		return nil, fmt.Errorf("failed to read header: %w", err)
	}

	idColumnIndex := 0

	for {
		record, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, fmt.Errorf("error reading CSV record: %w", err)
		}

		if idColumnIndex < len(record) {
			if record[idColumnIndex] == id {
				return t.mapper.toTodo(record), nil
			}
		} else {
			fmt.Printf("Warning: Record is too short for ID index %d: %v\n", idColumnIndex, record)
		}
	}

	return nil, fmt.Errorf("record with ID '%s' not found", id)
}

func (t TodoCsvRepository) getAllTodos() (*[]Todo, error) {
	//TODO implement me
	panic("implement me")
}

func (t TodoCsvRepository) deleteTodo(id int) error {
	//TODO implement me
	panic("implement me")
}

func _() {
	var _ TodoRepositoryInterface = (*TodoCsvRepository)(nil)
}
