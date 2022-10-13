package item

import "pueblomo/kanbancli/global"

type Task struct {
	Status             global.Status
	title, description string
}

func (t Task) Title() string {
	return t.title
}

func (t Task) Description() string {
	return t.description
}

func (t Task) FilterValue() string {
	return t.title
}

func New(title, description string) Task {
	return Task{Status: global.Todo, title: title, description: description}
}

func (t *Task) Next() {
	if t.Status == global.Done {
		t.Status = global.Todo
	} else {
		t.Status++
	}
}
