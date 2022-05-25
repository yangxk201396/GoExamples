package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	"github.com/yang201396/GoExamples/grpc/pkg/gtls"
	pb "github.com/yang201396/GoExamples/grpc/proto"
)

const PORT = "9003"

func main() {
	tlsClient := gtls.Client{
		ServerName: "www.abc.com",
		CertFile:   "/Users/yangxk/GOPATH/src/github.com/yang201396/GoExamples/grpc/conf/server/server.pem",
	}
	c, err := tlsClient.GetTLSCredentials()
	if err != nil {
		log.Fatalf("tlsClient.GetTLSCredentials err: %v", err)
	}

	conn, err := grpc.Dial(":"+PORT, grpc.WithTransportCredentials(c))
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
