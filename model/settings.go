package model

type Settings struct {
	Id          int    `xorm:"not null pk autoincr INT(10)"`
	Key         string `xorm:"not null unique VARCHAR(191)"`
	DisplayName string `xorm:"not null VARCHAR(191)"`
	Value       string `xorm:"not null TEXT"`
	Details     string `xorm:"default 'NULL' TEXT"`
	Type        string `xorm:"not null VARCHAR(191)"`
	Order       int    `xorm:"not null default 1 INT(11)"`
	Group       string `xorm:"default 'NULL' VARCHAR(191)"`
}
