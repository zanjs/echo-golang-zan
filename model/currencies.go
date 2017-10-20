package model

import (
	"time"
)

type Currencies struct {
	Id           int       `xorm:"not null pk autoincr INT(10)"`
	Name         string    `xorm:"not null VARCHAR(191)"`
	Mark         string    `xorm:"not null VARCHAR(191)"`
	ExchangeRate string    `xorm:"not null VARCHAR(191)"`
	CreatedAt    time.Time `xorm:"default 'NULL' TIMESTAMP"`
	UpdatedAt    time.Time `xorm:"default 'NULL' TIMESTAMP"`
}
