package models

import "database/sql"

type User struct {
	ID       int            `json:"id" validate:"required"`
	Name     string         `json:"name" validate:"required"`
	Password string         `json:"password" validate:"required"`
	Email    string         `json:"email" validate:"required"`
	Gender   sql.NullString `json:"gender" validate:"required"`
	Mobile   sql.NullString `json:"mobile" validate:"required"`
	// Orders   string `json:"orders" validate:"required"`
	ImageURL sql.NullString `json:"imageurl" validate:"required"`
	Address  sql.NullString `json:"address" validate:"required"`
}
