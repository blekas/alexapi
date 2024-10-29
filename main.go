package main

import (
  "encoding/json"
  "log"
  "net/http"
)

type Message struct {
  Message string `json:"message"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
  response := Message{Message: "Hello, world!"}
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(response)
}

func main() {
  http.HandleFunc("/api/hello", helloHandler)
  log.Println("Server started on http://localhost:8080")
  log.Fatal(http.ListenAndServe(":8080", nil))
}