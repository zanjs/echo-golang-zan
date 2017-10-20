package model

type PermissionGroups struct {
	Id   int    `xorm:"not null pk autoincr INT(10)"`
	Name string `xorm:"not null unique VARCHAR(191)"`
}
