package models

import (
	"time"

	"github.com/google/uuid"
)

type Document struct {
	// TODO: make tags, description and duedate optional
	ID          uuid.UUID  `gorm:"primary_key;type:char(36);default:(UUID())" json:"id"`
	Code        string     `gorm:"not null;unique" json:"code"`
	UserID      uuid.UUID  `gorm:"not null" json:"user_id"`
	Title       string     `gorm:"not null" json:"title"`
	Tags        *string    `json:"tags"`
	Description *string    `json:"description"`
	IssuedAt    time.Time  `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp" json:"issued_at"`
	DueDate     *time.Time `gorm:"default:NULL" json:"due_date"`

	User User `gorm:"foreignKey:UserID;references:ID" json:"user"`
}

type GetDocumentResponse struct {
	Page  int        `json:"page"`
	Limit int        `json:"limit"`
	Data  []Document `json:"data"`
	Total int        `json:"count"`
}

type CreateDocumentValidator struct {
	Title       string     `json:"title" validate:"required,min=3"`
	Code        string     `json:"code" validate:"required,min=5"`
	Tags        *string    `json:"tags"`
	Description *string    `json:"description"`
	IssuedAt    time.Time  `json:"issued_at" validate:"required"`
	DueDate     *time.Time `json:"due_date"`
}
