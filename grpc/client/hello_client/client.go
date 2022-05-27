package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/yang201396/GoExamples/grpc/proto"
)

func main() {
	creds, err := credentials.NewClientTLSFromFile("../../conf/server.pem", "panther")
	if err != nil {
		log.Printf("Failed to create TLS credentials %v", err)
	}
	conn, err := grpc.Dial(":50052", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
	defer conn.Close()

	if err != nil {
		log.Println(err)
	}

	c := pb.NewHelloWorldClient(conn)
	body := &pb.HelloWorldRequest{
		Referer: "Grpc",
	}

	r, err := c.SayHelloWorld(context.Background(), body)
	if err != nil {
		log.Println("Server SayHelloWorld err: ", err)
	}

	log.Println("Hello SayHelloWorld resp:", r.Message)
}
