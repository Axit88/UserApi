package rpc

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/Axit88/UserApi/src/config"
	ports "github.com/Axit88/UserApi/src/domain/userService/core/ports/incoming"
	pb "github.com/Axit88/UserApi/src/domain/userService/infrastructure/grpcClient/pb"
	"github.com/MindTickle/mt-go-logger/logger"

	"google.golang.org/grpc"
)

type Adapter struct {
	logger *logger.LoggerImpl
	api    ports.APIPort
	pb.UnimplementedUserApiServer
}

func NewGrpcClient(api ports.APIPort, l *logger.LoggerImpl) *Adapter {
	return &Adapter{
		api:    api,
		logger: l,
	}
}

func (grpca Adapter) GetUser(ctx context.Context, req *pb.GetUserInput) (*pb.GetUserOutput, error) {
	res, err := grpca.api.ProcessGetUser(req.UserId)
	if err != nil {
		grpca.logger.Errorf(ctx, "Application Is Not Able To Process Request", err)
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
		grpca.logger.Errorf(ctx, "Application Is Not Able To Process Request", err)
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
		grpca.logger.Errorf(ctx, "Application Is Not Able To Process Request", err)
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
		grpca.logger.Errorf(ctx, "Application Is Not Able To Process Request", err)
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
	pb.RegisterUserApiServer(grpcServer, userServiceServer)

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
	pb.RegisterUserApiServer(grpcServer, server)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve gRPC server over port 8080: %v", err)
	}
}
