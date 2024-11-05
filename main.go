package main

import (
  "encoding/json"
  "fmt"
	"io/ioutil"
  "log"
  "net/http"
)
//APIKEY: 11fb2fb3ff8c69d16e35c7450ec4cd62
type Message struct {
  Message string `json:"message"`
}

type ExternalAPIResponse struct {
	// Define the fields to match the expected response structure from the external API
	Message string `json:"message"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
  response := Message{Message: "Hello, world!"}
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

  // Unmarshal JSON if needed (optional)
  var externalData ExternalAPIResponse
  err = json.Unmarshal(body, &externalData)
  if err != nil {
    http.Error(w, "Failed to parse response", http.StatusInternalServerError)
    return
  }

  // Set response headers if needed and write the response body
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  w.Write(body) // Directly writes the response body
}

//func main2() {
//  http.HandleFunc("/api/hello", helloHandler)
//  log.Println("Server started on http://localhost:8080")
//  log.Fatal(http.ListenAndServe(":8080", nil))
//}

func main() {
  http.HandleFunc("/api/hello", helloHandler)
	http.HandleFunc("/api/weather", weatherHandler)
	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}