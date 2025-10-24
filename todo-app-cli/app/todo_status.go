package app

type TodoStatus string

const (
	New        TodoStatus = "new"
	InProgress TodoStatus = "in-progress"
	Completed  TodoStatus = "completed"
)

func (status TodoStatus) PrettyPrintString() string {
	switch status {
	case New:
		return "New"
	case InProgress:
		return "In-progress"
	case Completed:
	default:
		return "Completed"
	}

	return "<nil>"
}
