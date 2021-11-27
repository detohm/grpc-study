package main

import (
	"context"
	"log"
	"time"

	pb "github.com/detohm/grpc-study/proto"
	"google.golang.org/grpc"
)

func runCreate(client pb.DeliverySVCClient, product *pb.Product) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	delivery, err := client.Create(ctx, product)
	if err != nil {
		log.Printf("error occurred: %v", err)
	}
	log.Printf("got delivery: %v", delivery)

}

func main() {

	opts := []grpc.DialOption{}
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial("localhost:5000", opts...)
	if err != nil {
		log.Printf("error occurred: %v", err)
	}
	defer conn.Close()

	product := pb.Product{
		ProductId: 1,
		UserId:    1,
	}

	client := pb.NewDeliverySVCClient(conn)

	runCreate(client, &product)
}
