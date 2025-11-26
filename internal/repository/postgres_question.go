package repository

import (
	"context"

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

func (r *postgresQuestionRepository) Create(ctx context.Context, question *model.Question) error {
	return r.db.WithContext(ctx).Create(question).Error
}

func (r *postgresQuestionRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.Question{}, id).Error
}

func (r *postgresQuestionRepository) GetByID(ctx context.Context, id uint) (*model.Question, error) {
	var question model.Question
	err := r.db.WithContext(ctx).First(&question, id).Error
	if err != nil {
		return nil, err
	}
	return &question, nil
}

func (r *postgresQuestionRepository) GetAll(ctx context.Context) ([]model.Question, error) {
	var questions []model.Question
	err := r.db.WithContext(ctx).Find(&questions).Error
	return questions, err
}

func (r *postgresQuestionRepository) GetWithAnswers(ctx context.Context, id uint) (*model.Question, error) {
	var question model.Question
	err := r.db.WithContext(ctx).Preload("Answers").First(&question, id).Error
	if err != nil {
		return nil, err
	}
	return &question, nil
}