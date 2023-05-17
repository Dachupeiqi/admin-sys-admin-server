package main

import (
	"flag"
	"fmt"
	"ginAdminServer/remoterpc/talentrpc/internal/config"
	"ginAdminServer/remoterpc/talentrpc/internal/server"
	"ginAdminServer/remoterpc/talentrpc/internal/svc"
	__ "ginAdminServer/remoterpc/talentrpc/pb"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/talentrpc.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		__.RegisterTalentServer(grpcServer, server.NewTalentServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	//_ = consul.RegisterService(c.ListenOn, c.Consul)

	defer s.Stop()

	fmt.Printf("Starting remoterpc server at %s...\n", c.ListenOn)
	s.Start()
}
