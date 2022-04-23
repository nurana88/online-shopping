package infra

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func init() {
	dataSourceName := "root:pwd@tcp(127.0.0.1:3306)/testdb"

	var err error
	DB, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Println("Can not connect", err)
		return
	}

	defer DB.Close()

	fmt.Println("Successfully connected to DB...")

}
