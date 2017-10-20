package model

import (
	"time"
)

type Translations struct {
	Id         int       `xorm:"not null pk autoincr INT(10)"`
	TableName  string    `xorm:"not null unique(translations_table_name_column_name_foreign_key_locale_unique) VARCHAR(191)"`
	ColumnName string    `xorm:"not null unique(translations_table_name_column_name_foreign_key_locale_unique) VARCHAR(191)"`
	ForeignKey int       `xorm:"not null unique(translations_table_name_column_name_foreign_key_locale_unique) INT(10)"`
	Locale     string    `xorm:"not null unique(translations_table_name_column_name_foreign_key_locale_unique) VARCHAR(191)"`
	Value      string    `xorm:"not null TEXT"`
	CreatedAt  time.Time `xorm:"default 'NULL' TIMESTAMP"`
	UpdatedAt  time.Time `xorm:"default 'NULL' TIMESTAMP"`
}
