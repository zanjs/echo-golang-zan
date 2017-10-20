package model

import (
	"time"
)

type Pages struct {
	Id              int       `xorm:"not null pk autoincr INT(10)"`
	AuthorId        int       `xorm:"not null INT(11)"`
	Title           string    `xorm:"not null VARCHAR(191)"`
	Excerpt         string    `xorm:"default 'NULL' TEXT"`
	Body            string    `xorm:"default 'NULL' TEXT"`
	Image           string    `xorm:"default 'NULL' VARCHAR(191)"`
	Slug            string    `xorm:"not null unique VARCHAR(191)"`
	MetaDescription string    `xorm:"default 'NULL' TEXT"`
	MetaKeywords    string    `xorm:"default 'NULL' TEXT"`
	Status          string    `xorm:"not null default ''INACTIVE'' ENUM('INACTIVE','ACTIVE')"`
	CreatedAt       time.Time `xorm:"default 'NULL' TIMESTAMP"`
	UpdatedAt       time.Time `xorm:"default 'NULL' TIMESTAMP"`
}
