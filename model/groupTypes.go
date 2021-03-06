package model

import (
	"time"
)

type GroupTypes struct {
	Id          int       `xorm:"not null pk autoincr INT(10)"`
	Name        string    `xorm:"not null VARCHAR(100)"`
	Path        string    `xorm:"not null VARCHAR(100)"`
	Description string    `xorm:"not null VARCHAR(191)"`
	CreatedAt   time.Time `xorm:"default 'NULL' TIMESTAMP"`
	UpdatedAt   time.Time `xorm:"default 'NULL' TIMESTAMP"`
}
