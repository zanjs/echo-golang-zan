package model

import (
	"time"
)

type Orders struct {
	Id           int       `xorm:"not null pk autoincr INT(10)"`
	GroupId      string    `xorm:"not null VARCHAR(191)"`
	Alias        string    `xorm:"not null VARCHAR(191)"`
	Avatar       string    `xorm:"not null VARCHAR(191)"`
	Comment      string    `xorm:"not null VARCHAR(191)"`
	Products     string    `xorm:"not null VARCHAR(191)"`
	ProductsDesc string    `xorm:"not null TEXT"`
	Phone        string    `xorm:"not null VARCHAR(191)"`
	Address      string    `xorm:"not null VARCHAR(191)"`
	CreatedAt    time.Time `xorm:"default 'NULL' TIMESTAMP"`
	UpdatedAt    time.Time `xorm:"default 'NULL' TIMESTAMP"`
}
