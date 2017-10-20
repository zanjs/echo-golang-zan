package model

import (
	"time"
)

type PasswordResets struct {
	Email     string    `xorm:"not null index VARCHAR(191)"`
	Token     string    `xorm:"not null VARCHAR(191)"`
	CreatedAt time.Time `xorm:"default 'NULL' TIMESTAMP"`
}
