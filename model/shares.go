package model

import (
	"time"
)

type Shares struct {
	Id        int       `xorm:"not null pk autoincr INT(10)"`
	Title     string    `xorm:"not null VARCHAR(191)"`
	Path      string    `xorm:"not null VARCHAR(191)"`
	CreatedAt time.Time `xorm:"default 'NULL' TIMESTAMP"`
	UpdatedAt time.Time `xorm:"default 'NULL' TIMESTAMP"`
}
