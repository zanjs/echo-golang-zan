package model

import (
	"time"
)

type Posts struct {
	Id              int       `xorm:"not null pk autoincr INT(10)"`
	AuthorId        int       `xorm:"not null INT(11)"`
	CategoryId      int       `xorm:"default NULL INT(11)"`
	Title           string    `xorm:"not null VARCHAR(191)"`
	SeoTitle        string    `xorm:"default 'NULL' VARCHAR(191)"`
	Excerpt         string    `xorm:"default 'NULL' TEXT"`
	Body            string    `xorm:"not null TEXT"`
	Image           string    `xorm:"default 'NULL' VARCHAR(191)"`
	Slug            string    `xorm:"not null unique VARCHAR(191)"`
	MetaDescription string    `xorm:"default 'NULL' TEXT"`
	MetaKeywords    string    `xorm:"default 'NULL' TEXT"`
	Status          string    `xorm:"not null default ''DRAFT'' ENUM('PUBLISHED','DRAFT','PENDING')"`
	Featured        int       `xorm:"not null default 0 TINYINT(1)"`
	CreatedAt       time.Time `xorm:"default 'NULL' TIMESTAMP"`
	UpdatedAt       time.Time `xorm:"default 'NULL' TIMESTAMP"`
}
