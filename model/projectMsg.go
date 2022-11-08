package model

import "pueblomo/kanbancli/global"

type ProjectMsg struct {
	Type    global.MsgType
	Project string
}

func NewProjectMsg(msgType global.MsgType, project string) ProjectMsg {
	return ProjectMsg{Type: msgType, Project: project}
}
