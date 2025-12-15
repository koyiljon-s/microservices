package models

import (
    "time"
    "github.com/google/uuid"
    "gorm.io/gorm"
)

// User model
type User struct {
    ID           uuid.UUID   `gorm:"type:uuid;primaryKey"`
    Username     string      `gorm:"size:255;unique;not null"`
    Email        string      `gorm:"size:255;unique;not null"`
    PasswordHash string      `gorm:"size:255;not null"`
    IsActive     bool        `gorm:"default:true"`
    IsVerified   bool        `gorm:"default:false"`
    CreatedAt    time.Time   `gorm:"autoCreateTime"`
    LastLogin    *time.Time
    
    Profile      UserProfile `gorm:"constraint:OnDelete:CASCADE"`
    Roles        []Role      `gorm:"many2many:user_roles;constraint:OnDelete:CASCADE"`
    Sessions     []Session   `gorm:"constraint:OnDelete:CASCADE"`
}


func (u *User) BeforeCreate(tx *gorm.DB) error {
    if u.ID == uuid.Nil {
        u.ID = uuid.New()
    }
    return nil
}