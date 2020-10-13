package users_db

import (
	"database/sql"
	"fmt"
	"log"
)

var (
	Client *sql.DB

	username = "root"
	password = "139714"
	host     = "127.0.0.1"
	shema    = "users_db"
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username, password, host, shema,
	)
	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}

	//mysql.SetLogger(logger.GetLogger())
	log.Println("database successfuly configured")
}
