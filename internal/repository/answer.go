package repository

import (
	"context"

	"github.com/ppb03/question-answer-api/internal/models"

	"gorm.io/gorm"
)

type answerRepository struct {
	db *gorm.DB
}

func NewAnswerRepository(db *gorm.DB) AnswerRepository {
	return &answerRepository{db: db}
}

func (r *answerRepository) Create(ctx context.Context, answer *models.Answer) error {
	return r.db.WithContext(ctx).Create(answer).Error
}

func (r *answerRepository) GetByID(ctx context.Context, id uint) (*models.Answer, error) {
	var answer models.Answer
	err := r.db.WithContext(ctx).First(&answer, id).Error
	if err != nil {
		return nil, err
	}
	return &answer, nil
}

func (r *answerRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Answer{}, id).Error
}