package main

import (
	"context"
	"log"
	"net"

	grpc "google.golang.org/grpc"
	"stash.lbidts.com/~syeluri/gotooling/grpc/pb"
)

func main() {
	srv := grpc.NewServer()

	var serverImpl MathServerImpl
	pb.RegisterMathServer(srv, serverImpl)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("could not listen to :8080: %v", err)
	}

	log.Fatal(srv.Serve(l))
}

// MathServerImpl ...
type MathServerImpl struct{}

// Add ...
func (msi MathServerImpl) Add(ctx context.Context, in *pb.Tuple) (*pb.Number, error) {
	x := in.GetX().Value
	y := in.GetY().Value
	return &pb.Number{Value: x + y}, nil
}

// Sub ...
func (msi MathServerImpl) Sub(ctx context.Context, in *pb.Tuple) (*pb.Number, error) {
	x := in.GetX().Value
	y := in.GetY().Value
	return &pb.Number{Value: x - y}, nil
}
