# 🗺️ Geoprompt AI (LLM Location Finder with Google Maps Integration)

This project is a simple web-based location assistant that uses a **local LLM via Ollama (`llama3`)** to process natural language queries, extract latitude/longitude from Google Maps links, and display the location on an **embedded Google Map**.

---

## Features

- Ask location-related questions like:  
  _“Where can I find good Italian restaurants near Magelang?”_

- Uses **Ollama (llama3)** to extract lat/lng from Google Maps URLs
- Parses the JSON response and displays results:
  - As an embedded **Google Map**
  - As a clickable link to open in Google Maps
- API key stored securely in `.env`
- Includes **loading indicator** during LLM execution

---
## Requirements

- Go 1.18+
- [Ollama](https://ollama.com) installed and `llama3` model pulled
- Google Maps API key (Embed API enabled)

---

## Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/<your-username>/geoprompt_ai
cd llm-location-finder
```
### 2. Pull the llama3 model
```bash
ollama pull llama3
```

### 3. Create .env file (Rename the .env.example --> .env)
Change the `GOOGLE_MAPS_API_KEY`
```bash
GOOGLE_MAPS_API_KEY=your_google_maps_api_key
```

### 4. Run the Go backend or Build Project
```bash
go run main.go
```

OR

```bash
go build
./geoprompt_ai
```

## Prompt Template (used with LLM)
```bash
You are an API. Return only a JSON array of latitude and longitude coordinates extracted from Google Maps URLs. Do not explain anything.

Input:
Where can I find good Italian restaurants near Magelang, 56481? Return Google Maps links, extract lat/lng, and return only this JSON format:
[{"latitude": <number>, "longitude": <number>}]
```
