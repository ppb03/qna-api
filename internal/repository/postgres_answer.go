package repository

import (
	"context"

	"github.com/ppb03/qna-api/internal/model"

	"gorm.io/gorm"
)

type postgresAnswerRepository struct {
	db *gorm.DB
}

// NewPostgresAnswerRepository creates AnswerRepository instance which interacts with PostgreSQL database
func NewPostgresAnswerRepository(db *gorm.DB) AnswerRepository {
	return &postgresAnswerRepository{db: db}
}

func (r *postgresAnswerRepository) Create(ctx context.Context, answer *model.Answer) error {
	return r.db.WithContext(ctx).Create(answer).Error
}

func (r *postgresAnswerRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.Answer{}, id).Error
}

func (r *postgresAnswerRepository) GetByID(ctx context.Context, id uint) (*model.Answer, error) {
	var answer model.Answer
	err := r.db.WithContext(ctx).First(&answer, id).Error
	if err != nil {
		return nil, err
	}
	return &answer, nil
}