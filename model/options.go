package model

import (
	"time"
)

type Options struct {
	Id          int       `xorm:"not null pk autoincr INT(10)"`
	GroupId     string    `xorm:"not null VARCHAR(50)"`
	UserId      string    `xorm:"not null VARCHAR(50)"`
	Source      string    `xorm:"not null VARCHAR(50)"`
	ShareUserId string    `xorm:"not null VARCHAR(50)"`
	Ip          string    `xorm:"not null VARCHAR(100)"`
	CreatedAt   time.Time `xorm:"default 'NULL' TIMESTAMP"`
	UpdatedAt   time.Time `xorm:"default 'NULL' TIMESTAMP"`
}
