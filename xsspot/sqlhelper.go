package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const (
	dbDriverName = "sqlite3"
	dbName       = "../xssdata/data.db3"
)

type cross struct {
	Ip string
	Port int
	Url string
	Header string
	Body string
}

func createTable(db *sql.DB) error {
	sql := `create table if not exists "cross" (
		"id" integer primary key autoincrement,
		"ip" text not null,
		"port" integer not null,
		"url" text,
		"header" text,
		"body" text,
		'time' DATETIME NOT NULL DEFAULT (strftime('%Y-%m-%d %H:%M:%S', 'now', 'localtime'))
	)`
	_, err := db.Exec(sql)
	return err
}

func insertData(db *sql.DB, c cross) error {
	sql := `insert into cross (ip, port, url, header, body) values (?,?,?,?,?)`
	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(c.Ip, c.Port, c.Url, c.Header, c.Body)
	return err
}

func checkErr(e error) bool {
	if e != nil {
		log.Fatal(e)
		return true
	}
	return false
}
