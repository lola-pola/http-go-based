package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/go-redis/redis"
)

func main() {
	// Create a Redis client
	client := redis.NewClient(&redis.Options{
		Addr:     "static-redis:6379",
		Password: "",
		DB:       0,
	})

	// Make a GET request to the REST API
	startTime := time.Now()
	res, err := http.Get("https://example.com/api/endpoint")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Store the response in Redis
	err = client.Set("key", body, 0).Err()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Measure the latency
	elapsedTime := time.Since(startTime)
	fmt.Printf("Latency: %s\n", elapsedTime)
}



