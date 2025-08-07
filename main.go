package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

// Configuration struct to hold our API keys
type Config struct {
	NewsAPIKey   string
	OpenAIAPIKey string
	Port         string
}

// Load configuration from environment variables
func loadConfig() (*Config, error) {
	newsAPIKey := os.Getenv("NEWS_API_KEY")
	if newsAPIKey == "" {
		return nil, fmt.Errorf("NEWS_API_KEY environment variable is required")
	}

	openAIAPIKey := os.Getenv("OPENAI_API_KEY")
	if openAIAPIKey == "" {
		return nil, fmt.Errorf("OPENAI_API_KEY environment variable is required")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}

	return &Config{
		NewsAPIKey:   newsAPIKey,
		OpenAIAPIKey: openAIAPIKey,
		Port:         port,
	}, nil
}

// Global config variable
var config *Config

// API response structures
type NewsResponse struct {
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults"`
	Articles     []Article `json:"articles"`
}

type Article struct {
	Source      Source `json:"source"`
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	URLToImage  string `json:"urlToImage"`
	PublishedAt string `json:"publishedAt"`
	Content     string `json:"content"`
}

type Source struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type OpenAIRequest struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	MaxTokens   int       `json:"max_tokens"`
	Temperature float64   `json:"temperature"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OpenAIResponse struct {
	Choices []Choice `json:"choices"`
}

type Choice struct {
	Message Message `json:"message"`
}

// CORS middleware for API access
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		log.Printf("%s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}

// Fetch news from NewsAPI using environment variable
func fetchNews(endpoint string) (*NewsResponse, error) {
	url := fmt.Sprintf("https://newsapi.org/v2%s&apiKey=%s", endpoint, config.NewsAPIKey)

	// Log request with masked API key for security
	maskedURL := strings.Replace(url, config.NewsAPIKey, "[REDACTED]", 1)
	log.Printf("Making request to: %s", maskedURL)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch news: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	log.Printf("NewsAPI response status: %d", resp.StatusCode)
	if resp.StatusCode != http.StatusOK {
		log.Printf("NewsAPI error - status: %d", resp.StatusCode)
		return nil, fmt.Errorf("NewsAPI returned status %d", resp.StatusCode)
	}

	var newsResponse NewsResponse
	if err := json.Unmarshal(body, &newsResponse); err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %v", err)
	}

	log.Printf("Successfully parsed %d articles", len(newsResponse.Articles))
	return &newsResponse, nil
}

// Get top headlines endpoint
func getTopHeadlines(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	category := r.URL.Query().Get("category")
	var endpoint string

	if category != "" {
		endpoint = fmt.Sprintf("/top-headlines?country=us&category=%s", category)
	} else {
		endpoint = "/top-headlines?country=us"
	}

	newsResponse, err := fetchNews(endpoint)
	if err != nil {
		log.Printf("Error fetching news: %v", err)
		http.Error(w, fmt.Sprintf("Error fetching news: %v", err), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(newsResponse)
}

// Search news endpoint
func searchNews(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "Query parameter 'q' is required", http.StatusBadRequest)
		return
	}

	endpoint := fmt.Sprintf("/everything?q=%s", query)
	newsResponse, err := fetchNews(endpoint)
	if err != nil {
		log.Printf("Error searching news: %v", err)
		http.Error(w, fmt.Sprintf("Error searching news: %v", err), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(newsResponse)
}

// Transform news using OpenAI API
func transformNews(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var requestData struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	systemPrompt := "You are the Ministry of Truth from George Orwell's 1984. Transform news headlines and descriptions into dystopian propaganda using doublespeak, references to Big Brother, the Party, thoughtcrime, etc. Keep responses under 200 characters."

	openAIRequest := OpenAIRequest{
		Model: "gpt-3.5-turbo",
		Messages: []Message{
			{Role: "system", Content: systemPrompt},
			{Role: "user", Content: fmt.Sprintf("Transform this news: Title: %s, Description: %s", requestData.Title, requestData.Description)},
		},
		MaxTokens:   200,
		Temperature: 0.9,
	}

	jsonData, err := json.Marshal(openAIRequest)
	if err != nil {
		http.Error(w, "Error creating request", http.StatusInternalServerError)
		return
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", strings.NewReader(string(jsonData)))
	if err != nil {
		http.Error(w, "Error creating request", http.StatusInternalServerError)
		return
	}

	// Use environment variable for API key
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", config.OpenAIAPIKey))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Error making request to OpenAI", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		log.Printf("OpenAI API error - status: %d", resp.StatusCode)
		http.Error(w, "Error from OpenAI API", http.StatusInternalServerError)
		return
	}

	var openAIResponse OpenAIResponse
	if err := json.NewDecoder(resp.Body).Decode(&openAIResponse); err != nil {
		http.Error(w, "Error parsing OpenAI response", http.StatusInternalServerError)
		return
	}

	if len(openAIResponse.Choices) == 0 {
		http.Error(w, "No response from OpenAI", http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"transformedContent": openAIResponse.Choices[0].Message.Content,
	}

	json.NewEncoder(w).Encode(response)
}

// Health check endpoint
func healthCheck(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"status":  "healthy",
		"service": "Ministry of Truth Backend",
		"time":    time.Now().Format(time.RFC3339),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	// Load configuration from environment variables
	var err error
	config, err = loadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	log.Printf("Ministry of Truth Backend starting on port %s", config.Port)

	r := mux.NewRouter()

	// Apply CORS middleware to all routes
	r.Use(corsMiddleware)

	// API routes
	r.HandleFunc("/api/news/headlines", getTopHeadlines).Methods("GET")
	r.HandleFunc("/api/news/search", searchNews).Methods("GET")
	r.HandleFunc("/api/transform", transformNews).Methods("POST")
	r.HandleFunc("/health", healthCheck).Methods("GET")

	// Serve static files
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

	log.Printf("Server starting on port %s", config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, r))
}
