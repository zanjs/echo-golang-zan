package model

type DataRows struct {
	Id          int    `xorm:"not null pk autoincr INT(10)"`
	DataTypeId  int    `xorm:"not null index INT(10)"`
	Field       string `xorm:"not null VARCHAR(191)"`
	Type        string `xorm:"not null VARCHAR(191)"`
	DisplayName string `xorm:"not null VARCHAR(191)"`
	Required    int    `xorm:"not null default 0 TINYINT(1)"`
	Browse      int    `xorm:"not null default 1 TINYINT(1)"`
	Read        int    `xorm:"not null default 1 TINYINT(1)"`
	Edit        int    `xorm:"not null default 1 TINYINT(1)"`
	Add         int    `xorm:"not null default 1 TINYINT(1)"`
	Delete      int    `xorm:"not null default 1 TINYINT(1)"`
	Details     string `xorm:"default 'NULL' TEXT"`
	Order       int    `xorm:"not null default 1 INT(11)"`
}
