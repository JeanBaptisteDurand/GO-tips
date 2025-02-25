package main

import (
	"fmt"
	"time"
)

// Server struct with channels for messages and shutdown signal
type Server struct {
	quitCh chan struct{}  // Empty struct channel for quitting
	msgCh  chan string    // String channel for messages
}

// Constructor function to create a new Server
func newServer() *Server {
	return &Server{
		quitCh: make(chan struct{}),  // Quit channel
		msgCh:  make(chan string, 128), // Buffered channel for messages
	}
}

// Start function to begin server loop
func (s *Server) start() {
	fmt.Println("server starting")
	s.loop() // Start server loop (blocks execution)
}

// Send message to the server
func (s *Server) sendMessage(msg string) {
	s.msgCh <- msg
}

// Quit the server
func (s *Server) quit() {
	s.quitCh <- struct{}{}
}

// Server loop that listens for messages and quit signals
func (s *Server) loop() {
	for {
		select {
		case <-s.quitCh: // If quit signal is received
			fmt.Println("quitting server")
			break // Break out of loop

		case msg := <-s.msgCh: // If a message is received
			s.handleMessage(msg)
		}
	}
	fmt.Println("server is shutting down gracefully")
}

// Handle received messages
func (s *Server) handleMessage(msg string) {
	fmt.Println("We received a message:", msg)
}

// Main function
func main() {
	server := newServer()

	// Simulate server shutdown after 5 seconds
	go func() {
		time.Sleep(time.Second * 5)
		server.quit()
	}()

	server.start()
}
