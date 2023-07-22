package main

import (
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	dbHandler *Database
}

func NewHandler(db *Database) *Server {
	return &Server{
		dbHandler: db,
	} 
}

func (sv *Server) Run(port string) {
	fmt.Println("Server is running on port:", port)
	
	// endpoints
	http.HandleFunc("/", sv.handleRequest)

	addr := fmt.Sprintf(":%s", port)
	log.Fatalf("Error starting the server: %v", http.ListenAndServeTLS(addr, "", "", nil))
}

func (sv *Server) handleRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte("works")); err != nil {
		fmt.Printf("Something is wrong: %v ", err.Error())
		http.Error(w, "Could not connect, server is down", http.StatusInternalServerError)
	}
}
