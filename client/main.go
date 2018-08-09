package main

import (
    "log"
    "time"
    "context"
    
    "google.golang.org/grpc"
    "google.golang.org/grpc/metadata"

    jwt "grpc-example/proto/jwt_service"
)

const (
	address     = "localhost:50051"
)

func main() {
    // Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
    c := jwt.NewJWTServiceClient(conn)
    
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()
    // ends setting up

    // set up header
    header := metadata.New(map[string]string{})
    // r, err := c.Verify(ctx, &jwt.Request{Token: "test123"}, grpc.Header(&header))
    r, err := c.Echo(ctx, &jwt.Request{}, grpc.Header(&header))

    if err != nil {
		log.Fatalf("could not greet: %v", err)
    }
    log.Println("header foo:", header["foo"])
    log.Println("header bar:", header["bar"])
    log.Println("content-type", header["content-type"])
	log.Printf("verify status: %s", r.Status)
}