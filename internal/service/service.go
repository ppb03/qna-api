// Package service defines interfaces for service layer of question-answer API.
package service

import (
	"context"
	"errors"
	"log/slog"
	"regexp"

	"github.com/ppb03/qna-api/internal/model"
)

// Client-side errors
var (
	ErrEmptyText         = errors.New("text cannot be empty")
	ErrEmptyUserID       = errors.New("user ID cannot be empty")
	ErrInvalidUserID     = errors.New("user ID must be a valid UUID")
	ErrQuestionNotExists = errors.New("no question with such ID")
	ErrAnswerNotExists   = errors.New("no answer with such ID")
)

// Internal errors
var (
	ErrRepositoryFailure = errors.New("repository failure")
)

// QuestionService defines the interface for operations related to questions on service layer.
//
// Standart implementation can be obtained via NewQuestionService() function.
type QuestionService interface {
	// Create creates a new question and persists it in database via underlying repository
	Create(ctx context.Context, text string) (*model.Question, error)

	// Delete removes a question from the database based on its ID via underlying repository.
	Delete(ctx context.Context, id uint) error

	// GetAll retrieves all questions from the database via underlying repository.
	GetAll(ctx context.Context) ([]model.Question, error)

	// GetByID retrieves a question from the database based on its ID along with its associated answers via underlying repository.
	GetByID(ctx context.Context, id uint) (*model.Question, error)
}

// AnswerService defines the interface for operations related to answers on service layer.
//
// Standart implementation can be obtained via NewAnswerService() function.
type AnswerService interface {
	// Create creates a new answer and persists it to the database via underlying repository.
	Create(ctx context.Context, questionID uint, userID, text string) (*model.Answer, error)

	// Delete removes an answer from the database based on its ID via underlying repository.
	Delete(ctx context.Context, id uint) error

	// GetByID retrieves an answer from the database based on its ID via underlying repository.
	GetByID(ctx context.Context, id uint) (*model.Answer, error)
}

func internalError(err, errClass error) error {
	joinedErr := errors.Join(errClass, err)
	slog.Error("unexpected internal error: " + joinedErr.Error())
	return joinedErr
}

// ! Should have used the uuid.UUID type from the external library for UserID initially, but I thought of it too late.
func isValidUUID(uuid string) bool {
	uuidRegex := regexp.MustCompile(`^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$`)
	return uuidRegex.MatchString(uuid)
}
