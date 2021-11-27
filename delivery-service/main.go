package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/detohm/grpc-study/proto"
	"google.golang.org/grpc"
)

type DeliveryServer struct {
	pb.UnimplementedDeliverySVCServer
}

func NewDeliveryServer() *DeliveryServer {
	return &DeliveryServer{}
}

func (d *DeliveryServer) Create(ctx context.Context, p *pb.Product) (*pb.Delivery, error) {
	log.Printf("create called with product : %v", p)
	return &pb.Delivery{
		DeliveryId: 1,
		TrackingId: "d3u2dfn2lh2e2dss",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:5000")
	if err != nil {
		fmt.Println("")
	}
	grpcServer := grpc.NewServer()
	pb.RegisterDeliverySVCServer(grpcServer, NewDeliveryServer())
	grpcServer.Serve(lis)
}
