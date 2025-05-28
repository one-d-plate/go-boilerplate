package entity

import (
	"time"

	"github.com/uptrace/bun"
)

type Level struct {
	bun.BaseModel `bun:"table:level"`

	ID          int64     `bun:"id,pk,autoincrement"`
	Name        string    `bun:"name,notnull"`
	Value       int       `bun:"value,notnull"`
	Description string    `bun:"description"`
	CreatedAt   time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp"`
	UpdatedAt   time.Time `bun:"updated_at,nullzero,notnull,default:current_timestamp"`
}
