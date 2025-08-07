package api
package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

// Configuration struct to hold our API keys
type Config struct {
	NewsAPIKey   string
	OpenAIAPIKey string
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

	return &Config{
		NewsAPIKey:   newsAPIKey,
		OpenAIAPIKey: openAIAPIKey,
	}, nil
}

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

// CORS helper
func setCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}

// Fetch news from NewsAPI
func fetchNews(endpoint string, config *Config) (*NewsResponse, error) {
	url := fmt.Sprintf("https://newsapi.org/v2%s&apiKey=%s", endpoint, config.NewsAPIKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch news: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("NewsAPI returned status %d", resp.StatusCode)
	}

	var newsResponse NewsResponse
	if err := json.Unmarshal(body, &newsResponse); err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %v", err)
	}

	return &newsResponse, nil
}

// Transform news using OpenAI
func transformContent(title, description string, config *Config) (map[string]string, error) {
	systemPrompt := "You are the Ministry of Truth from George Orwell's 1984. Transform news headlines and descriptions into dystopian propaganda using doublespeak, references to Big Brother, the Party, thoughtcrime, etc. Keep responses under 200 characters."

	openAIRequest := OpenAIRequest{
		Model: "gpt-3.5-turbo",
		Messages: []Message{
			{Role: "system", Content: systemPrompt},
			{Role: "user", Content: fmt.Sprintf("Transform this news: Title: %s, Description: %s", title, description)},
		},
		MaxTokens:   200,
		Temperature: 0.9,
	}

	jsonData, err := json.Marshal(openAIRequest)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", strings.NewReader(string(jsonData)))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", config.OpenAIAPIKey))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("OpenAI API returned status %d", resp.StatusCode)
	}

	var openAIResponse OpenAIResponse
	if err := json.NewDecoder(resp.Body).Decode(&openAIResponse); err != nil {
		return nil, err
	}

	if len(openAIResponse.Choices) == 0 {
		return nil, fmt.Errorf("no response from OpenAI")
	}

	return map[string]string{
		"transformedContent": openAIResponse.Choices[0].Message.Content,
	}, nil
}

// Main serverless function handler
func Handler(w http.ResponseWriter, r *http.Request) {
	setCORS(w)

	// Handle preflight requests
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Load configuration
	config, err := loadConfig()
	if err != nil {
		log.Printf("Config error: %v", err)
		http.Error(w, "Server configuration error", http.StatusInternalServerError)
		return
	}

	path := r.URL.Path
	log.Printf("Request: %s %s", r.Method, path)

	// Route handling
	switch {
	case path == "/api/health":
		handleHealth(w, r)
	case strings.HasPrefix(path, "/api/news/headlines"):
		handleHeadlines(w, r, config)
	case strings.HasPrefix(path, "/api/news/search"):
		handleSearch(w, r, config)
	case path == "/api/transform" && r.Method == "POST":
		handleTransform(w, r, config)
	default:
		http.Error(w, "Not found", http.StatusNotFound)
	}
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{
		"status":  "healthy",
		"service": "Ministry of Truth Backend",
		"time":    time.Now().Format(time.RFC3339),
	}
	json.NewEncoder(w).Encode(response)
}

func handleHeadlines(w http.ResponseWriter, r *http.Request, config *Config) {
	w.Header().Set("Content-Type", "application/json")

	category := r.URL.Query().Get("category")
	var endpoint string

	if category != "" {
		endpoint = fmt.Sprintf("/top-headlines?country=us&category=%s", category)
	} else {
		endpoint = "/top-headlines?country=us"
	}

	newsResponse, err := fetchNews(endpoint, config)
	if err != nil {
		log.Printf("Error fetching news: %v", err)
		http.Error(w, fmt.Sprintf("Error fetching news: %v", err), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(newsResponse)
}

func handleSearch(w http.ResponseWriter, r *http.Request, config *Config) {
	w.Header().Set("Content-Type", "application/json")

	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "Query parameter 'q' is required", http.StatusBadRequest)
		return
	}

	endpoint := fmt.Sprintf("/everything?q=%s", query)
	newsResponse, err := fetchNews(endpoint, config)
	if err != nil {
		log.Printf("Error searching news: %v", err)
		http.Error(w, fmt.Sprintf("Error searching news: %v", err), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(newsResponse)
}

func handleTransform(w http.ResponseWriter, r *http.Request, config *Config) {
	w.Header().Set("Content-Type", "application/json")

	var requestData struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	result, err := transformContent(requestData.Title, requestData.Description, config)
	if err != nil {
		log.Printf("Transform error: %v", err)
		http.Error(w, "Error transforming content", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}