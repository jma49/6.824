package mr

type RequestTaskArgs struct{}

type RequestTaskReply struct {
	Task Task
}

type ReportTaskStatusArgs struct {
	TaskID     int
	TaskStatus string
}

type ReportTaskStatusReply struct{}
