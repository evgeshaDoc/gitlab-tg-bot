package main

import (
	"database/sql"
	"fmt"

	"github.com/evgeshaDoc/gitlab-tg-bot/internal"
	_ "modernc.org/sqlite"
)

func main() {
	config, err := internal.GetConfig()
	if err != nil {
		panic(err)
	}

	db, err := sql.Open("sqlite", config.DbFile)
	if err != nil {
		panic(err)
	}

	res, err := db.Exec("select * from Tokens")
	if err != nil {
		panic(err)
	}
	fmt.Printf("res: %v\n", res)
}
