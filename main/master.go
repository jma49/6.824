package main

import (
	"6.824/mr"
	"sync"
)

type Master struct {
	mu      sync.Mutex
	tasks   []mr.Task
	nReduce int
}

func (m *Master) AssighTask(args *mr.RequestTaskArgs, reply *mr.RequestTaskReply) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	//
	return nil
}

func MakeMaster(files []string, nReduce int) *Master {
	m := Master{
		tasks:   []mr.Task{},
		nReduce: nReduce,
	}
	// 初始化任务队列
	return &m
}
