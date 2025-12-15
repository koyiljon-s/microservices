package models

import (
    "github.com/google/uuid"
    "gorm.io/gorm"
)

type UserProfile struct {
    ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
    UserID    uuid.UUID `gorm:"type:uuid;uniqueIndex;not null"`
    User      User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	
    Bio       string    `gorm:"type:text"`
    AvatarURL string    `gorm:"size:500"`
    Karma     int       `gorm:"default:0"`
}


func (up *UserProfile) BeforeCreate(tx *gorm.DB) error {
    if up.ID == uuid.Nil {
        up.ID = uuid.New()
    }
    return nil
}