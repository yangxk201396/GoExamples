package main

import (
	"context"
	"log"
	"time"

	"github.com/yang201396/GoExamples/grpc/pkg/gtls"
	pb "github.com/yang201396/GoExamples/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const PORT = "9001"

func main() {
	tlsClient := gtls.Client{
		ServerName: "www.abc.com", // 和openssl/server/server.conf 的 [alt_names] 对应
		CaFile:     "/Users/yangxk/GOPATH/src/github.com/yang201396/GoExamples/grpc/conf/ca.pem",
		CertFile:   "/Users/yangxk/GOPATH/src/github.com/yang201396/GoExamples/grpc/conf/client/client.pem",
		KeyFile:    "/Users/yangxk/GOPATH/src/github.com/yang201396/GoExamples/grpc/conf/client/client.key",
	}

	c, err := tlsClient.GetCredentialsByCA()
	if err != nil {
		log.Fatalf("GetTLSCredentialsByCA err: %v", err)
	}

	conn, err := grpc.Dial(":"+PORT, grpc.WithTransportCredentials(c))
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
	defer conn.Close()

	log.Println("tcp:", PORT)

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	defer cancel()

	client := pb.NewSearchServiceClient(conn)
	resp, err := client.Search(ctx, &pb.SearchRequest{
		Request: "gRPC",
	})
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				log.Fatalln("client.Search err: deadline")
			}
		}
		log.Fatalf("client.Search err: %v", err)
	}

	log.Printf("resp: %s", resp.GetResponse())
}
