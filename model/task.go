package model

import "pueblomo/kanbancli/global"

type Task struct {
	Status                  global.Status
	title, tag, description string
}

func (t Task) Title() string {
	return t.title
}

func (t Task) Description() string {
	return t.tag
}

func (t Task) Tag() string {
	return t.tag
}

func (t Task) ShowDescription() string {
	return t.description
}

func (t Task) FilterValue() string {
	return t.title
}

func New(title, tag, description string) Task {
	return Task{Status: global.Todo, title: title, tag: tag, description: description}
}

func (t *Task) Next() {
	if t.Status == global.Done {
		t.Status = global.Todo
	} else {
		t.Status++
	}
}
