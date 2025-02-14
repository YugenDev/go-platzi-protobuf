package repository

import (
	"context"

	"github.com/platzi/go-platzi-protobuf/models"
)

type Repository interface {
	GetStudent(ctx context.Context, id string) (*models.Student, error)
	SetStudent(ctx context.Context, student *models.Student) error
	GetTest(ctx context.Context, id string) (*models.Test, error)
	SetTest(ctx context.Context, test *models.Test) error
	SetQuestion(ctx context.Context, question *models.Question) error
	//GetQuestions(ctx context.Context, testId string) ([]*models.Question, error)
}

var implementation Repository

func SetRepository(repository Repository) {
	implementation = repository
}

func GetStudent(ctx context.Context, id string) (*models.Student, error) {
	return implementation.GetStudent(ctx, id)
}

func SetStudent(ctx context.Context, student *models.Student) error {
	return implementation.SetStudent(ctx, student)
}

func GetTest(ctx context.Context, id string) (*models.Test, error) {
	return implementation.GetTest(ctx, id)
}

func SetTest(ctx context.Context, test *models.Test) error {
	return implementation.SetTest(ctx, test)
}

func SetQuestions(ctx context.Context, question *models.Question) error {
	return implementation.SetQuestion(ctx, question)
}

//func GetQuestions(ctx context.Context, testId string) ([]*models.Question, error) {
//	return implementation.GetQuestions(ctx, testId)
//}
