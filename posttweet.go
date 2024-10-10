package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"bufio"

	"github.com/dghubble/oauth1"
)

// Response structure for Twitter API v2
type TwitterResponse struct {
	Data struct {
		ID string `json:"id"`
	} `json:"data"`
}

// Error structure for Twitter API
type TwitterError struct {
	Errors []struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"errors"`
}

func main() {
	// Get OAuth credentials from environment variables
	consumerKey := os.Getenv("TWITTER_API_KEY")
	consumerSecret := os.Getenv("TWITTER_API_SECRET")
	accessToken := os.Getenv("TWITTER_ACCESS_TOKEN")
	accessSecret := os.Getenv("TWITTER_ACCESS_SECRET")

	// Check if any credentials are missing
	if consumerKey == "" || consumerSecret == "" || accessToken == "" || accessSecret == "" {
		fmt.Println("Error: Twitter API credentials are not set in environment variables.")
		return
	}

	// Prompt for tweet text
	var tweetText string
	fmt.Print("Enter your tweet: ")
	reader := bufio.NewReader(os.Stdin)
	tweetText, err := reader.ReadString('\n') // Read the full line, including spaces
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Trim any newline characters from the input
	tweetText = tweetText[:len(tweetText)-1] // Remove the trailing newline character

	// Validate tweet length
	if len(tweetText) > 280 {
		fmt.Println("Error: Tweet exceeds the 280-character limit.")
		return
	}

	// OAuth1 config
	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	// Prepare tweet payload
	payload := map[string]string{"text": tweetText}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	// POST request to create a new tweet
	req, err := http.NewRequest("POST", "https://api.twitter.com/2/tweets", bytes.NewBuffer(payloadBytes))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	// Perform POST request
	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Println("Error posting tweet:", err)
		return
	}
	defer resp.Body.Close()

	// Handle response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	if resp.StatusCode != http.StatusCreated {
		var twitterError TwitterError
		if json.Unmarshal(body, &twitterError) == nil {
			for _, errDetail := range twitterError.Errors {
				fmt.Printf("Error Code: %d, Message: %s\n", errDetail.Code, errDetail.Message)
			}
		} else {
			fmt.Println("Error response from Twitter:", resp.Status, string(body))
		}
		return
	}

	// Parse response to get the tweet ID
	var twitterResponse TwitterResponse
	if err := json.Unmarshal(body, &twitterResponse); err != nil {
		fmt.Println("Error unmarshalling response:", err)
		return
	}

	// Print the Tweet ID
	fmt.Printf("Tweet posted successfully! Tweet ID: %s\n", twitterResponse.Data.ID)
}
