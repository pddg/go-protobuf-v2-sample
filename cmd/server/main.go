package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/pddg/go-protobuf-v2-sample/hello"
	"github.com/pddg/go-protobuf-v2-sample/hello/pb"
)

var (
	port int
	host string
)

func main() {
	flag.IntVar(&port, "port", 8080, "Port number")
	flag.StringVar(&host, "host", "0.0.0.0", "Listen host name")
	flag.Parse()

	listenAddr, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		log.Fatalf("Failed to listen '%s:%d'\n", host, port)
	}
	log.Printf("HelloServiceServer is listening on tcp://%s:%d\n", host, port)

	s := grpc.NewServer()
	helloServer := hello.NewHelloServiceServer()
	pb.RegisterHelloServiceServer(s, helloServer)

	// Enable reflection
	reflection.Register(s)

	if err := s.Serve(listenAddr); err != nil {
		log.Fatal(err)
	}
}
