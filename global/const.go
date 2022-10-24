package global

type Status int

const (
	Todo Status = iota
	InProgress
	Done
)

const Divisor = 4

const StoragePath = "kanbancli"
const StorageName = StoragePath + "/kanbanclistorage.json"
