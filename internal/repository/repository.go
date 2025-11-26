// Package repository defines interfaces and their implementations
// for repositories responsible for managing Question and Answer data persistence
// for repository layer of question-answer API.
package repository

import (
	"context"
	"errors"

	"github.com/ppb03/qna-api/internal/model"
)

var (
	ErrQuestionNotFound = errors.New("no question with such ID")
	ErrAnswerNotFound   = errors.New("no answer with such ID")
)

// QuestionRepository defines the interface for operations related to questions on repository layer.
//
// Standart implementation interacts with PostgreSQL and can be obtained via NewPostgresQuestionRepository() function.
type QuestionRepository interface {
	// Create creates a new question and persists it to the database.
	Create(ctx context.Context, question *model.Question) (*model.Question, error)

	// Delete removes a question from the database based on its ID.
	Delete(ctx context.Context, id uint) error

	// GetByID retrieves a question from the database based on its ID along with its associated answers.
	GetByID(ctx context.Context, id uint) (*model.Question, error)

	// GetAll retrieves all questions from the database.
	GetAll(ctx context.Context) ([]model.Question, error)
}

// AnswerRepository defines the interface for operations related to answers on repository layer.
//
// Standart implementation interacts with PostgreSQL and can be obtained via NewPostgresAnswerRepository() function.
type AnswerRepository interface {
	// Create creates a new answer and persists it to the database.
	Create(ctx context.Context, answer *model.Answer) (*model.Answer, error)

	// Delete removes an answer from the database based on its ID.
	Delete(ctx context.Context, id uint) error

	// GetByID retrieves an answer from the database based on its ID.
	GetByID(ctx context.Context, id uint) (*model.Answer, error)
}
