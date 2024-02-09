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

func (c *Category) FindAll() ([]Category, error) {
	rows, err := c.Db.Query("select * from categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var categories []Category
	for rows.Next() {
		var id, name, description string
		if err = rows.Scan(&id, &name, &description); err != nil {
			return nil, err
		}
		categories = append(categories, Category{ID: id, Name: name, Description: description})
	}

	return categories, nil
}