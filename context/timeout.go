package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()
	ctx := context.WithValue(context.Background(), "username", "jedurand")

	userID, err := fetchUserID(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The response took %v -> %v\n", time.Since(start), userID)
}

func fetchUserID(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*100)
	defer cancel() // Ensures the context is canceled when the function returns

	val := ctx.Value("username")
	fmt.Println("The value =", val)

	// Result struct to hold response from the third-party call
	type result struct {
		userId string
		err    error
	}

	resultCh := make(chan result, 1)

	// Run third-party HTTP call in a goroutine
	go func() {
		userId, err := thirdpartyHTTPCall()
		resultCh <- result{
			userId: userId,
			err:    err,
		}
	}()

	// Use select to wait for either the response or context timeout
	select {
	case <-ctx.Done(): // Case when context deadline exceeds or is manually canceled
		return "", ctx.Err()
	case res := <-resultCh: // Case when the HTTP call returns a result
		return res.userId, res.err
	}
}

// Simulated third-party HTTP call that takes 500ms
func thirdpartyHTTPCall() (string, error) {
	time.Sleep(time.Millisecond * 500)
	return "user id 1", nil
}
