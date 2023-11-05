package server

import (
	"fmt"
	"net"

	"github.com/evgsrkn/go-ddd-example/user/internal/ports"
	"github.com/evgsrkn/go-ddd-example/user/pkg/pb"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func ServeGrpc(log *zap.Logger, srv *grpc.Server, api *ports.GrpcUserServer) {
	pb.RegisterUserServiceServer(srv, api)
	reflection.Register(srv)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", 8081))
	if err != nil {
		log.Error("cannot bind server", zap.Error(err))
		return
	}

	if err := srv.Serve(lis); err != nil {
		log.Error("cannot listen server", zap.Error(err))
		return
	}
}
