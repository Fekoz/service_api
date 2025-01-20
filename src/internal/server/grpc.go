package server

import (
	"service_api/internal/controller"

	conf "bitbucket.org/klaraeng/package_config/conf"
	apiAdmin "bitbucket.org/klaraeng/package_proto/service_api/admin"
	api "bitbucket.org/klaraeng/package_proto/service_api/client"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(
	c *conf.Server,
	client *controller.ClientService,
	admin *controller.AdminService,
	logger log.Logger,
) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			metadata.Server(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	api.RegisterClientServer(srv, client)
	apiAdmin.RegisterAdminServer(srv, admin)
	return srv
}
