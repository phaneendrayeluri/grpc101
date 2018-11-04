package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"google.golang.org/grpc"
	"stash.lbidts.com/~syeluri/gotooling/grpc/pb"
)

func main() {

	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "missing subcommand: list or add")
		os.Exit(1)
	}

	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not connect to backend: %v\n", err)
		os.Exit(1)
	}

	client := pb.NewMathClient(conn)

	switch flag.Arg(0) {
	case "+":
		log.Printf("Operation Successful - Result : %v", OutboundAdd(context.Background(), client))
	case "-":
		log.Printf("Operation Successful - Result : %v", OutboundSub(context.Background(), client))
	default:
		log.Println("Undefined Operation, Type Help")
	}

}

// OutboundAdd ...
func OutboundAdd(ctx context.Context, cli pb.MathClient) interface{} {
	x, _ := strconv.Atoi(flag.Arg(1))
	y, _ := strconv.Atoi(flag.Arg(2))
	n, err := cli.Add(ctx, &pb.Tuple{X: &pb.Number{Value: int32(x)}, Y: &pb.Number{Value: int32(y)}})
	if err != nil {
		log.Fatalf("Failed Add %v", err)
		return ""
	}
	return n.Value
}

// OutboundSub ...
func OutboundSub(ctx context.Context, cli pb.MathClient) interface{} {
	x, _ := strconv.Atoi(flag.Arg(1))
	y, _ := strconv.Atoi(flag.Arg(2))
	n, err := cli.Sub(ctx, &pb.Tuple{X: &pb.Number{Value: int32(x)}, Y: &pb.Number{Value: int32(y)}})
	if err != nil {
		log.Fatalf("Failed Add %v", err)
		return ""
	}
	return n.Value
}
