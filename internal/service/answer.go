package service

import (
	"context"
	"errors"

	"github.com/ppb03/qna-api/internal/model"
	"github.com/ppb03/qna-api/internal/repository"
)

type answerService struct {
	answerRepository   repository.AnswerRepository
	questionRepository repository.QuestionRepository
}

// NewAnswerService creates AnswerService instance with standart implementation
func NewAnswerService(answerRepository repository.AnswerRepository, questionRepository repository.QuestionRepository) AnswerService {
	return &answerService{answerRepository: answerRepository, questionRepository: questionRepository}
}

func (as *answerService) Create(ctx context.Context, questionID uint, userID, text string) (*model.Answer, error) {
	if text == "" {
		return nil, ErrEmptyText
	}
	
	if userID == "" {
		return nil, ErrEmptyUserID
	}

	if !isValidUUID(userID) {
		return nil, ErrInvalidUserID
	}

	_, err := as.questionRepository.GetByID(ctx, questionID)
	if err != nil {
		if errors.Is(err, repository.ErrQuestionNotFound) {
			return nil, ErrQuestionNotExists
		}
		return nil, internalError(err, ErrRepositoryFailure)
	}

	answer, err := as.answerRepository.Create(ctx, &model.Answer{QuestionID: questionID, UserID: userID, Text: text})
	if err != nil {
		return nil, internalError(err, ErrRepositoryFailure)
	}
	return answer, nil
}

func (as *answerService) Delete(ctx context.Context, id uint) error {
	if err := as.answerRepository.Delete(ctx, id); err != nil {
		if errors.Is(err, repository.ErrAnswerNotFound) {
			return ErrAnswerNotExists
		}
		return internalError(err, ErrRepositoryFailure)
	}
	return nil
}

func (as *answerService) GetByID(ctx context.Context, id uint) (*model.Answer, error) {
	answer, err := as.answerRepository.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrAnswerNotFound) {
			return nil, ErrAnswerNotExists
		}
		return nil, internalError(err, ErrRepositoryFailure)
	}
	return answer, nil
}
