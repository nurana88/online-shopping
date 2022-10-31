package command

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type BaseCommand struct {
	ctx context.Context
}

func NewBaseCommand(ctx context.Context) BaseCommand {
	return BaseCommand{
		ctx: ctx,
	}
}

func (cmd *BaseCommand) newDBConnection() (*sql.DB, error) {
	dataSourceName := "root:J>uyQ4aL@tcp(127.0.0.1:3306)/testdb"

	var err error
	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Println("Couldn't connect to DB", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Can not be connected in Ping", err)
		return nil, err
	}
	fmt.Println("Successfully connected to DB...")

	return db, nil
}
