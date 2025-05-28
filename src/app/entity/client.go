package entity

import (
	"time"

	"github.com/uptrace/bun"
)

type Client struct {
	bun.BaseModel `bun:"clients"`

	ID           int64     `bun:",pk,autoincrement"`
	Name         string    `bun:",notnull"`
	Description  string    `bun:",nullzero"`
	RedirectURI  string    `bun:",nullzero"`
	ContactEmail string    `bun:",nullzero"`
	IsActive     bool      `bun:",notnull,default:true"`
	ApiKey       string    `bun:",unique,notnull"`
	ApiSecret    string    `bun:",notnull"`
	CreatedAt    time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt    time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}

type GrantTokenRedis struct {
	ApiKey     string `json:"api_key"`
	GrantToken string `json:"grant_token"`
}

type JwtGrantToken struct {
	Sub      int64  `json:"sub"`
	ApiKey   string `json:"api_key"`
	Type     string `json:"type"`
	Exp      int64  `json:"exp"`
	IssuedAt int64  `json:"issued_at"`
}
