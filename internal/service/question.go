package service

import (
	"context"
	"errors"

	"github.com/ppb03/qna-api/internal/model"
	"github.com/ppb03/qna-api/internal/repository"	
)

type questionService struct {
	repository repository.QuestionRepository
}

// NewQuestionService creates QuestionService instance with standart implementation
func NewQuestionService(repository repository.QuestionRepository) QuestionService {
	return &questionService{repository: repository}
}

func (qs *questionService) Create(ctx context.Context, text string) (*model.Question, error){
	if text == "" {
		return nil, ErrEmptyText
	}

	question, err := qs.repository.Create(ctx, &model.Question{Text: text})
	if err != nil {
		return nil, internalError(err, ErrRepositoryFailure)
	}
	return question, nil
}

func (qs *questionService) Delete(ctx context.Context, id uint) error {
	if err := qs.repository.Delete(ctx, id); err != nil {
		if errors.Is(err, repository.ErrQuestionNotFound) {
			return ErrQuestionNotExists
		}
		return internalError(err, ErrRepositoryFailure)
	}
	return nil
}

func (qs *questionService) GetByID(ctx context.Context, id uint) (*model.Question, error) {
	question, err := qs.repository.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrQuestionNotFound) {
			return nil, ErrQuestionNotExists
		}
		return nil, internalError(err, ErrRepositoryFailure)
	} 
	return question, nil
}

func (qs *questionService) GetAll(ctx context.Context) ([]model.Question, error) {
	questions, err := qs.repository.GetAll(ctx)
	if err != nil {
		return nil, internalError(err, ErrRepositoryFailure)
	}
	return questions, nil
}

