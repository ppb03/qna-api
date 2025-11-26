package repository

import (
	"context"
	"errors"

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

func (r *postgresAnswerRepository) Create(ctx context.Context, answer *model.Answer) (*model.Answer, error) {
	return answer, r.db.WithContext(ctx).Create(answer).Error
}

func (r *postgresAnswerRepository) Delete(ctx context.Context, id uint) error {
	if err := r.db.WithContext(ctx).Delete(&model.Answer{}, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrAnswerNotFound
		}
		return err
	}
	return nil
}

func (r *postgresAnswerRepository) GetByID(ctx context.Context, id uint) (*model.Answer, error) {
	var answer model.Answer
	if err := r.db.WithContext(ctx).First(&answer, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrAnswerNotFound
		}
		return nil, err
	}
	return &answer, nil
}