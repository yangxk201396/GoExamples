package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	"github.com/yang201396/GoExamples/grpc/pkg/gtls"
	pb "github.com/yang201396/GoExamples/grpc/proto"
)

const PORT = "9004"

type Auth struct {
	AppKey    string
	AppSecret string
}

func (a *Auth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{"app_key": a.AppKey, "app_secret": a.AppSecret}, nil
}

func (a *Auth) RequireTransportSecurity() bool {
	return true
}

func main() {
	tlsClient := gtls.Client{
		ServerName: "www.abc.com",
		CertFile:   "/Users/yangxk/GOPATH/src/github.com/yang201396/GoExamples/grpc/conf/server/server.pem",
	}
	c, err := tlsClient.GetTLSCredentials()
	if err != nil {
		log.Fatalf("tlsClient.GetTLSCredentials err: %v", err)
	}

	auth := Auth{
		AppKey:    "eddycjy",
		AppSecret: "20181005",
	}
	conn, err := grpc.Dial(":"+PORT, grpc.WithTransportCredentials(c), grpc.WithPerRPCCredentials(&auth))
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
	defer conn.Close()

	client := pb.NewSearchServiceClient(conn)
	resp, err := client.Search(context.Background(), &pb.SearchRequest{
		Request: "gRPC",
	})
	if err != nil {
		log.Fatalf("client.Search err: %v", err)
	}

	log.Printf("resp: %s", resp.GetResponse())
}
