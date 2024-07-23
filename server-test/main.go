package main

import (
	"log"
	"net"

	"github.com/platzi/go-platzi-protobuf/database"
	"github.com/platzi/go-platzi-protobuf/server"
	"github.com/platzi/go-platzi-protobuf/testpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listener, err := net.Listen("tcp", ":5070")
	if err != nil {
		log.Fatal(err)
	}

	repo, err := database.NewPostgresRepository("postgres://postgres:postgres@localhost:54321/postgres?sslmode=disable")
	server := server.NewTestServer(repo)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	testpb.RegisterTestServiceServer(s, server)

	reflection.Register(s)

	if err := s.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
