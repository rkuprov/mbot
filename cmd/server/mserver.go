package main

import (
	"context"
	"net"

	"github.com/go-chi/chi/v5"
	"google.golang.org/grpc"
)

func main() {
	r := chi.NewRouter()
	SetupRoutes(r)
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	grpc.ServiceRegistrar(service, s)
	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}
}

type mbotserver struct {
	server.UnimplementedMBotServerServer
}

func (m mbotserver) CreateCustomer(c context.Context, req *server.CreateCustomerRequest) (*server.CreateCustomerReply, error) {
	return &server.CreateCustomerReply{
		Message: "success",
		Id:      req.Id,
		Token:   req.Token,
	}, nil
}
