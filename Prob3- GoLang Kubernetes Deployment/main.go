package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// handler serves the current date & time with styled HTML response
func handler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format("Monday, Jan 2, 2006 - 15:04:05 MST")
	response := fmt.Sprintf(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Current Time</title>
			<style>
				body { font-family: Arial, sans-serif; text-align: center; padding: 50px; background-color: #f4f4f4; }
				.container { background: #fff; padding: 20px; border-radius: 10px; box-shadow: 0 4px 8px rgba(0,0,0,0.2); display: inline-block; }
				h1 { color: #333; }
				p { font-size: 1.5em; color: #007bff; margin: 0; }
			</style>
		</head>
		<body>
			<div class="container">
				<h1>Current Date & Time</h1>
				<p>%s</p>
			</div>
		</body>
		</html>`, currentTime)

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, response)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// Graceful shutdown
	go func() {
		fmt.Println("🚀 Server running on http://localhost:8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("❌ Server error: %v\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	fmt.Println("\n🛑 Shutting down server...")
	server.Close()
	fmt.Println("✅ Server gracefully stopped")
}
