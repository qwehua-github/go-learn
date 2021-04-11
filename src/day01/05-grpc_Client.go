package main

import (
	"context"
	"day01/pb"
	"fmt"
	"google.golang.org/grpc"
)

func main() {
	//1. 连接grpc服务,参数选择安全的
	grpcConn, _ := grpc.Dial("127.0.0.1:8800", grpc.WithInsecure())
	defer grpcConn.Close()
	//2. 初始化grpc的客户端
	grpcClient := pb.NewSayNameClient(grpcConn)

	//创建一个Techer对象
	var teacher pb.Teacher
	teacher.Name = "doge"
	teacher.Age = 18

	//3. 调用远程服务
	//context.TODO是一个空对象
	t, _ := grpcClient.SayHello(context.TODO(), &teacher)

	fmt.Println(t)
}
