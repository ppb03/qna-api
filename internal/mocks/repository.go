// Package mocks contains mock implementations of repository interfaces.
//
// These mocks simulate database interactions without requiring actual database connections.
package mocks

import (
	"context"

	"github.com/ppb03/qna-api/internal/model"
	
	"github.com/stretchr/testify/mock"
)

type MockQuestionRepository struct {
	mock.Mock
}

func (m *MockQuestionRepository) Create(ctx context.Context, question *model.Question) (*model.Question, error) {
	args := m.Called(ctx, question)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Question), args.Error(1)
}

func (m *MockQuestionRepository) Delete(ctx context.Context, id uint) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockQuestionRepository) GetByID(ctx context.Context, id uint) (*model.Question, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Question), args.Error(1)
}

func (m *MockQuestionRepository) GetAll(ctx context.Context) ([]model.Question, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Question), args.Error(1)
}

type MockAnswerRepository struct {
	mock.Mock
}

func (m *MockAnswerRepository) Create(ctx context.Context, answer *model.Answer) (*model.Answer, error) {
	args := m.Called(ctx, answer)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Answer), args.Error(1)
}

func (m *MockAnswerRepository) Delete(ctx context.Context, id uint) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockAnswerRepository) GetByID(ctx context.Context, id uint) (*model.Answer, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Answer), args.Error(1)
}