package main

import (
	"bytes"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)

type PromptRequest struct {
	Query string `json:"query"`
}

type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func main() {
	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	http.HandleFunc("/", serveIndex)
	http.HandleFunc("/ask", handleAsk)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Server running on http://localhost:8080")
	log.Println("Google Maps API Key :", os.Getenv("GOOGLE_MAPS_API_KEY"))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	data := struct {
		ApiKey string
	}{
		ApiKey: os.Getenv("GOOGLE_MAPS_API_KEY"),
	}
	tmpl.Execute(w, data)
}

func handleAsk(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	var req PromptRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	basePrompt := `You are an API. Return only a JSON array of latitude and longitude coordinates extracted from Google Maps URLs. Do not explain anything.

Input:
` + req.Query + `

Return only this JSON format:
[{"latitude": <number>, "longitude": <number>}]`

	cmd := exec.Command("ollama", "run", "llama3", basePrompt)
	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		http.Error(w, "Error running Ollama: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var coords []Coordinates
	if err := json.Unmarshal(out.Bytes(), &coords); err != nil {
		http.Error(w, "Failed to parse JSON: "+err.Error()+"\nOutput:\n"+out.String(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(coords)
}
