package model

import (
	"time"
)

type WeUsers struct {
	Id            int       `xorm:"not null pk autoincr INT(10)"`
	Name          string    `xorm:"not null VARCHAR(191)"`
	Email         string    `xorm:"not null unique VARCHAR(191)"`
	Password      string    `xorm:"not null VARCHAR(191)"`
	OpenId        string    `xorm:"not null unique VARCHAR(191)"`
	UnionId       string    `xorm:"default 'NULL' index VARCHAR(191)"`
	Nickname      string    `xorm:"not null VARCHAR(191)"`
	Avatar        string    `xorm:"not null VARCHAR(191)"`
	RememberToken string    `xorm:"default 'NULL' VARCHAR(100)"`
	UName         string    `xorm:"default 'NULL' VARCHAR(50)"`
	UPhone        string    `xorm:"default 'NULL' VARCHAR(50)"`
	UWechat       string    `xorm:"default 'NULL' VARCHAR(50)"`
	UAddress      string    `xorm:"default 'NULL' VARCHAR(255)"`
	CreatedAt     time.Time `xorm:"default 'NULL' TIMESTAMP"`
	UpdatedAt     time.Time `xorm:"default 'NULL' TIMESTAMP"`
}
