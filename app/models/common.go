package models

import (
	"github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

type ID struct {
	ID uint `json:"id" gorm:"primaryKey"`
}

type Uuid struct {
	Uuid uuid.UUID `json:"uuid" gorm:"not null;index;comment:UUID"`
}

type Timestamps struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SoftDeletes struct {
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
