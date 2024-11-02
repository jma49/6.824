package main

import (
	"mit_6.824/mr"
)

func Worker(mapF func(string, string) []mr.KeyValue, reduceF func(string, []string) string) {
	// Worker 请求任务并执行逻辑
}
