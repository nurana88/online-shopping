package Infrastructure

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func init() {
	dataSourceName := "root:J>uyQ4aL@tcp(127.0.0.1:3306)/testdb"

	var err error
	DB, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Println("Couldn't connect to DB", err)
		return
	}

	err = DB.Ping()
	if err != nil {
		fmt.Println("Can not be connected in Ping", err)
		return
	}
	fmt.Println("Successfully connected to DB...")

}
