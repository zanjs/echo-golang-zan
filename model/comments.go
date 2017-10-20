package model

import (
	"time"
)

type Comments struct {
	Id              int       `xorm:"not null pk autoincr INT(10)"`
	Index           string    `xorm:"not null VARCHAR(191)"`
	GroupId         string    `xorm:"not null VARCHAR(191)"`
	UserId          string    `xorm:"not null VARCHAR(191)"`
	Alias           string    `xorm:"not null VARCHAR(191)"`
	Avatar          string    `xorm:"not null VARCHAR(191)"`
	Comment         string    `xorm:"default '''' VARCHAR(191)"`
	ProductComment  string    `xorm:"default 'NULL' TEXT"`
	ProductJson     string    `xorm:"default 'NULL' LONGTEXT"`
	TotalPrice      string    `xorm:"default 0.00 DECIMAL(12,2)"`
	Name            string    `xorm:"default 'NULL' VARCHAR(100)"`
	Phone           string    `xorm:"default 'NULL' VARCHAR(191)"`
	Wechat          string    `xorm:"default 'NULL' VARCHAR(100)"`
	Address         string    `xorm:"default 'NULL' VARCHAR(191)"`
	LocationAddress string    `xorm:"default 'NULL' VARCHAR(191)"`
	CreatedAt       time.Time `xorm:"default 'NULL' TIMESTAMP"`
	UpdatedAt       time.Time `xorm:"default 'NULL' TIMESTAMP"`
}
