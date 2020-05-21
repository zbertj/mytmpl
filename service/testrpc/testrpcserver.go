package testrpc

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "mytmpl/service/testrpc/proto"

	"google.golang.org/grpc"
)

// UserInfoService 定义一个对象
type UserInfoService struct {
}

// NewUserInfoService 创建一个grpc服务
func NewUserInfoService() *UserInfoService {
	return &UserInfoService{}
}

// Run  NewUserInfoService这个服务的启动函数
func (s *UserInfoService) Run(grpcHost string, grpcPort int32) error {
	srv := grpc.NewServer()
	pb.RegisterUserInfoServiceServer(srv, s)

	// 提供对外接口
	lis, err := net.Listen("tcp4", fmt.Sprintf("%s:%d", grpcHost, grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return err
	}

	return srv.Serve(lis)
}

// GetUserInfo 向外提供的接口方法
func (s *UserInfoService) GetUserInfo(ctx context.Context, req *pb.UserRequest) (res *pb.UserResponse, err error) {
	fmt.Println(req.String())

	res = &pb.UserResponse{
		Id:    1,
		Name:  req.GetName(),
		Age:   22,
		Hobby: []string{"唱歌", "跳舞"},
	}
	err = nil
	return
}

/*
import (
	"fmt"
	"mytmpl/service/testrpc"
)

func main() {
	fmt.Println("--------------running-----------------")

	testrpc := testrpc.NewUserInfoService()
	testrpc.Run("127.0.0.1", 8080)
	fmt.Println("rpc  is   running ")
}


*/
