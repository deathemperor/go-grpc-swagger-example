package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	jwt "grpc-example/proto/jwt_service"
)

type server struct{}

// setting up "class" JWTServiceServer
func newServer() jwt.JWTServiceServer {
	return new(server)
}

// implementation for Verify method
func (s *server) Verify(ctx context.Context, in *jwt.Request) (*jwt.Response, error) {
	log.Println("received msg: ", in)
	log.Println("received token: ", in.Token)
	var out = new(jwt.Response)
	grpc.SendHeader(ctx, metadata.New(map[string]string{
		"foo": "foo1",
		"bar": "bar1",
	}))
	out.Status = "success"

	return out, nil
}

func (s *server) Echo(ctx context.Context, in *jwt.Request) (*jwt.Response, error) {
	log.Println("message to echo: ", in.Token)
	var out = new(jwt.Response)
	grpc.SendHeader(ctx, metadata.New(map[string]string{
		"foo": "foo1",
		"bar": "bar1",
	}))
	out.Status = "success"
	out.Token = in.Token

	return out, nil
}
