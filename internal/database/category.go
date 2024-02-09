package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Category struct {
	Db *sql.DB
	ID string
	Name string
	Description string
}

func NewCategory(db *sql.DB) *Category {
	return &Category{Db: db}
}

func (c *Category) Create(name, description string) (Category, error) {
	id := uuid.New().String()
	_, err := c.Db.Exec("insert into categories (id, name, description) values ($1, $2, $3)", id, name, description)
	if err != nil {
		return Category{}, err
	}

	return Category{ID: id, Name: name, Description: description}, nil
}