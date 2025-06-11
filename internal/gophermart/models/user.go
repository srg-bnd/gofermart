package models

import (
	"github.com/google/uuid"
	"time"
)

type UserStatus int

//const UserStatusDisabled UserStatus = 0
//const UserStatusActive UserStatus = 1

type User struct {
	ID           uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Login        string     `gorm:"uniqueIndex;not null"`
	Password     string     `gorm:"-" json:"-"`
	PasswordHash string     `gorm:"not null" json:"-"`
	Status       UserStatus `gorm:"default:0"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
