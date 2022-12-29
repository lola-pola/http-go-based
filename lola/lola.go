package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-redis/redis"
)

func main() {
	// Connect to Redis
	client := redis.NewClient(&redis.Options{
		Addr:     "static-redis:6379",
		Password: "",
		DB:       0,
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Measure the latency time
		start := time.Now()

		// Write the response
		fmt.Fprint(w, "Hello, World!")

		// Store the response in Redis
		err := client.Set("response", "Hello, World!", 0).Err()
		if err != nil {
			fmt.Println(err)
			return
		}

		// Calculate and print the latency time
		latency := time.Since(start)
		fmt.Printf("Latency time: %s\n", latency)
	})

	http.ListenAndServe(":8080", nil)
}