package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// const (
// 	driver = "mysql"
// 	user   = "root"
// 	pass   = "root"
// 	host   = "localhost"
// 	port   = "3306"
// 	db     = "minhas_autorizacoes"
// )

var (
	DB  *sql.DB
	err error
)

func OpenDB() {
	DB, err = sql.Open("mysql", "root:root@/minhas_autorizacoes?parseTime=true")

	if err != nil {
		panic(err)
	} else {
		fmt.Println("")
		fmt.Println("<! ---------- Conexão com o banco realizada! ---------- !>")
		fmt.Println("")
	}

	showTables()

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
}

func showTables() {
	rows, err := DB.Query("SHOW TABLES;")

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var tableName string
		rows.Scan(&tableName)
		fmt.Println(tableName)
	}
}
