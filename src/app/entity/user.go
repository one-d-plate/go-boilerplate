package entity

import (
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"users"`

	ID           int64     `bun:",pk,autoincrement"`
	Username     string    `bun:",unique,notnull"`
	Password     string    `bun:",notnull"`
	Name         string    `bun:",notnull"`
	TokenVersion int       `bun:",notnull,default:0"`
	Role         string    `bun:",notnull,default:'1'"`
	Level        string    `bun:",notnull,default:'1.1'"`
	CreatedAt    time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt    time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}

const (
	RoleUser       string = "1"
	RoleAdmin      string = "2"
	RoleSuperAdmin string = "3"
)

type UserJwt struct {
	ID           int64  `json:"id"`
	Username     string `json:"username"`
	Role         string `json:"role"`
	Level        string `json:"level"`
	TokenVersion int    `json:"token_version"`
	Exp          int64  `json:"exp"`
	IssueAt      int64  `json:"issued_at"`
}
