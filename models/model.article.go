package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ModelArticle struct {
	ID        string    `json:"id" gorm:"primary_key"`
	Title     string    `json:"first_name" gorm:"type:varchar;  not null"`
	Body      string    `json:"last_name" gorm:"type:varchar; not null"`
	Author    string    `json:"email" gorm:"type:varchar; unique; not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (m *ModelArticle) BeforeCreate(db *gorm.DB) error {
	m.ID = uuid.NewString()
	m.CreatedAt = time.Now()
	return nil
}

func (m *ModelArticle) BeforeUpdate(db *gorm.DB) error {
	m.ID = uuid.NewString()
	m.UpdatedAt = time.Now()
	return nil
}
