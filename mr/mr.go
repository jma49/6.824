package mr

type TaskType int

const (
	MapTask TaskType = iota
	ReduceTask
)

type TaskStatus int

const (
	Idle TaskStatus = iota
	InProgress
	Completed
)

type Task struct {
	TaskType TaskType
	TaskID   int
	FileName string
	NReduce  int
	Status   TaskStatus
}
