package main

import (
	"fmt"
	"net/http"
	"time"
	// "math/rand"
	"github.com/go-redis/redis"
)


func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	// Connect to Redis
	client := redis.NewClient(&redis.Options{
		Addr:     "static-redis:6379",
		// Addr:     "127.0.0.1:26379",
		Password: "",
		DB:       0,
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Measure the latency time
		
		startingtime := time.Now()
		// Write the response
		fmt.Fprint(w, " Start ")
		

		
		
		

		// rand.Seed(time.Now().UnixNano())
		// b := rand.Intn(100) + 1
		// fmt.Println(b)
		fmt.Println(fibonacci(120))
		// Store the response in Redis
		// latencys := time.Since(start)
		// lolas := latencys.String()
		start := time.Now()
		err := client.Set(startingtime.String() ,start , 0).Err()
		if err != nil {
			fmt.Println(err)
			return
		}

		// Calculate and print the latency time
		latency := time.Since(start)
		fmt.Fprint(w ,latency)
		fmt.Printf("Latency time: %s\n", latency)
	})
	
	http.ListenAndServe(":8085", nil)
}