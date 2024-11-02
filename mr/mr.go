package mr

// TaskType 定义 TaskType 枚举，用于标识任务类型
type TaskType int

const (
	MapTask    TaskType = iota // Map 任务类型
	ReduceTask                 // Reduce 任务类型
)

// TaskStatus 定义 TaskStatus 枚举，用于标识任务状态
type TaskStatus int

const (
	Idle       TaskStatus = iota // 任务未开始
	InProgress                   // 任务正在进行
	Completed                    // 任务已完成
)

// Task 定义 Task 结构体，用于描述每个任务
type Task struct {
	TaskType TaskType   // 任务类型 (Map 或 Reduce)
	TaskID   int        // 任务 ID
	FileName string     // 输入文件名 (仅对 Map 任务有意义)
	Status   TaskStatus // 任务状态
}

// KeyValue 定义 KeyValue 结构体，用于存储 mapF 的输出键值对
type KeyValue struct {
	Key   string
	Value string
}
