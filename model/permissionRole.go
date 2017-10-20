package model

type PermissionRole struct {
	PermissionId int `xorm:"not null pk index INT(10)"`
	RoleId       int `xorm:"not null pk index INT(10)"`
}
