// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0

package db

import (
	"database/sql"
	"time"
)

type Budget struct {
	ID          string         `json:"id"`
	UserID      string         `json:"user_id"`
	Amount      int32          `json:"amount"`
	Description string         `json:"description"`
	CreatedAt   sql.NullTime   `json:"created_at"`
	CreatedBy   string         `json:"created_by"`
	UpdatedAt   time.Time      `json:"updated_at"`
	UpdatedBy   string         `json:"updated_by"`
	DeletedBy   sql.NullString `json:"deleted_by"`
	DeletedAt   sql.NullTime   `json:"deleted_at"`
}

type CategoryTag struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type User struct {
	UserID    string    `json:"user_id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy string    `json:"updated_by"`
}