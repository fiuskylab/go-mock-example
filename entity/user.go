package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Tabler interface {
	TableName() string
}

type UserModel struct {
	ID        uuid.UUID `gorm:"primaryKey;default:gen_random_uuid()"`
	Username  string    `db:"username" json:"username" gorm:"varchar(128);unique;not null"`
	Password  string    `db:"password" json:"password" gorm:"varchar(128);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (u UserModel) TableName() string {
	return "user"
}

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
