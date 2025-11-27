// Package model defines the core domain entities and their data structure.
//
// These types represent the business objects and are used throughout the application layers.
package model

import (
	"time"
)


// Question represents a question entity in the system.
// It contains the question text, creation timestamp, and associated answers.
type Question struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Text      string    `json:"text" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	Answers   []Answer  `json:"answers,omitempty" gorm:"foreignKey:QuestionID;constraint:OnDelete:CASCADE"`
}

// Answer represents an answer to a question in the system.
// It links to a specific question and includes the responder's identity and answer content.
type Answer struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	QuestionID uint      `json:"question_id" gorm:"index;not null"`
	UserID     string    `json:"user_id" gorm:"not null;index"`
	Text       string    `json:"text" gorm:"not null"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
}
