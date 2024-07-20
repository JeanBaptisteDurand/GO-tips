package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// Function to fetch and parse the webpage
func fetchPage(url string) (string, error) {
	// Request the HTML page.
	res, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to fetch page: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "", fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", fmt.Errorf("failed to parse page: %v", err)
	}

	// Extract the text information from <h2 id="description">
	text := doc.Find("h2#description").Text()

	return text, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	url := "https://tradingeconomics.com/united-states/interest-rate" // Replace with your target URL
	text, err := fetchPage(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the results to the response
	fmt.Fprintf(w, "Text: %s\n", text)
}

func main() {
	url := "https://tradingeconomics.com/united-states/interest-rate" // Replace with your target URL
	text, err := fetchPage(url)
	if err != nil {
		log.Fatalf("Error fetching page: %v", err)
	}

	fmt.Println("Extracted Text:")
	fmt.Println(text)
}
