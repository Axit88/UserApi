package rpc

import (
	"log"
	"net"

	"github.com/Axit88/UserApi/src/domain/userService/adapters/incoming/grpc/pb"
	ports "github.com/Axit88/UserApi/src/domain/userService/core/ports/incoming"

	"google.golang.org/grpc"
)

type Adapter struct {
	api ports.APIPort
	pb.UnimplementedTestApiServer
}

func NewAdapter(api ports.APIPort) *Adapter {
	return &Adapter{api: api}
}

func (grpca Adapter) Run() {
	var err error

	listen, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("failed to listen on port 9000: %v", err)
	}

	arithmeticServiceServer := grpca
	grpcServer := grpc.NewServer()
	pb.RegisterTestApiServer(grpcServer, arithmeticServiceServer)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve gRPC server over port 9000: %v", err)
	}
}
