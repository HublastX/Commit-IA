package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type LLMRequest struct {
	Prompt string `json:"prompt"`
}

type LLMResponse struct {
	Response string `json:"response"`
	Error    string `json:"error,omitempty"`
}

type GoogleRequest struct {
	Contents         []GoogleContent `json:"contents"`
	GenerationConfig struct {
		Temperature     float64 `json:"temperature"`
		MaxOutputTokens int     `json:"maxOutputTokens"`
	} `json:"generationConfig"`
}

type GoogleContent struct {
	Parts []GooglePart `json:"parts"`
}

type GooglePart struct {
	Text string `json:"text"`
}

type GoogleCandidate struct {
	Content GoogleContent `json:"content"`
}

type GoogleAPIResponse struct {
	Candidates []GoogleCandidate `json:"candidates"`
	Error      *struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
	} `json:"error,omitempty"`
}

func callGoogleLLM(prompt string) (string, error) {
	apiKey := os.Getenv("GOOGLE_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("GOOGLE_API_KEY not set")
	}

	requestBody := GoogleRequest{
		Contents: []GoogleContent{
			{
				Parts: []GooglePart{
					{Text: prompt},
				},
			},
		},
	}
	requestBody.GenerationConfig.Temperature = 0.3
	requestBody.GenerationConfig.MaxOutputTokens = 2048

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %v", err)
	}

	url := fmt.Sprintf("https://generativelanguage.googleapis.com/v1beta/models/gemini-2.0-flash-exp:generateContent?key=%s", apiKey)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	var response GoogleAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", fmt.Errorf("failed to decode response: %v", err)
	}

	if response.Error != nil {
		return "", fmt.Errorf("Google API error: %s", response.Error.Message)
	}

	if len(response.Candidates) == 0 || len(response.Candidates[0].Content.Parts) == 0 {
		return "", fmt.Errorf("no response content returned")
	}

	return response.Candidates[0].Content.Parts[0].Text, nil
}

func llmHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(LLMResponse{Error: "Only POST method allowed"})
		return
	}

	var req LLMRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(LLMResponse{Error: "Invalid JSON format"})
		return
	}

	if req.Prompt == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(LLMResponse{Error: "Prompt field is required"})
		return
	}

	response, err := callGoogleLLM(req.Prompt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(LLMResponse{Error: fmt.Sprintf("Error generating response: %v", err)})
		return
	}

	json.NewEncoder(w).Encode(LLMResponse{Response: response})
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok", "service": "commit-ia-api"})
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/llm", llmHandler).Methods("POST")
	router.HandleFunc("/health", healthHandler).Methods("GET")

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	})

	handler := c.Handler(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("🚀 API running on port %s\n", port)
	fmt.Println("Endpoints:")
	fmt.Println("  POST /llm - Send message to LLM")
	fmt.Println("  GET /health - Health check")

	log.Fatal(http.ListenAndServe(":"+port, handler))
}
