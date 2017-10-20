package model

import (
	"time"
)

type DataTypes struct {
	Id                  int       `xorm:"not null pk autoincr INT(10)"`
	Name                string    `xorm:"not null unique VARCHAR(191)"`
	Slug                string    `xorm:"not null unique VARCHAR(191)"`
	DisplayNameSingular string    `xorm:"not null VARCHAR(191)"`
	DisplayNamePlural   string    `xorm:"not null VARCHAR(191)"`
	Icon                string    `xorm:"default 'NULL' VARCHAR(191)"`
	ModelName           string    `xorm:"default 'NULL' VARCHAR(191)"`
	PolicyName          string    `xorm:"default 'NULL' VARCHAR(191)"`
	Controller          string    `xorm:"default 'NULL' VARCHAR(191)"`
	Description         string    `xorm:"default 'NULL' VARCHAR(191)"`
	GeneratePermissions int       `xorm:"not null default 0 TINYINT(1)"`
	ServerSide          int       `xorm:"not null default 0 TINYINT(4)"`
	CreatedAt           time.Time `xorm:"default 'NULL' TIMESTAMP"`
	UpdatedAt           time.Time `xorm:"default 'NULL' TIMESTAMP"`
}
