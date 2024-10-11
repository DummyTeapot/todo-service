package model

import "time"

type Task struct {
	ID          int64     `bun:"id,pk,autoincrement" json:"id"`
	Title       string    `bun:"title,notnull" json:"title"`
	Description string    `bun:"description" json:"description"`
	CreatedAt   time.Time `bun:"created_at,default:current_timestamp" json:"created_at"`
	Completed   bool      `bun:"completed,notnull,default:false" json:"completed"`
}
