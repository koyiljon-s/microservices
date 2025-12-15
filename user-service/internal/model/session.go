package models

import (
    "time"
    "github.com/google/uuid"
    "gorm.io/gorm"
)

type Session struct {
    ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
    UserID     uuid.UUID `gorm:"type:uuid;not null;index"`
    User       User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
    Token      string    `gorm:"size:500;unique;not null"`
    ExpiresAt  time.Time `gorm:"not null"`
    DeviceInfo string    `gorm:"size:500"`
    CreatedAt  time.Time `gorm:"autoCreateTime"`
}

func (s *Session) BeforeCreate(tx *gorm.DB) error {
    if s.ID == uuid.Nil {
        s.ID = uuid.New()
    }
    return nil
}