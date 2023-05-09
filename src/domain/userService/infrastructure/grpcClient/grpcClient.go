package rpc

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/Axit88/UserApi/src/config"
	ports "github.com/Axit88/UserApi/src/domain/userService/core/ports/incoming"
	"github.com/Axit88/UserApi/src/domain/userService/infrastructure/grpcClient/pb"

	"google.golang.org/grpc"
)

type Adapter struct {
	api ports.APIPort
	pb.UnimplementedTestApiServer
}

func NewAdapter(api ports.APIPort) *Adapter {
	return &Adapter{api: api}
}

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

	err := grpca.api.ProcessAddUser(req.UserId, req.UserName)

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
		Message: fmt.Sprintf("User %v Updated Successfully", req.UserId),
	}

	return &output, nil
}

func (grpca Adapter) Run() {
	var err error
	var cfn, _ = config.NewConfig()
	url := cfn.UserServiceUrl.GrpcUrl

	listen, err := net.Listen("tcp", url)
	if err != nil {
		log.Fatalf("failed to listen on port 8080: %v", err)
	}

	userServiceServer := grpca
	grpcServer := grpc.NewServer()
	pb.RegisterTestApiServer(grpcServer, userServiceServer)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve gRPC server over port 8080: %v", err)
	}
}

func StartServer(server *Adapter) {
	var err error
	var cfn, _ = config.NewConfig()
	url := cfn.UserServiceUrl.GrpcUrl

	listen, err := net.Listen("tcp", url)
	if err != nil {
		log.Fatalf("failed to listen on port 8080: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterTestApiServer(grpcServer, server)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve gRPC server over port 8080: %v", err)
	}
}
