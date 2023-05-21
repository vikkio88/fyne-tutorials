package models

import (
	"fmt"

	"fyne.io/fyne/v2/data/binding"
)

type Todo struct {
	Description string
	Done        bool
}

func NewTodo(description string) Todo {
	return Todo{description, false}
}

func NewTodoFromDataItem(di binding.DataItem) Todo {
	v, _ := di.(binding.Untyped).Get()
	return v.(Todo)
}

func (t Todo) String() string {
	return fmt.Sprintf("%s  - %t", t.Description, t.Done)
}
