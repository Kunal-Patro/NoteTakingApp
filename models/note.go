package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Note struct {
	ID          uuid.UUID      `json:"nid" gorm:"type:uuid;default:uuid_generate_v4();primaryKey;"`
	Title       string         `json:"title" gorm:"type:text;unique_index:idx_title_desc;not null;"`
	Description string         `json:"description" gorm:"type:text;unique_index:idx_title_desc;"`
	CreatedAt   time.Time      `json:"-" gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `json:"-" gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
	NotebookID  uuid.UUID      `json:"notebook_id"`
	Notebook    Notebook       `gorm:"foreignKey:NotebookID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
