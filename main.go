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
	Message string `json:"message"`
}

// Function to fetch data from the external API with dynamic endpoint
func fetchExternalData(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

  // Unmarshal JSON if needed (optional)
  var externalData ApiResponse
  err = json.Unmarshal(body, &externalData)
  if err != nil {
    http.Error(w, "Failed to parse response", http.StatusInternalServerError)
    return nil, err
  }

	return body, nil
}

// Single handler for /api/hello, /weather/now, and /weather/forecast
func genericHandler(w http.ResponseWriter, r *http.Request) {
  switch r.URL.Path {
  case "/api/hello":
		// Handle /api/hello
		response := ApiResponse{Message: "Hello, world!! This is my API :)"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)

  case "/weather/now":
  case "/weather/forecast":
		// Handle /weather/now and /weather/forecast
    if (r.URL.Path == "/weather/forecast") {
      url := "https://api.openweathermap.org/data/2.5/forecast?lat=40.599&lon=22.951&appid=11fb2fb3ff8c69d16e35c7450ec4cd62"
    }
    else {
      url := "https://api.openweathermap.org/data/2.5/weather?lat=40.599&lon=22.951&appid=11fb2fb3ff8c69d16e35c7450ec4cd62"
    }
    body, err := fetchExternalData(url)
		if err != nil {
			http.Error(w, "Failed to reach external API", http.StatusInternalServerError)
			return
		}
  }

  // Set response headers if needed and write the response body
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  w.Write(body) // Directly writes the response body
}

func main() {
  http.HandleFunc("/api/hello", genericHandler)
	http.HandleFunc("/weather/now", genericHandler)
  http.HandleFunc("/weather/forecst", genericHandler)
	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}