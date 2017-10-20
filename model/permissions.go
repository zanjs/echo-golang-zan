package model

import (
	"time"
)

type Permissions struct {
	Id                int       `xorm:"not null pk autoincr INT(10)"`
	Key               string    `xorm:"not null index VARCHAR(191)"`
	TableName         string    `xorm:"default 'NULL' VARCHAR(191)"`
	CreatedAt         time.Time `xorm:"default 'NULL' TIMESTAMP"`
	UpdatedAt         time.Time `xorm:"default 'NULL' TIMESTAMP"`
	PermissionGroupId int       `xorm:"default NULL INT(10)"`
}
