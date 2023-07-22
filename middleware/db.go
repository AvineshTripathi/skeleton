package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)



type Database struct {
	Client *sql.DB 
}

func NewDb(host, port, user, pwd, dbName string) *Database {

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require", host, port, user, pwd, dbName)
	client , err  := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return &Database{
		Client: client,
	} 
}

func (db *Database) Read() {}

func (db *Database) Write() {}

func (db *Database) Update() {}