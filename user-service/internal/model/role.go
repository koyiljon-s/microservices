package models

import (
    "github.com/google/uuid"
    "gorm.io/gorm"
)

type Role struct {
    ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
    Name        string    `gorm:"size:100;unique;not null"`
    Description string    `gorm:"type:text"`
    Users       []User    `gorm:"many2many:user_roles;"`
}

func (r *Role) BeforeCreate(tx *gorm.DB) error {
    if r.ID == uuid.Nil {
        r.ID = uuid.New()
    }
    return nil
}