package service

import (
	"context"

	"github.com/ppb03/question-answer-api/internal/models"
)

type QuestionService interface {
	Create(ctx context.Context, question *models.Question) error
	Delete(ctx context.Context, id uint) error

	GetAll(ctx context.Context) ([]models.Question, error)
	GetByID(ctx context.Context, id uint) (*models.Question, error)
	GetWithAnswers(ctx context.Context, id uint) (*models.Question, error)
}

type AnswerService interface {
	Create(ctx context.Context, answer *models.Answer) error
	Delete(ctx context.Context, id uint) error

	GetByID(ctx context.Context, id uint) (*models.Answer, error)
}
