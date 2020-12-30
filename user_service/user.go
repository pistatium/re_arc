package main

import (
	"context"
	"fmt"
	pb "github.com/pistatium/re_ark/protos"
)

type UserService struct {}

func (u UserService) GetUser(ctx context.Context, request *pb.GetUserRequest) (*pb.UserResponse, error) {
	userId := request.UserId
	return &pb.UserResponse{
		Id: userId,
		Name: fmt.Sprintf("user-%s", userId),
	}, nil
}


