package service

import (
	"context"

	"github.com/ppb03/qna-api/internal/model"
	"github.com/ppb03/qna-api/internal/repository"
)

type questionService struct {
	repository repository.QuestionRepository
}

// NewPostgresQuestionRepository creates QuestionRepository instance which interacts with PostgreSQL database
// NewQuestionService creates QuestionService instance with standart implementation
func NewQuestionService(repository repository.QuestionRepository) QuestionService {
	return &questionService{repository: repository}
}

func (qs *questionService) Create(ctx context.Context, question *model.Question) error {
	return qs.repository.Create(ctx, question)
}

func (qs *questionService) GetAll(ctx context.Context) ([]model.Question, error) {
	return qs.repository.GetAll(ctx)
}

func (qs *questionService) GetByID(ctx context.Context, id uint) (*model.Question, error) {
	return qs.repository.GetByID(ctx, id)
}

func (qs *questionService) GetWithAnswers(ctx context.Context, id uint) (*model.Question, error) {
	return qs.repository.GetWithAnswers(ctx, id)
}

func (qs *questionService) Delete(ctx context.Context, id uint) error {
	return qs.repository.Delete(ctx, id)
}