package main

import (
	"context" //上下文 goroutine
	"day01/pb"
	"google.golang.org/grpc"
	"net"
)

//定义类
type Children struct {
}

//从pb下的proto.go绑定方法
func (this *Children) SayHello(ctx context.Context, t *pb.Teacher) (*pb.Teacher, error) {
	t.Name += "is Sleeping"
	t.Age += 2
	return t, nil
}

func main() {
	//1. 初始化一个grpc对象
	grpcServer := grpc.NewServer()
	//2. 注册服务
	pb.RegisterSayNameServer(grpcServer, new(Children))
	//3. 设置监听，指定IP，port
	listener, _ := net.Listen("tpc", "127.0.0.1:8800")
	defer listener.Close()
	//4. 启动服务。 --server()
	grpcServer.Serve(listener)
}
