package main

import (
	"context"
	"grpcTestOne/proto"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func main() {
	lis, err := net.Listen("tcp", ":3001") //React + 1
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	gcrpServer := grpc.NewServer()
	//proto generates ReservationService Server
	proto.RegisterReservationServiceServer(gcrpServer, &server{})
	reflection.Register(gcrpServer)

	if err := gcrpServer.Serve(lis); err != nil {
		log.Fatalf("Error %v", err)
	}

}

func (s *server) MakeReservation(ctx context.Context, req *proto.BookTrip) (*proto.Trip, error) {
	log.Printf("Recieved message from client %s", req.PassengerName)
	return &proto.Trip{PassengerName: "Default pass", DriverName: "April"}, nil
}
