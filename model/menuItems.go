package model

import (
	"time"
)

type MenuItems struct {
	Id         int       `xorm:"not null pk autoincr INT(10)"`
	MenuId     int       `xorm:"default NULL index INT(10)"`
	Title      string    `xorm:"not null VARCHAR(191)"`
	Url        string    `xorm:"not null VARCHAR(191)"`
	Target     string    `xorm:"not null default ''_self'' VARCHAR(191)"`
	IconClass  string    `xorm:"default 'NULL' VARCHAR(191)"`
	Color      string    `xorm:"default 'NULL' VARCHAR(191)"`
	ParentId   int       `xorm:"default NULL INT(11)"`
	Order      int       `xorm:"not null INT(11)"`
	CreatedAt  time.Time `xorm:"default 'NULL' TIMESTAMP"`
	UpdatedAt  time.Time `xorm:"default 'NULL' TIMESTAMP"`
	Route      string    `xorm:"default 'NULL' VARCHAR(191)"`
	Parameters string    `xorm:"default 'NULL' TEXT"`
}
