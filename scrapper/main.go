package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// Function to fetch and parse the webpage
func fetchPage(url string) (string, string, error) {
	// Request the HTML page.
	res, err := http.Get(url)
	if err != nil {
		return "", "", fmt.Errorf("failed to fetch page: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "", "", fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", "", fmt.Errorf("failed to parse page: %v", err)
	}

	// Extract the text information from <h2 id="description">
	text := doc.Find("h2#description").Text()

	// Extract the actual interest rate from the table
	actualRate := doc.Find("div#ctl00_ContentPlaceHolder1_ctl00_ctl02_Panel1 .table-responsive table tbody tr td").Eq(1).Text()

	return text, actualRate, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	url := "https://tradingeconomics.com/united-states/interest-rate" // Replace with your target URL
	text, rate, err := fetchPage(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the results to the response
	fmt.Fprintf(w, "Description: %s\n", text)
	fmt.Fprintf(w, "Actual Interest Rate: %s\n", rate)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
