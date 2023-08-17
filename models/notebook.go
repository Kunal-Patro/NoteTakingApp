package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Notebook struct {
	ID   uuid.UUID `json:"nbid" gorm:"type:uuid;default:uuid_generate_v4();primaryKey;"`
	Name string    `json:"name" gorm:"type:varchar(32); not null;"`
	// Notes     []Note         `json:"notes" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt time.Time      `json:"-" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"-" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
	UserID    uuid.UUID      `json:"user_id"`
	User      User           `json:"-" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
