package models

import "github.com/google/uuid"

type Auth struct {
	AuthID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID          uuid.UUID
	TokenExpiration float64 `gorm:"type:float;"`
	User            User    `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
