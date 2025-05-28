package entity

import (
	"time"

	"github.com/uptrace/bun"
)

type GrantRequestsLog struct {
	bun.BaseModel `bun:"grant_requests_log"`

	ID        int64     `bun:",pk,autoincrement"`
	ClientID  int64     `bun:",notnull"`
	IPAddress string    `bun:",notnull"`
	UserAgent string    `bun:",notnull"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`

	// Relations (opsional)
	Client *Client `bun:"rel:belongs-to,join:client_id=id"`
}
