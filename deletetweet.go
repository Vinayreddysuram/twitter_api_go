package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/dghubble/oauth1"
)

func main() {
	// OAuth1 config with your API keys from environment variables
	apiKey := os.Getenv("TWITTER_API_KEY")
	apiSecret := os.Getenv("TWITTER_API_SECRET")
	accessToken := os.Getenv("TWITTER_ACCESS_TOKEN")
	accessSecret := os.Getenv("TWITTER_ACCESS_SECRET")

	if apiKey == "" || apiSecret == "" || accessToken == "" || accessSecret == "" {
		fmt.Println("Twitter API credentials are not set in the environment variables.")
		return
	}

	config := oauth1.NewConfig(apiKey, apiSecret)
	token := oauth1.NewToken(accessToken, accessSecret)

	// Create an HTTP client with OAuth
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	httpClient := config.Client(ctx, token)

	// Get the Tweet ID from console input
	var tweetID string
	fmt.Print("Enter the Tweet ID you want to delete: ")
	_, err := fmt.Scanln(&tweetID)
	if err != nil {
		fmt.Println("Error reading tweet ID:", err)
		return
	}

	// URL for deleting the tweet
	deleteUrl := fmt.Sprintf("https://api.twitter.com/1.1/statuses/destroy/%s.json", tweetID)
	req, err := http.NewRequest("POST", deleteUrl, nil)
	if err != nil {
		fmt.Println("Error creating delete request:", err)
		return
	}

	// Perform the delete request
	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Println("Error deleting tweet:", err)
		return
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		fmt.Printf("Failed to delete the tweet. Status: %s, Response: %s\n", resp.Status, body)
		return
	}

	fmt.Println("Tweet deleted successfully!")
}
