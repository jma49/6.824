package main

import (
	"log"
	"net"
	"net/rpc"
)

func coordinator(files []string, nReduce int) {
	master := MakeMaster(files, nReduce) // 调用 MakeMaster 初始化 Master

	err := rpc.Register(master) // 注册 RPC 服务并处理错误
	if err != nil {
		log.Fatalf("Failed to register RPC server: %v", err)
	}

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatalf("Listen error: %v", err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err == nil {
			go rpc.ServeConn(conn)
		}
	}
}
