package main

import (
	"os"	
)

func main() {

	// Connect the DB
	db := ConnectDb()

	// Server setup 
	server := NewHandler(db)
	server.Run("8080")
}


func ConnectDb() *Database {
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")
	pwd := os.Getenv("DATABASE_PASSWORD")
	dbName := os.Getenv("DATABASE_DBNAME")
	user := os.Getenv("DATABASE_USER")

	db := NewDb(host, port, user, pwd, dbName)

	db.Client.Ping()

	return db
}
