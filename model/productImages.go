package model

import (
	"time"
)

type ProductImages struct {
	Id        int       `xorm:"not null pk autoincr INT(10)"`
	ProductId string    `xorm:"not null VARCHAR(50)"`
	Path      string    `xorm:"not null VARCHAR(191)"`
	CreatedAt time.Time `xorm:"default 'NULL' TIMESTAMP"`
	UpdatedAt time.Time `xorm:"default 'NULL' TIMESTAMP"`
}
