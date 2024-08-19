package main

import (
	"apiproducts/src/pb/products"
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	products.ProductServiceServer
}

func (s *server) Create(ctx context.Context, product *products.Product) (*products.Product, error) {
	return &products.Product{}, nil
}

func (s *server) FindAll(ctx context.Context, product *products.Product) (*products.ProductList, error) {
	return &products.ProductList{}, nil
}

func main() {
	fmt.Println("Starting grpc server")
	srv := server{}

	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalln("error on vreate listener. Error ", err)
	}

	s := grpc.NewServer()
	products.RegisterProductServiceServer(s, &srv)

	if err := s.Serve(listener); err != nil {
		log.Fatalln("error on server. Error: ", err)
	}

}
