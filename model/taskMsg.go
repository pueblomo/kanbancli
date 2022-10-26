package model

import "pueblomo/kanbancli/global"

type TaskMsg struct {
	Type global.MsgType
	Task Task
}

func NewTaskMsg(msgType global.MsgType, task Task) TaskMsg {
	return TaskMsg{Type: msgType, Task: task}
}
