package repository

import (
	"context"

	"github.com/ppb03/question-answer-api/internal/models"
)

type QuestionRepository interface {
	Create(ctx context.Context, question *models.Question) error
	Delete(ctx context.Context, id uint) error
	
	GetByID(ctx context.Context, id uint) (*models.Question, error)
	GetAll(ctx context.Context) ([]models.Question, error)
	GetWithAnswers(ctx context.Context, id uint) (*models.Question, error)
}

type AnswerRepository interface {
	Create(ctx context.Context, answer *models.Answer) error
	Delete(ctx context.Context, id uint) error

	GetByID(ctx context.Context, id uint) (*models.Answer, error)
}