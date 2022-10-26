package model

import "pueblomo/kanbancli/global"

type Task struct {
	Status  global.Status
	Subject string
	Tag     string
	Desc    string
}

func (t Task) Title() string {
	return t.Subject
}

func (t Task) Description() string {
	return t.Tag
}

func (t Task) FilterValue() string {
	return t.Subject + t.Tag
}

func New(title, tag, description string) Task {
	return Task{Status: global.Todo, Subject: title, Tag: tag, Desc: description}
}

func (t *Task) Next() {
	if t.Status == global.Done {
		t.Status = global.Todo
	} else {
		t.Status++
	}
}
