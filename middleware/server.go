package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
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
	cert, err := tls.LoadX509KeyPair("./certs/go_server.crt", "./certs/go_server.key")
	if err != nil {
		panic("Error loading certificates: " + err.Error())
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	addr := fmt.Sprintf(":%s", port)
	ser := &http.Server{
		Addr: addr,
		Handler: nil,
		TLSConfig: tlsConfig,
	}

	
	// endpoints
	http.HandleFunc("/", sv.handlePing)
	http.HandleFunc("/users/id?", sv.handleRead)
	http.HandleFunc("/user", sv.handleInsert)
	
	log.Fatalf("Error starting the server: %v", ser.ListenAndServeTLS("",""))
	fmt.Println("Server is running on port:", port)
}

func (sv *Server) handlePing(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte("works")); err != nil {
		fmt.Printf("Something is wrong: %v ", err.Error())
		http.Error(w, "Could not connect, server is down", http.StatusInternalServerError)
	}
}

func (sv *Server) handleInsert(w http.ResponseWriter, r *http.Request) {
	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid reques payload", http.StatusBadRequest)
	}

	if err = sv.dbHandler.Insert(&newUser); err != nil {
		http.Error(w, "Couldn't insert user", http.StatusExpectationFailed)
	}

	w.WriteHeader(http.StatusAccepted)
	resp := fmt.Sprintf("User Added: %s", newUser.Name)
	if _, err := w.Write([]byte(resp)); err != nil {
		fmt.Printf("Something is wrong: %v ", err.Error())
		http.Error(w, "Could not connect, server is down", http.StatusInternalServerError)
	}

}

func (sv *Server) handleRead(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid parameter", http.StatusBadRequest)
	}

	user, err := sv.dbHandler.Read(id)
	if err != nil {
		http.Error(w, "Couldn't read user", http.StatusNoContent)
	}

	json.NewEncoder(w).Encode(user)
}

func (sv *Server) handleUpdate(w http.ResponseWriter, r *http.Request) {
	
}