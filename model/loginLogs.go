package model

import (
	"time"
)

type LoginLogs struct {
	Id        int       `xorm:"not null pk autoincr INT(10)"`
	UserId    string    `xorm:"not null VARCHAR(50)"`
	UserName  string    `xorm:"not null VARCHAR(191)"`
	OpenId    string    `xorm:"not null VARCHAR(191)"`
	Ip        string    `xorm:"not null VARCHAR(191)"`
	Device    string    `xorm:"not null VARCHAR(191)"`
	CreatedAt time.Time `xorm:"default 'NULL' TIMESTAMP"`
	UpdatedAt time.Time `xorm:"default 'NULL' TIMESTAMP"`
}
