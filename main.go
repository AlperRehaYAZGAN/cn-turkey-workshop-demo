package main

// simple webservice that returns hello world
import (
	"context"
	"fmt"
	"net/http"
	"os"

	// redis/v8
	"github.com/go-redis/redis/v8"
)

func main() {
	// get redis url and password from env
	redisURL := os.Getenv("REDIS_URL")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	if redisURL == "" || redisPassword == "" {
		fmt.Println("REDIS_URL or REDIS_PASSWORD is not set")
	}

	// connects redis
	rContext := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisURL,
		Password: redisPassword, // no password set
		DB:       0,             // use default DB
	})

	// checks if redis is connected
	_, err := rdb.Ping(rContext).Result()
	if err != nil {
		fmt.Println("Redis is not connected")
	}

	// create counter
	err = rdb.Set(rContext, "counter", 0, 0).Err()
	if err != nil {
		fmt.Println("Counter is not created")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// incrments the counter
		counter, err := rdb.Incr(rContext, "counter").Result()
		if err != nil {
			// return error to user
			fmt.Fprintf(w, "Page viewed %d times because error occured", 0)
		} else {
			// return counter to user
			fmt.Fprintf(w, "Page viewed %d times", counter)
		}
	})

	http.ListenAndServe(":4040", nil)
}
