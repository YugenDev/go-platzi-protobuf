package server

import (
	"github.com/platzi/go-platzi-protobuf/repository"
)

type Server struct {
	repo repository.Repository
}

func NewStudentServer(repo repository.Repository) *Server {
	return &Server{repo: repo}
}
