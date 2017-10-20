package model

import (
	"time"
)

type Categories struct {
	Id        int       `xorm:"not null pk autoincr INT(10)"`
	ParentId  int       `xorm:"default NULL index INT(10)"`
	Order     int       `xorm:"not null default 1 INT(11)"`
	Name      string    `xorm:"not null VARCHAR(191)"`
	Slug      string    `xorm:"not null unique VARCHAR(191)"`
	CreatedAt time.Time `xorm:"default 'NULL' TIMESTAMP"`
	UpdatedAt time.Time `xorm:"default 'NULL' TIMESTAMP"`
}
