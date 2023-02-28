package database

import (
	"gorm.io/jinzhu/gorm"
	_ "gorm.io/jinzhu/gorm/dialets/sqlite"
)

var(
	DBconn *gorm.DB
)