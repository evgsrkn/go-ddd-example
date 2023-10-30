package rpc

import (
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Option = func() *grpc.ServerOption

func New(opts ...Option) *grpc.Server {
	var options []grpc.ServerOption
	for _, opt := range opts {
		options = append(options, *opt())
	}

	return grpc.NewServer(options...)
}

func WithLoggerZap(logger *zap.Logger) Option {
	return func() *grpc.ServerOption {
		grpc_zap.ReplaceGrpcLoggerV2(logger)
		opts := []grpc_zap.Option{
			grpc_zap.WithLevels(grpc_zap.DefaultCodeToLevel),
		}

		opt := grpc.ChainUnaryInterceptor(
			grpc_ctxtags.UnaryServerInterceptor(
				grpc_ctxtags.WithFieldExtractor(
					grpc_ctxtags.CodeGenRequestFieldExtractor,
				),
			),
			grpc_zap.UnaryServerInterceptor(logger, opts...),
		)

		return &opt
	}
}
