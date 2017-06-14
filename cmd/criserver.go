package main

import (
        "flag"
        "log"
	"net"

        "github.com/brendandburns/cri-base/generated"
        "github.com/brendandburns/cri-base/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	addr = flag.String("addr", ":50051", "The address to listen on")
)

func main() {
        flag.Parse()
	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	generated.RegisterRuntimeServiceServer(s, &server.RuntimeServer{})
	generated.RegisterImageServiceServer(s, &server.ImageServer{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
