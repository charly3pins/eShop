package postgres

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type ConnectionOptions struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}

func NewConnection(c ConnectionOptions) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", c.Host, c.Port, c.User, c.Password, c.DbName)
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	db.SingularTable(true)

	return db, nil
}
