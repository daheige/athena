// Package rpc for grpc service method impl.
package rpc

import (
	"context"
	"errors"
	"log"

	"github.com/daheige/gmicro/v2"

	"github.com/daheige/athena/internal/application"
	"github.com/daheige/athena/internal/pb"
)

// NewGreeterService 创建greeter service实例
func NewGreeterService(userService *application.UserService) pb.GreeterServiceServer {
	return &greeterService{
		userService: userService,
	}
}

// rpc service entry
type greeterService struct {
	// 这里必须包含这个解构体才可以，否则就是没有实现
	pb.UnimplementedGreeterServiceServer

	userService *application.UserService
}

// SayHello 返回 name,message
func (s *greeterService) SayHello(ctx context.Context, in *pb.HelloReq) (*pb.HelloReply, error) {
	log.Println("req data: ", in)
	md := gmicro.GetIncomingMD(ctx)
	log.Println("request md: ", md)
	if in.Id == 0 {
		return nil, errors.New("id invalid")
	}

	// 获取用户信息
	user, err := s.userService.GetUser(ctx, in.Id)
	if err != nil {
		return nil, err
	}

	reply := &pb.HelloReply{
		Name:    user.User,
		Message: "hello," + user.Name,
	}

	return reply, nil
}

// Info 返回address,message
func (s *greeterService) Info(ctx context.Context, in *pb.InfoReq) (*pb.InfoReply, error) {
	log.Println("request name: ", in.Name)

	return &pb.InfoReply{
		Address: "shenzhen",
		Message: "ok",
	}, nil
}

// BatchUsers 批量获取用户信息
func (s *greeterService) BatchUsers(ctx context.Context, in *pb.BatchUsersReq) (*pb.BatchUsersReply, error) {
	log.Println("request data: ", in)
	users, err := s.userService.BatchUsers(ctx, in.Ids)
	if err != nil {
		return nil, err
	}

	reply := &pb.BatchUsersReply{
		Users: make([]*pb.UserEntity, 0, len(users)+1),
	}
	for k := range users {
		reply.Users = append(reply.Users, &pb.UserEntity{
			Id:   users[k].ID,
			Name: users[k].Name,
		})
	}

	return reply, nil
}
