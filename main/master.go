package main

import (
	"mit_6.824/mr"
	"sync"
)

type Master struct {
	mu      sync.Mutex
	tasks   []mr.Task // 使用 mr 包定义的 Task 结构
	nReduce int       // Reduce 任务的数量
}

// MakeMaster 初始化 Master 的实例
func MakeMaster(files []string, nReduce int) *Master {
	m := Master{
		tasks:   make([]mr.Task, len(files)),
		nReduce: nReduce,
	}

	// 初始化 Map 任务
	for i, filename := range files {
		m.tasks[i] = mr.Task{
			TaskType: mr.MapTask,
			TaskID:   i,
			FileName: filename,
			Status:   mr.Idle,
		}
	}
	return &m
}
