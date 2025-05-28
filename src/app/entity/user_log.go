package entity

import (
	"time"

	"github.com/uptrace/bun"
)

type UserLoginLog struct {
	bun.BaseModel `bun:"table:user_login_logs"`

	ID         int       `bun:"id,pk,autoincrement"`
	Username   string    `bun:"username,notnull"`
	LogJournal string    `bun:"log_journal,notnull"`
	IPAddress  string    `bun:"ip_address,nullzero"`
	UserAgent  string    `bun:"user_agent,nullzero"`
	Success    bool      `bun:"success,notnull,default:true"`
	LoggedAt   time.Time `bun:"logged_at,notnull,default:now()"`
	CreatedAt  time.Time `bun:"created_at,notnull,default:now()"`
}
