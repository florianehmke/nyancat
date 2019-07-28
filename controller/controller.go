package controller

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc"

	pb "github.com/florianehmke/nyancat/api"
)

const (
	address = "localhost:50051"
)

func Start() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMinerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.MineCat(ctx, &pb.MineRequest{Id: uuid.New().String()})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("%s: %s", r.Id, r.Cat)
}
