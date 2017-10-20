package model

import (
	"time"
)

type Roles struct {
	Id          int       `xorm:"not null pk autoincr INT(10)"`
	Name        string    `xorm:"not null unique VARCHAR(191)"`
	DisplayName string    `xorm:"not null VARCHAR(191)"`
	CreatedAt   time.Time `xorm:"default 'NULL' TIMESTAMP"`
	UpdatedAt   time.Time `xorm:"default 'NULL' TIMESTAMP"`
}
