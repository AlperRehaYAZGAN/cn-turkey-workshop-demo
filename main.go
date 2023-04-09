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
	appPort := os.Getenv("APP_PORT")
	if redisURL == "" || redisPassword == "" || appPort == "" {
		fmt.Println("REDIS_URL or REDIS_PASSWORD or APP_PORT is not set")
	}
	fmt.Printf("App is starting with redis url: %s and password: %s and http port on %s", redisURL, redisPassword, appPort)

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
		fmt.Printf("error: unable to ping redis with initial connection: %v", err)
	}

	// create counter
	err = rdb.Set(rContext, "counter", 0, 0).Err()
	if err != nil {
		fmt.Printf("error: unable to set counter to 0: %v", err)
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

	// start server
	fmt.Println("Server is starting on port " + appPort)
	http.ListenAndServe(":"+appPort, nil)
}
