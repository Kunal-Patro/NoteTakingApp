package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID           uuid.UUID `json:"uid" gorm:"type:uuid;default:uuid_generate_v4();primaryKey;"`
	Email        string    `json:"email" gorm:"type:varchar(128);index:idx_users_email,unique;not null;unique;"`
	PasswordHash string    `json:"-" gorm:"type:varchar(512);not null;"`
	FirstName    string    `json:"first_name" gorm:"type:varchar(32);"`
	LastName     string    `json:"last_name" gorm:"type:varchar(32);"`
	Phone        string    `json:"phone" gorm:"type:varchar(16);"`
	DateOfBirth  string    `json:"date_of_birth" gorm:"type:date;"`
	// Notebooks    []Notebook     `json:"notebooks" gorm:"foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	CreatedAt time.Time      `json:"-" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"-" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
