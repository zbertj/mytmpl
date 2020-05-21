package main

import (
	"context"
	"fmt"
	"time"

	pb "mytmpl/service/testrpc/proto"

	"google.golang.org/grpc"
)

// InitRPC 初始化函数，创建一个连接 Set up a connection to the server.
func InitRPC() (pb.UserInfoServiceClient, error) {
	conn, err := grpc.Dial("127.0.0.1:8080",
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(256<<20)))
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	//defer conn.Close()
	rpc := pb.NewUserInfoServiceClient(conn)
	return rpc, nil
}

func main() {
	rpc, err := InitRPC()
	if err != nil {
		return
	}

	req := &pb.UserRequest{
		Name: "zbjjj",
	}

	for {
		fmt.Println("GetDataFromRpc")

		for i := 1; i < 10; i++ {
			res, err := rpc.GetUserInfo(context.Background(), req)
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Println(res)
		}
		time.Sleep(10 * time.Second)
	}
}
