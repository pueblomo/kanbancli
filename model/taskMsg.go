package model

type TaskMsg struct {
	Create bool
	Task   Task
}

func NewTaskMsg(create bool, task Task) TaskMsg {
	return TaskMsg{Create: create, Task: task}
}
