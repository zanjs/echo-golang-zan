package model

import (
	"time"
)

type Users struct {
	Id            int       `xorm:"not null pk autoincr INT(10)"`
	RoleId        int       `xorm:"default NULL INT(11)"`
	Name          string    `xorm:"not null VARCHAR(191)"`
	Email         string    `xorm:"not null unique VARCHAR(191)"`
	Avatar        string    `xorm:"default ''users/default.png'' VARCHAR(191)"`
	Password      string    `xorm:"not null VARCHAR(191)"`
	RememberToken string    `xorm:"default 'NULL' VARCHAR(100)"`
	CreatedAt     time.Time `xorm:"default 'NULL' TIMESTAMP"`
	UpdatedAt     time.Time `xorm:"default 'NULL' TIMESTAMP"`
}
