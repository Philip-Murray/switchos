package main

import (
	"context"
	"errors"
	"log"
	"net"
	"os"
	"proto-playground/proto"
	"time"

	"cloud.google.com/go/pubsub"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var topic *pubsub.Topic

func main() {
	os.Setenv("PUBSUB_EMULATOR_HOST", "localhost:8085")
	ps_ctx := context.Background()
	notified := false
	var client *pubsub.Client
	var err error
	for {
		client, err = pubsub.NewClient(ps_ctx, "karhoo-local")
		if err == nil {
			break
		} else if !notified {
			log.Fatalf("Failed to create pubsub client %v", err)
			log.Fatalf("Will reattempt every second to connect to PubSub, please start PubSub terminal")
			notified = true
		}
		time.Sleep(time.Second)
	}
	topic, err = client.CreateTopic(ps_ctx, "events.TripBooked")
	if err != nil {
		log.Fatalf("Failed to create topic %v", err)
		panic()
	}
	defer topic.Stop()

	lis, err := net.Listen("tcp", ":3001") //React port + 1
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

type server struct{}

func (s *server) MakeReservation(ctx context.Context, req *proto.BookTrip) (*proto.Trip, error) {
	log.Printf("Recieved request to book trip from client %s", req.PassengerName)
	ps_ctx := context.Background()

	if _, err := topic.Publish(ps_ctx, &pubsub.Message{Data: []byte("payload")}).Get(ps_ctx); err != nil {
		log.Fatalf(err)
		return &proto.Trip{}, errors.New("Could not confirm reservation %v", err)
	}
	new_trip := proto.Trip{
		PassengerName: req.PassengerName,
		DriverName:    "Marek",
	}
	return &new_trip, nil
}
