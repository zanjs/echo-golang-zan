package model

import (
	"time"
)

type Images struct {
	Id        int       `xorm:"not null pk autoincr INT(10)"`
	GroupId   string    `xorm:"not null VARCHAR(191)"`
	Filename  string    `xorm:"not null VARCHAR(191)"`
	CreatedAt time.Time `xorm:"default 'NULL' TIMESTAMP"`
	UpdatedAt time.Time `xorm:"default 'NULL' TIMESTAMP"`
}
