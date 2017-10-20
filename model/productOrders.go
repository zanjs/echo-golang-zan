package model

import (
	"time"
)

type ProductOrders struct {
	Id        int       `xorm:"not null pk autoincr INT(10)"`
	GroupId   string    `xorm:"not null VARCHAR(191)"`
	UserId    string    `xorm:"not null VARCHAR(191)"`
	ProductId string    `xorm:"not null VARCHAR(191)"`
	Quantity  int       `xorm:"default NULL INT(11)"`
	CreatedAt time.Time `xorm:"default 'NULL' TIMESTAMP"`
	UpdatedAt time.Time `xorm:"default 'NULL' TIMESTAMP"`
}
