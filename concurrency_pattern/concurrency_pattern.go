package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

// UserProfile struct to hold user data
type UserProfile struct {
	ID       int
	Comments []string
	Likes    int
	Friends  []int
}

// Response struct to send data through channels
type Response struct {
	data any
	err  error
}

func main() {
	start := time.Now()

	// Fetch user profile concurrently
	userProfile, err := handleGetUserProfile(10)
	if err != nil {
		log.Fatal(err)
	}

	// Print results
	fmt.Println(userProfile)
	fmt.Println("Fetching the user profile took", time.Since(start))
}

// handleGetUserProfile fetches Likes, Friends, and Comments concurrently
func handleGetUserProfile(id int) (*UserProfile, error) {
	respch := make(chan Response, 3) // Buffered channel to hold responses
	wg := &sync.WaitGroup{}

	// Launch 3 concurrent goroutines
	wg.Add(3)
	go getComments(id, respch, wg)
	go getLikes(id, respch, wg)
	go getFriends(id, respch, wg)

	wg.Wait()  // Wait for all goroutines to complete
	close(respch) // Close the channel after receiving all responses

	// Process responses
	userProfile := &UserProfile{ID: id}
	for resp := range respch {
		if resp.err != nil {
			return nil, resp.err
		}

		// Type assertion to populate UserProfile fields
		switch msg := resp.data.(type) {
		case int:
			userProfile.Likes = msg
		case []int:
			userProfile.Friends = msg
		case []string:
			userProfile.Comments = msg
		}
	}

	return userProfile, nil
}

// Fetches Likes (Simulating an API call)
func getLikes(id int, respch chan Response, wg *sync.WaitGroup) {
	defer wg.Done() // Mark this task as done
	time.Sleep(time.Millisecond * 200) // Simulated delay
	respch <- Response{data: 11, err: nil}
}

// Fetches Friends (Simulating an API call)
func getFriends(id int, respch chan Response, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Millisecond * 100)
	friendIDs := []int{11, 34, 854, 455}
	respch <- Response{data: friendIDs, err: nil}
}

// Fetches Comments (Simulating an API call)
func getComments(id int, respch chan Response, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Millisecond * 200)
	comments := []string{
		"Hey, that was great",
		"Hey Buddy",
		"Ow, I didn’t know that",
	}
	respch <- Response{data: comments, err: nil}
}
