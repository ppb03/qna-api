package service

import (
	"context"
	"errors"

	"github.com/ppb03/question-answer-api/internal/models"
	"github.com/ppb03/question-answer-api/internal/repository"
)

type questionService struct {
	repository repository.QuestionRepository
}

func NewQuestionService(repository repository.QuestionRepository) QuestionService {
	return &questionService{repository: repository}
}

func (s *questionService) Create(ctx context.Context, question *models.Question) error {
	if question.Text == "" {
		return errors.New("question text cannot be empty")
	}
	return s.repository.Create(ctx, question)
}

func (s *questionService) GetAll(ctx context.Context) ([]models.Question, error) {
	return s.repository.GetAll(ctx)
}

func (s *questionService) GetByID(ctx context.Context, id uint) (*models.Question, error) {
	return s.repository.GetByID(ctx, id)
}

func (s *questionService) GetWithAnswers(ctx context.Context, id uint) (*models.Question, error) {
	return s.repository.GetWithAnswers(ctx, id)
}

func (s *questionService) Delete(ctx context.Context, id uint) error {
	return s.repository.Delete(ctx, id)
}