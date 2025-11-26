package service

import (
	"context"
	"errors"

	"github.com/ppb03/qna-api/internal/model"
	"github.com/ppb03/qna-api/internal/repository"

	"gorm.io/gorm"
)

type answerService struct {
	answerRepository   repository.AnswerRepository
	questionRepository repository.QuestionRepository
}

// NewAnswerService creates AnswerService instance with standart implementation
func NewAnswerService(answerRepository repository.AnswerRepository, questionRepository repository.QuestionRepository) AnswerService {
	return &answerService{answerRepository: answerRepository, questionRepository: questionRepository}
}

func (s *answerService) Create(ctx context.Context, answer *model.Answer) error {
	if answer.Text == "" {
		return errors.New("answer text cannot be empty")
	}
	if answer.UserID == "" {
		return errors.New("user ID cannot be empty")
	}

	_, err := s.questionRepository.GetByID(ctx, answer.QuestionID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("question does not exist")
		}
		return err
	}
	return s.answerRepository.Create(ctx, answer)
}

func (s *answerService) GetByID(ctx context.Context, id uint) (*model.Answer, error) {
	return s.answerRepository.GetByID(ctx, id)
}

func (s *answerService) Delete(ctx context.Context, id uint) error {
	return s.answerRepository.Delete(ctx, id)
}
