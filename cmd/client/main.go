package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"

	"github.com/pddg/go-protobuf-v2-sample/hello/pb"
)

var server string

func main() {
	flag.StringVar(&server, "server", "localhost:8080", "Server address")
	flag.Parse()
	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}
	args := flag.Args()

	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, server, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewHelloServiceClient(conn)
	response, err := client.Hello(ctx, &pb.HelloRequest{Name: args[0]})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response.Message)
}