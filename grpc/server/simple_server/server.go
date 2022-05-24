package main

import (
	"context"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"runtime/debug"

	"github.com/yang201396/GoExamples/grpc/pkg/gtls"
	pb "github.com/yang201396/GoExamples/grpc/proto"
)

type SearchService struct{}

func (s *SearchService) Search(ctx context.Context, r *pb.SearchRequest) (*pb.SearchResponse, error) {
	if ctx.Err() == context.Canceled {
		return nil, status.Errorf(codes.Canceled, "searchService.Search canceled")
	}
	return &pb.SearchResponse{Response: r.GetRequest() + " Server"}, nil
}

const PORT = "9001"

func main() {
	tlsServer := gtls.Server{
		CaFile:   "/Users/yangxk/GOPATH/src/github.com/yang201396/GoExamples/grpc/conf/ca.pem",
		CertFile: "/Users/yangxk/GOPATH/src/github.com/yang201396/GoExamples/grpc/conf/server/server.pem",
		KeyFile:  "/Users/yangxk/GOPATH/src/github.com/yang201396/GoExamples/grpc/conf/server/server.key",
	}
	c, err := tlsServer.GetCredentialsByCA()
	if err != nil {
		log.Fatalf("GetTLSCredentialsByCA err: %v", err)
	}

	opts := []grpc.ServerOption{
		grpc.Creds(c),
		grpc_middleware.WithUnaryServerChain(
			RecoveryInterceptor,
			LoggingInterceptor,
		),
	}

	server := grpc.NewServer(opts...)
	pb.RegisterSearchServiceServer(server, &SearchService{})

	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	log.Println("tcp:", PORT)
	err = server.Serve(lis)
	if err != nil {
		return
	}
}

func LoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf("gRPC method: %s, %v", info.FullMethod, req)
	resp, err := handler(ctx, req)
	log.Printf("gRPC method: %s, %v", info.FullMethod, resp)
	return resp, err
}

func RecoveryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	defer func() {
		if e := recover(); e != nil {
			debug.PrintStack()
			err = status.Errorf(codes.Internal, "Panic err: %v", e)
		}
	}()

	return handler(ctx, req)
}
