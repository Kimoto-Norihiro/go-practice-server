package database

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

func NewMySql(dns string) (db *gorm.DB, err error) {
	db, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}