package mysql

import (
	"database/sql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Dialector(db *sql.DB) gorm.Dialector {
	return mysql.New(mysql.Config{Conn: db})
}
