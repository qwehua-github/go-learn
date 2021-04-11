package main

import (
	"fmt"
	"net"
	"net/rpc/jsonrpc"
)

//创建一个类
type World struct {
}

//绑定类的方法
func (this *World) HelloWorld(name string, resp *string) error {
	*resp = name + "你好！"
	return nil
	//return errors.New("未知错误")
}

func main() {

	// 1. 注册RPC服务
	RegisterService(new(World))
	//err := rpc.RegisterName("helloDoge", new(World))
	//if err != nil {
	//	fmt.Println("注册 rpc 服务失败！", err)
	//	return
	//}
	// 2. 设置监听
	listener, err := net.Listen("tcp", "127.0.0.1:8800")
	if err != nil {
		fmt.Println("设置监听失败！", err)
		return
	}
	defer listener.Close()
	fmt.Printf("开始监听...")

	for {
		// 3. 建立链接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("net.Accept err:", err)
			return
		}
		defer conn.Close()
		fmt.Println("链接成功...")
		// 4. 绑定服务
		//rpc.ServeConn(conn)
		//jsonrpc绑定服务
		jsonrpc.ServeConn(conn)
		fmt.Println("绑定成功...")
	}
}
