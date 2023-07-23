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

type User struct {
    ID    int
    Name  string
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

func (db *Database) CreateTable() error {
	 _, err := db.Client.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id SERIAL PRIMARY KEY,
            name VARCHAR(50)
        )
    `)
    return err
}

func (db *Database) Read(userID int) (*User, error) {
	user := &User{}
	err := db.Client.QueryRow("SELECT id, name FROM users WHERE id = $1", userID).
        Scan(&user.ID, &user.Name)
    return user, err
}

func (db *Database) Insert(User *User) error {
	err := db.Client.QueryRow("INSERT INTO users (name) VALUES ($1)", User.Name).Scan()
	if err != nil {
		return err
	}
	return nil 
}

func (db *Database) Update(User *User) {}