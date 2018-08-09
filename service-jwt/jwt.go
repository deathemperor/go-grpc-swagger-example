package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	jwt "grpc-example/proto/jwt"
)

type jwtServer struct{}

// setting up "class" JWTServiceServer
func newJwtServer() jwt.JWTServiceServer {
	return new(jwtServer)
}

// implementation for Verify method
func (s *jwtServer) Verify(ctx context.Context, in *jwt.Request) (*jwt.Response, error) {
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

func (s *jwtServer) Echo(ctx context.Context, in *jwt.Request) (*jwt.Response, error) {
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
