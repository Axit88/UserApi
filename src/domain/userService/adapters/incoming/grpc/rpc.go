package rpc

import (
	"context"
	"fmt"

	"github.com/Axit88/UserApi/src/domain/userService/adapters/incoming/grpc/pb"

	"github.com/Axit88/UserApi/src/domain/userService/core/model"
)

func (grpca Adapter) GetUser(ctx context.Context, req *pb.GetUserInput) (*pb.GetUserOutput, error) {
	res, err := grpca.api.ProcessGetUser(req.UserId)

	if err != nil {
		return nil, err
	}
	ans := pb.GetUserOutput{
		UserId:   res.UserId,
		UserName: res.UserName,
	}

	return &ans, nil
}

func (grpca Adapter) AddUser(ctx context.Context, req *pb.AddUserInput) (*pb.AddUserOutput, error) {

	input := model.User{
		UserId:   req.UserId,
		UserName: req.UserName,
	}
	err := grpca.api.ProcessAddUser(&input)

	if err != nil {
		return nil, err
	}
	output := pb.AddUserOutput{
		Message: fmt.Sprintf("User %v Added Successfully", req.UserName),
	}

	return &output, nil
}

func (grpca Adapter) DeleteUser(ctx context.Context, req *pb.DeleteUserInput) (*pb.DeleteUserOutput, error) {

	err := grpca.api.ProcessDeleteUser(req.UserId)

	if err != nil {
		return nil, err
	}
	output := pb.DeleteUserOutput{
		Message: fmt.Sprintf("User %v Deleted Successfully", req.UserId),
	}

	return &output, nil
}

func (grpca Adapter) UpdateUser(ctx context.Context, req *pb.UpdateUserInput) (*pb.UpdateUserOutput, error) {

	err := grpca.api.ProcessUpdateUser(req.UserId, req.UserName)

	if err != nil {
		return nil, err
	}
	output := pb.UpdateUserOutput{
		Message: fmt.Sprintf("User %v Updated Successfully", req.UserName),
	}

	return &output, nil
}
