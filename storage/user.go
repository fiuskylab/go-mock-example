package storage

import (
	"github.com/fiuskylab/go-mock-example/entity"
	"golang.org/x/crypto/bcrypt"
)

func (s *Storage) CreateUser(u entity.UserRequest) (entity.UserModel, error) {
	m := new(entity.UserModel)
	m.Username = u.Username
	hashPassBytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return entity.UserModel{}, err
	}

	m.Password = string(hashPassBytes)

	if err := s.PGSQL.Create(m).Error; err != nil {
		return entity.UserModel{}, err
	}
	return *m, nil
}
