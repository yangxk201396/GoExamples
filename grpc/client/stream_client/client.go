package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"time"

	pb "github.com/yang201396/GoExamples/grpc/proto"
)

const (
	PORT = "9002"
)

func main() {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	conn, err := grpc.Dial(":"+PORT, opts...)
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
	defer conn.Close()

	client := pb.NewStreamServiceClient(conn)

	err = printLists(client, &pb.StreamRequest{Pt: &pb.StreamPoint{Name: "List", Value: 2018}})
	if err != nil {
		log.Fatalf("printLists.err: %v", err)
	}

	err = printRecord(client, &pb.StreamRequest{Pt: &pb.StreamPoint{Name: "Record", Value: 2018}})
	if err != nil {
		log.Fatalf("printRecord.err: %v", err)
	}

	err = printRoute(client, &pb.StreamRequest{Pt: &pb.StreamPoint{Name: "Route", Value: 2018}})
	if err != nil {
		log.Fatalf("printRoute.err: %v", err)
	}
}

func printLists(client pb.StreamServiceClient, r *pb.StreamRequest) error {
	stream, err := client.List(context.Background(), r)
	if err != nil {
		return err
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			log.Printf("List Recv err: %v, 通信结束", io.EOF)
			return nil
		}
		if err != nil {
			return err
		}

		log.Printf("List Recv(pj.name: %s, pt.value: %d)", resp.Pt.Name, resp.Pt.Value)
	}
}

func printRecord(client pb.StreamServiceClient, r *pb.StreamRequest) error {
	stream, err := client.Record(context.Background())
	if err != nil {
		return err
	}

	for n := 0; n < 6; n++ {
		r.Pt.Value += int32(n)
		err := stream.Send(r)
		if err != nil {
			return err
		}
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		return err
	}

	log.Printf("Record CloseAndRecv(pj.name: %s, pt.value: %d)", resp.Pt.Name, resp.Pt.Value)

	return nil
}

func printRoute(client pb.StreamServiceClient, r *pb.StreamRequest) error {
	stream, err := client.Route(context.Background())
	if err != nil {
		return err
	}

	for n := 0; n <= 6; n++ {
		r.Pt.Value += int32(n)
		err = stream.Send(r)
		if err != nil {
			log.Printf("Route Send err: %v", err.Error())
			return err
		}
	}

	for n := 0; n < 10; n++ {
		resp, err := stream.Recv()
		if err == io.EOF {
			log.Printf("Route Recv err: %v, 通信结束", io.EOF)
			break
		}
		if err != nil {
			log.Printf("Route Recv err: %v", err.Error())
			return err
		}
		log.Printf("Route: resp(pj.name: %s, pt.value: %d)", resp.Pt.Name, resp.Pt.Value)
	}

	time.Sleep(3 * time.Second)
	err = stream.CloseSend()
	if err != nil {
		log.Printf("Route CloseSend err: %v", err.Error())
		return err
	}

	return nil
}
