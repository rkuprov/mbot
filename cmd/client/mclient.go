package main

import (
	"context"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	server "mbot/pkg/proto/mserver"
)

func main() {
	var conn *grpc.ClientConn

	conn, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := server.NewMBotServerClient(conn)
	resp, err := c.CreateCustomer(context.Background(), &server.CreateCustomerRequest{
		Id:      "1",
		Name:    "John Doe",
		Email:   "doe@gmail.com",
		Contact: "1234567890",
		Token:   uuid.New().String(),
	})
	if err != nil {
		panic(err)
	}
	println(resp.Message)
}
