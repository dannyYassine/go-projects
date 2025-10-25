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

func NewTodoRepositoryInterface() TodoRepositoryInterface {
	return NewTodoCsvRepository()
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
	// Ensure the record exists first
	if todo == nil || todo.Id == "" {
		return nil, fmt.Errorf("invalid todo: missing id")
	}

	filePath := filepath.Join("app", "todo.csv")

	// Open existing file for reading
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	// Allow variable number of fields across records (header vs data rows)
	reader.FieldsPerRecord = -1

	// Read and preserve header (if present)
	header, err := reader.Read()
	if err != nil && err != io.EOF {
		return nil, fmt.Errorf("failed to read header: %w", err)
	}

	records := make([][]string, 0, 16)
	found := false

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error reading CSV record: %w", err)
		}

		// Guard against short/invalid rows (e.g., header without ID)
		if len(record) == 0 {
			records = append(records, record)
			continue
		}

		if !found && record[0] == todo.Id {
			// Replace with updated record
			records = append(records, t.mapper.toRecord(todo))
			found = true
			continue
		}

		records = append(records, record)
	}

	if !found {
		return nil, fmt.Errorf("record with ID '%s' not found", todo.Id)
	}

	// Rewrite the file with header and updated records
	out, err := os.Create(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file for writing: %w", err)
	}
	defer out.Close()

	w := csv.NewWriter(out)

	// If there was a header, write it back as-is
	if len(header) > 0 {
		if err := w.Write(header); err != nil {
			return nil, fmt.Errorf("failed to write header: %w", err)
		}
	}

	if err := w.WriteAll(records); err != nil {
		return nil, fmt.Errorf("failed to write records: %w", err)
	}
	w.Flush()
	if err := w.Error(); err != nil {
		return nil, fmt.Errorf("csv writer error: %w", err)
	}

	return todo, nil
}

func (t TodoCsvRepository) getTodo(id string) (*Todo, error) {
	filePath := filepath.Join("app", "todo.csv")
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	// Allow variable number of fields across records (header vs data rows)
	reader.FieldsPerRecord = -1

	// skip header if present
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
	filePath := filepath.Join("app", "todo.csv")
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	// Allow variable number of fields across records (header vs data rows)
	reader.FieldsPerRecord = -1

	// skip header if present
	_, err = reader.Read()
	if err != nil && err != io.EOF {
		return nil, fmt.Errorf("failed to read header: %w", err)
	}

	var allTodos []Todo

	for {
		record, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, fmt.Errorf("error reading CSV record: %w", err)
		}
		todo := t.mapper.toTodo(record)
		allTodos = append(allTodos, *todo)
	}

	return &allTodos, nil
}

func (t TodoCsvRepository) deleteTodo(id string) error {
	// Ensure the record exists first
	if id == "" {
		return fmt.Errorf("invalid todo: missing id")
	}

	filePath := filepath.Join("app", "todo.csv")

	// Open existing file for reading
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	// Allow variable number of fields across records (header vs data rows)
	reader.FieldsPerRecord = -1

	// Read and preserve header (if present)
	header, err := reader.Read()
	if err != nil && err != io.EOF {
		return fmt.Errorf("failed to read header: %w", err)
	}

	records := make([][]string, 0, 16)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("error reading CSV record: %w", err)
		}

		// Guard against short/invalid rows (e.g., header without ID)
		if len(record) == 0 {
			records = append(records, record)
			continue
		}

		if record[0] == id {
			continue
		}

		records = append(records, record)
	}

	// Rewrite the file with header and updated records
	out, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file for writing: %w", err)
	}
	defer out.Close()

	w := csv.NewWriter(out)

	// If there was a header, write it back as-is
	if len(header) > 0 {
		if err := w.Write(header); err != nil {
			return fmt.Errorf("failed to write header: %w", err)
		}
	}

	if err := w.WriteAll(records); err != nil {
		return fmt.Errorf("failed to write records: %w", err)
	}
	w.Flush()
	if err := w.Error(); err != nil {
		return fmt.Errorf("csv writer error: %w", err)
	}

	return nil
}

func _() {
	var _ TodoRepositoryInterface = (*TodoCsvRepository)(nil)
}
