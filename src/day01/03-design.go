package main

import (
	"net/rpc"
	"net/rpc/jsonrpc"
)

//要求，服务端注册rpc对象，能编译期检出 注册对象是否合法

//创建接口，在接口定义方法（多态）
type MyInterface interface {
	//绑定的结构体三个参数，传入和传出
	HelloWorld(string, *string) error
}

//	定义个函数，封装函数
func RegisterService(i MyInterface) {
	//使用的
	rpc.RegisterName("helloDoge", i)
}

//-----客户端重载-----
//创建一个类
type Myclient struct {
	c *rpc.Client
}

//	实现函数，原型参照上面的interface来实现
func (this *Myclient) HelloWorld(a string, b *string) error {
	//参数一 参照上面的Interface， RegisterName 而来。 a,传入参数，b传出参数
	return this.c.Call("helloDoge.HelloWorld", a, b)
}

//初始化客户端
func IntiClient(addr string) Myclient {
	conn, _ := jsonrpc.Dial("tcp", addr)
	return Myclient{c: conn}
}
