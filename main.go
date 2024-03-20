package main

import (
	"context"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/zackzhangkai/grpc-pubsub-example/api/proto"
	"github.com/zackzhangkai/grpc-pubsub-example/pkg/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/keepalive"
	"k8s.io/klog/v2"
	"log"
	"net"
	"time"
)

func main() {
	lis, _ := net.Listen("tcp", "127.0.0.1:1234")
	opts := []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandlerContext(func(ctx context.Context, rec interface{}) (err error) {
			klog.Warningf("Recovered in f %v", rec)
			return grpc.Errorf(codes.Internal, "Recovered from panic")
		}),
	}
	grpcServer := grpc.NewServer(
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			MinTime:             5 * time.Second,
			PermitWithoutStream: true, // Allow pings even when there are no active streams
		}),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			Time:    2 * time.Hour,
			Timeout: 20 * time.Second,
		}),
		grpc_middleware.WithUnaryServerChain(
			grpc_recovery.UnaryServerInterceptor(opts...),
		),
		grpc_middleware.WithStreamServerChain(
			grpc_recovery.StreamServerInterceptor(opts...),
		),
	)
	klog.Info("start server, listen: 1234")
	proto.RegisterPubSubServiceServer(grpcServer, service.NewService())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	klog.Warningf("stop server")
	grpcServer.GracefulStop()
}
