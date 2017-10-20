package model

import (
	"time"
)

type Placards struct {
	Id          int       `xorm:"not null pk autoincr INT(10)"`
	Description string    `xorm:"not null TEXT"`
	TypeId      string    `xorm:"not null VARCHAR(191)"`
	CreatedAt   time.Time `xorm:"default 'NULL' TIMESTAMP"`
	UpdatedAt   time.Time `xorm:"default 'NULL' TIMESTAMP"`
}