package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type ApiResponse struct {
	Message string `json:"message"`
}

func checkUsername(username string) (string, error) {
	url := "https://api.solo.to/" + username
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("accept", "application/json")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/133.0.0.0 Safari/537.36")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var apiResponse ApiResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return "", err
	}

	switch apiResponse.Message {
	case "page reserved or blocked":
		return "Username is reserved or blocked", nil
	case "page not found":
		return "Username is available", nil
	default:
		return "Username is taken", nil
	}
}

func main() {
	file, err := os.Open("names.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	availableUsernames := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		username := strings.TrimSpace(scanner.Text())
		if username != "" {
			status, err := checkUsername(username)
			if err != nil {
				fmt.Printf("Error checking username %s: %v\n", username, err)
				continue
			}
			fmt.Printf("Username: %s - Status: %s\n", username, status)

			if status == "Username is available" {
				availableUsernames = append(availableUsernames, username)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	if len(availableUsernames) > 0 {
		outputFile, err := os.Create("available_usernames.txt")
		if err != nil {
			fmt.Println("Error creating output file:", err)
			return
		}
		defer outputFile.Close()

		for _, username := range availableUsernames {
			_, err := outputFile.WriteString(username + "\n")
			if err != nil {
				fmt.Println("Error writing to output file:", err)
				return
			}
		}
		fmt.Println("Available usernames saved to available_usernames.txt")
	} else {
		fmt.Println("No available usernames found.")
	}
}
