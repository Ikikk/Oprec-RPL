package models

import (
	"time"
)

type Lists struct {
	ID          uint64    `json:"id"`
	Title       string    `json:"title"`
	Tag         string    `json:"tag"`
	Description string    `json:"description"`
	Check       bool      `json:"check"`
	Created_at  time.Time `json:"created_at"`
	Deleted_at  time.Time `json:"deleted_at"`
	Updated_at  time.Time `json:"updated_at"`
}
