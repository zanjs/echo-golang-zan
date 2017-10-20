package model

import (
	"time"
)

type Products struct {
	Id        int       `xorm:"not null pk autoincr INT(10)"`
	GroupId   string    `xorm:"not null VARCHAR(50)"`
	Name      string    `xorm:"not null VARCHAR(100)"`
	Price     string    `xorm:"not null default 0.00 DECIMAL(12,2)"`
	Quantity  int       `xorm:"default NULL INT(11)"`
	Sell      int       `xorm:"not null default 0 INT(11)"`
	CreatedAt time.Time `xorm:"default 'NULL' TIMESTAMP"`
	UpdatedAt time.Time `xorm:"default 'NULL' TIMESTAMP"`
}
