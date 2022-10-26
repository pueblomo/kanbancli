package global

type Status int

const (
	Todo Status = iota
	InProgress
	Done
)

type MsgType int

const (
	Create MsgType = iota
	Update
	Show
)

const Divisor = 4

const StoragePath = "kanbancli"
const StorageName = StoragePath + "/kanbanclistorage.json"
