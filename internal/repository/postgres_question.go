package repository

import (
	"context"
	"errors"

	"github.com/ppb03/qna-api/internal/model"

	"gorm.io/gorm"
)

type postgresQuestionRepository struct {
	db *gorm.DB
}

// NewPostgresQuestionRepository creates QuestionRepository instance which interacts with PostgreSQL database
func NewPostgresQuestionRepository(db *gorm.DB) QuestionRepository {
	return &postgresQuestionRepository{db: db}
}

func (r *postgresQuestionRepository) Create(ctx context.Context, question *model.Question) (*model.Question, error) {
	return question, r.db.WithContext(ctx).Create(question).Error
}

func (r *postgresQuestionRepository) Delete(ctx context.Context, id uint) error {
	if err := r.db.WithContext(ctx).Delete(&model.Question{}, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrQuestionNotFound
		}
		return err
	}
	return nil
}

func (r *postgresQuestionRepository) GetByID(ctx context.Context, id uint) (*model.Question, error) {
	var question model.Question
	if err := r.db.WithContext(ctx).Preload("Answers").First(&question, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrQuestionNotFound
		}
		return nil, err
	}
	return &question, nil
}

func (r *postgresQuestionRepository) GetAll(ctx context.Context) ([]model.Question, error) {
	var questions []model.Question
	err := r.db.WithContext(ctx).Find(&questions).Error
	return questions, err
}
