package config

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "log"
)

var DB *sql.DB

func InitDB() {
    var err error
    dsn := "root:@tcp(127.0.0.1:3306)/lpkn-api_asset?charset=utf8mb4&parseTime=True&loc=Local"
    DB, err = sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal(err)
    }
    if err = DB.Ping(); err != nil {
        log.Fatal(err)
    }
    log.Println("Database connecteds")
}
