package main

import (
	"context"
	"log"
	"net"
	"proto-playground/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"cloud.google.com/go/pubsub"
)

func main() {
	os.Setenv("PUBSUB_EMULATOR_HOST", "localhost:8085")
	ps_ctx := context.Background()
	client, err := pubsub.NewClient(ctx, proj)
	var client *pubsub.Client; var err_mem error
	for {
		client, err := pubsub.NewClient(ps_ctx, "karhoo-local")
		if err == nil {
			break
		} else if err != err_mem {
			log.Fatalf("Failed to create pubsub client %v", err)
			log.Fatalf("Will reattempt every 0.5 seconds")
			err_mem = err
		}
		time.Sleep(time.Second * 0.5)
	}
	var topic *Topic; err_mem = nil
	for {
		topic, err 
	}


	lis, err := net.Listen("tcp", ":3001") //React port + 1

	if err != nil {
		log.Fatalf("Error %v", err)
	}
	client, 
	gcrpServer := grpc.NewServer()
	//proto generates ReservationService Server
	proto.RegisterReservationServiceServer(gcrpServer, &server{})
	reflection.Register(gcrpServer)

	if err := gcrpServer.Serve(lis); err != nil {
		log.Fatalf("Error %v", err)
	}

}

type server struct {}

func (s *server) MakeReservation(ctx context.Context, req *proto.BookTrip) (*proto.Trip, error) {
	log.Printf("Recieved request to book trip from client %s", req.PassengerName)
	new_trip := proto.Trip {
		PassengerName: req.PassengerName,
		DriverName:    "Marek",
	}
	return &new_trip, nil
}
