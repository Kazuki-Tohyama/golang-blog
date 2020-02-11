package database

import (
	"github.com/jinzhu/gorm"
)

type SqlHandler interface {
	Find(interface{}, ...interface{}) (*gorm.DB)
	Exec(string, ...interface{}) (*gorm.DB)
	First(interface{}, ...interface{}) (*gorm.DB)
	Raw(string, ...interface{}) (*gorm.DB)
	Create(...interface{}) (*gorm.DB)
	Save(...interface{}) (*gorm.DB)
	Delete(...interface{}) (*gorm.DB)
	Where(interface{}, ...interface{}) (*gorm.DB)
}

type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

type Row interface {
	Scan(...interface{}) error
	Next() bool
	Close() error
}
