package main

import (
    "net"
    "log"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"

    jwt "grpc-example/proto/jwt"
)
import _ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"

const (
	port = ":50051"
)

func main() {
    lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
    }
    log.Printf("Listening on port %s", port)
	s := grpc.NewServer()
	jwt.RegisterJWTServiceServer(s, newJwtServer())
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}