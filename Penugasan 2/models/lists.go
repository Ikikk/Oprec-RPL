package models

// import (
// 	"time"
// )

type Lists struct {
	ID          uint64 `json:"id"`
	Title       string `json:"title"`
	Tag         uint64 `json:"tag"`
	Description string `json:"description"`
	Check       bool   `json:"check"`
}
