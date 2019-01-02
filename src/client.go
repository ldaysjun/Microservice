

package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"

	pb "TimeService"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewTimServiceClient(conn)

	// Contact the server and print out its response.

	r, err := c.TimedTask(context.Background(), &pb.TimedTaskRequest{Start:1})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r)
}