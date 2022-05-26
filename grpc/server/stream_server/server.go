package main

import (
	"google.golang.org/grpc"
	"io"
	"log"
	"net"

	pb "github.com/yang201396/GoExamples/grpc/proto"
)

type StreamService struct{}

const (
	PORT = "9002"
)

func main() {
	server := grpc.NewServer()
	pb.RegisterStreamServiceServer(server, &StreamService{})

	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	err = server.Serve(lis)
	if err != nil {
		return
	}
}

func (s *StreamService) List(r *pb.StreamRequest, stream pb.StreamService_ListServer) error {
	log.Printf("List: req(pt.name: %s, pt.value: %d)", r.Pt.Name, r.Pt.Value)

	for n := 0; n <= 6; n++ {
		err := stream.Send(&pb.StreamResponse{
			Pt: &pb.StreamPoint{
				Name:  r.Pt.Name,
				Value: r.Pt.Value + int32(n),
			},
		})
		if err != nil {
			log.Println("Route Send err:", err.Error())
			return err
		}
	}

	return nil
}

func (s *StreamService) Record(stream pb.StreamService_RecordServer) error {
	for {
		r, err := stream.Recv()
		if err == io.EOF {
			log.Printf("Record Recv err: %v, 通信完毕", io.EOF)
			return stream.SendAndClose(&pb.StreamResponse{Pt: &pb.StreamPoint{Name: "Record", Value: 111}})
		}
		if err != nil {
			log.Println("Record Recv err:", err.Error())
			return err
		}
		log.Printf("Record Recv(pt.name: %s, pt.value: %d)", r.Pt.Name, r.Pt.Value)
	}
}

func (s *StreamService) Route(stream pb.StreamService_RouteServer) error {
	for n := 0; n < 10; n++ {
		err := stream.Send(&pb.StreamResponse{
			Pt: &pb.StreamPoint{
				Name:  "Server Route",
				Value: int32(n),
			},
		})
		if err != nil {
			log.Println("Route Send err:", err.Error())
			return err
		}
	}

	for {
		r, err := stream.Recv()
		if err == io.EOF {
			log.Printf("Route Recv err: %v, 通信完毕", io.EOF)
			return nil
		}
		if err != nil {
			log.Println("Route Recv err:", err.Error())
			return err
		}
		log.Printf("Route Recv(pt.name: %s, pt.value: %d)", r.Pt.Name, r.Pt.Value)
	}
}
