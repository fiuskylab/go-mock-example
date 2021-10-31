package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `gorm:"primaryKey;default:gen_random_uuid()"`
	Username  string    `db:"username" json:"username" gorm:"varchar(128);unique;not null"`
	Password  string    `db:"password" json:"password" gorm:"varchar(128);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
