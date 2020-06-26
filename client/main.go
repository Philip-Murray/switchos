package main

import (
	"grpcTestOne/proto"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var con *grpc.ClientConn
	con, err := grpc.Dial("localhost:3001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	defer con.Close()
	client := proto.NewReservationServiceClient(con)

	book_trip_request := proto.BookTrip{
		PassengerName: "Philip",
	}

	confirmed_trip, err := client.MakeReservation(context.Background(), &book_trip_request)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	log.Printf("Response from server: %s", confirmed_trip.DriverName)

}
