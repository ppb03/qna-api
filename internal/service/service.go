// Package service defines interfaces for service layer of question-answer API.
package service

import (
	"context"

	"github.com/ppb03/qna-api/internal/model"
)

// QuestionService defines the interface for operations related to questions on service layer.
//
// Standart implementation can be obtained via NewQuestionService() function.
type QuestionService interface {
	// Create creates a new question and persists it in database via underlying repository
	Create(ctx context.Context, question *model.Question) error
	
	// Delete removes a question from the database based on its ID via underlying repository.
	Delete(ctx context.Context, id uint) error

	// GetAll retrieves all questions from the database via underlying repository.
	GetAll(ctx context.Context) ([]model.Question, error)

	GetByID(ctx context.Context, id uint) (*model.Question, error)

	// GetWithAnswers retrieves a question from the database based on its ID along with its associated answers via underlying repository.
	GetWithAnswers(ctx context.Context, id uint) (*model.Question, error)
}

// AnswerService defines the interface for operations related to answers on service layer.
//
// Standart implementation can be obtained via NewAnswerService() function.
type AnswerService interface {
	// Create creates a new answer and persists it to the database via underlying repository.
	Create(ctx context.Context, answer *model.Answer) error

	// Delete removes an answer from the database based on its ID via underlying repository.
	Delete(ctx context.Context, id uint) error

	// GetByID retrieves an answer from the database based on its ID via underlying repository.
	GetByID(ctx context.Context, id uint) (*model.Answer, error)
}
