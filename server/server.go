package server

import (
	"context"

	"github.com/platzi/go-platzi-protobuf/models"
	"github.com/platzi/go-platzi-protobuf/repository"
	"github.com/platzi/go-platzi-protobuf/studentpb"
)

type Server struct {
	repo repository.Repository
	studentpb.UnimplementedStudentServiceServer
}

func NewStudentServer(repo repository.Repository) *Server {
	return &Server{repo: repo}
}

func (s *Server) GetStudent(ctx context.Context, request *studentpb.GetStudentRequest) (*studentpb.Student, error) {
	student, err := s.repo.GetStudent(ctx, request.GetId())
	if err != nil {
		return nil, err
	}

	return &studentpb.Student{
		Id:   student.Id,
		Name: student.Name,
		Age:  student.Age,
	}, nil
}

func (s *Server) SetStudent(ctx context.Context, request *studentpb.Student) (*studentpb.SetStudentResposne, error) {
	student := models.Student{
		Id:   request.GetId(),
		Name: request.GetName(),
		Age:  request.GetAge(),
	}

	err := s.repo.SetStudent(ctx, &student)
	if err != nil {
		return nil, err
	}

	return &studentpb.SetStudentResposne{
		Id: student.Id,
	}, nil
}
