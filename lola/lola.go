package main

import (
	"fmt"
	"net/http"
	"time"
	// "math/rand"
	"strconv"
	"os"
	"github.com/go-redis/redis"
)

func main() {
	

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "GET" {
		// Parse the query string parameters
		params := r.URL.Query()

		
		validate := params.Get("validate")

		hostname, erroo := os.Hostname()
		if erroo != nil {
			fmt.Println(erroo)
		}

		if validate == "stamp" {
			start := time.Now()		
			latency := time.Since(start)
			fmt.Println(latency)
			return
		}

		if validate == "true" {
			start := time.Now()		
			fmt.Fprintf(w, "Validated, %s!", validate , hostname )
			client := &http.Client{}
			req, err := http.NewRequest("GET", "http://from-github-http-servers", nil)
			res, err := client.Do(req)
			if err != nil {
				fmt.Println(err)
				return
			}else{
				fmt.Println(res.StatusCode)
			}
			latency := time.Since(start)
			fmt.Fprintf(w, "time inside, %s!", latency ,hostname )
			return
		}


		if validate == "google" {
			start := time.Now()		
			fmt.Fprintf(w, "google, %s!", validate )
			client := &http.Client{}
			req, err := http.NewRequest("GET", "https://google.com", nil)
			res, err := client.Do(req)
			if err != nil {
				fmt.Println(err)
				return
			}else{
				fmt.Println(res.StatusCode)
			}
			latency := time.Since(start)
			fmt.Fprintf(w, "time outside, %s!", latency )
			return
		}


		// Get the value of the "name" parameter
		numStr := params.Get("num")

		// Write the response
		fmt.Fprintf(w, "Number to calculate, %s!", numStr )
		
		// Get the hostname 


		redisname := params.Get("redis") 
		// Write the response
		fmt.Fprintf(w, "REDIS, %s!", redisname )
		
		
		if redisname != "" {
			fmt.Fprintf(w, "REDIS, %s!", redisname )
		}else{
			redisname = "from-github-redis-0:6379"
		}

		// Connect to Redis
		client := redis.NewClient(&redis.Options{
			
			Addr:     redisname,	
			PoolSize: 100000,
			// Addr:     "127.0.0.1:26379",
			Password: "",
			DB:       0,
			
		})


		// Measure the latency time
		
		startingtime := time.Now()
		// Write the response
		fmt.Fprint(w, " Start Calculating, %s!", startingtime, hostname)

		num, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Println(err)
			return
		}
		// Calculate the Fibonacci number
		 
		// go fibonacci(num)
		fmt.Println(num)

		fmt.Println(fibonacci(1))
		start := time.Now()		
		errs := client.Set(startingtime.String() ,start , 0).Err()
		if errs != nil {
			fmt.Println(errs)
			return
		}

		// Calculate and print the latency time
		latency := time.Since(start)
		fmt.Fprint(w ,latency)
		fmt.Fprint(w, " Latency time, %s!", latency, hostname)
		fmt.Printf("Latency time: %s\n", latency, hostname)
		}
	})
	
	http.ListenAndServe(":8085", nil)
}




func fibonacci(n int) int {
	
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}





