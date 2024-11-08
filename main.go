package main

import (
  "encoding/json"
  "fmt"
	"io/ioutil"
  "log"
  "net/http"
)
//APIKEY: 11fb2fb3ff8c69d16e35c7450ec4cd62
type ApiResponse struct {
	// Define the fields to match the expected response structure from the external API
	Message string `json:"message"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
  response := ApiResponse{Message: "Hello, world!! Welcome to my API :)"}
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(response)
}


func weatherHandler(w http.ResponseWriter, r *http.Request) {
  // Call the external API
  resp, err := http.Get("https://api.openweathermap.org/data/2.5/weather?lat=40.599&lon=22.951&appid=11fb2fb3ff8c69d16e35c7450ec4cd62")
  if err != nil {
    http.Error(w, "Failed to reach external API", http.StatusInternalServerError)
    return
  }
  defer resp.Body.Close()

  // Read the response body
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    http.Error(w, "Failed to read response", http.StatusInternalServerError)
    return
  }

  // Set response headers if needed and write the response body
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  w.Write(body) // Directly writes the response body
}

func forecastHandler(w http.ResponseWriter, r *http.Request) {
  // Call the external API
  resp, err := http.Get("https://api.openweathermap.org/data/2.5/forecast?lat=40.599&lon=22.951&appid=11fb2fb3ff8c69d16e35c7450ec4cd62")
  if err != nil {
    http.Error(w, "Failed to reach external API", http.StatusInternalServerError)
    return
  }
  defer resp.Body.Close()

  // Read the response body
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    http.Error(w, "Failed to read response", http.StatusInternalServerError)
    return
  }

  // Set response headers if needed and write the response body
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  w.Write(body) // Directly writes the response body
}

func main() {
  http.HandleFunc("/api/hello", helloHandler)
	http.HandleFunc("/api/weather", weatherHandler)
  http.HandleFunc("/api/forecast", forecastHandler)
	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}