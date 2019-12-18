package main

// Requests

// /messages - return all messages (method GET)
// /msg - add new message (method POST - data = { msg: string })

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Message struct
type Message struct {
	Msg string `json:"msg"`
	ID  string `json:"id"`
}

var messages []Message

func getMessages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}

func sendMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var message Message
	_ = json.NewDecoder(r.Body).Decode(&message)
	message.ID = strconv.Itoa(rand.Intn(1000000))
	messages = append(messages, message)
	json.NewEncoder(w).Encode(message)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/messages", getMessages).Methods("GET")
	r.HandleFunc("/msg", sendMessage).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", r))
}
