package storage

import (
	"fmt"

	"github.com/fiuskylab/go-mock-example/common"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage struct {
	PGSQL  *gorm.DB
	Common *common.Common
}

func NewStorage(c *common.Common) (*Storage, error) {
	s := Storage{
		Common: c,
	}

	var pgsqlURL string
	url := `host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`

	pgsqlURL = fmt.Sprintf(
		url,
		c.Env.PGSQL_HOST,
		c.Env.PGSQL_PORT,
		c.Env.PGSQL_USER,
		c.Env.PGSQL_PASSWORD,
		c.Env.PGSQL_DBNAME,
	)

	db, err := gorm.Open(postgres.Open(pgsqlURL), &gorm.Config{})

	if err != nil {
		c.Log.Error(err.Error())
		return &s, err
	}
	s.PGSQL = db

	return &s, nil
}
