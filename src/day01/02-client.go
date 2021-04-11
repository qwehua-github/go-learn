package main

import (
	"fmt"
)

//func main01() {
//	// 1. 用rpc链接服务器
//	//conn, err := rpc.Dial("tcp", "127.0.0.1:8800")
//	//json版本的cli
//	conn, err := jsonrpc.Dial("tcp", "127.0.0.1:8800")
//	if err != nil {
//		fmt.Println("Dial err:", err)
//		return
//	}
//	defer conn.Close()
//	fmt.Println("rpc链接服务器")
//	// 2. 调用远程函数
//	var reply string // 接受返回值 ， ——传出参数
//	err = conn.Call("helloDoge.HelloWorld", "傻东", &reply)
//	if err != nil {
//		fmt.Println("Call() err:", err)
//		return
//	}
//
//	fmt.Println("远程参数返回：", reply)
//}

func main() {
	myClient := IntiClient("127.0.0.1:8800")

	var resp string
	err := myClient.HelloWorld("韭神", &resp)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("返回值为", resp)
}
