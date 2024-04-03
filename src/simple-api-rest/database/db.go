package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func OpenDB() {
	dsn := "root:root@tcp(127.0.0.1:3306)/db"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panic("Erro de conexao com o banco")
	}
}
