package model

import (
	"time"
)

type Groups struct {
	Id               int       `xorm:"not null pk autoincr INT(10)"`
	Title            string    `xorm:"default 'NULL' VARCHAR(255)"`
	Description      string    `xorm:"not null TEXT"`
	TypeId           int       `xorm:"not null INT(11)"`
	HeadId           string    `xorm:"not null VARCHAR(50)"`
	Alias            string    `xorm:"not null VARCHAR(191)"`
	Avatar           string    `xorm:"not null VARCHAR(191)"`
	Image            string    `xorm:"default 'NULL' TEXT"`
	Open             int       `xorm:"not null default 1 INT(11)"`
	CurrencyId       string    `xorm:"default 'NULL' VARCHAR(50)"`
	QrCodeV          int       `xorm:"default 1 TINYINT(4)"`
	QrCode           string    `xorm:"default 'NULL' VARCHAR(255)"`
	QrCodePath       string    `xorm:"default 'NULL' VARCHAR(255)"`
	RequiredUName    int       `xorm:"default 0 TINYINT(1)"`
	RequiredUPhone   int       `xorm:"default 0 TINYINT(1)"`
	RequiredUWechat  int       `xorm:"default 0 TINYINT(1)"`
	RequiredUAddress int       `xorm:"default 0 TINYINT(1)"`
	CreatedAt        time.Time `xorm:"default 'NULL' TIMESTAMP"`
	UpdatedAt        time.Time `xorm:"default 'NULL' TIMESTAMP"`
}
