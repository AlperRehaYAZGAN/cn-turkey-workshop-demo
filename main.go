package main

// simple webservice that returns hello world
import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	// redis/v8
	"github.com/go-redis/redis/v8"

	// template
	"html/template"
)

func main() {
	// working directory
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalln("error: unable to get current working directory")
	}

	// index.html
	t, err := template.ParseFiles(pwd + "/templates/index.html")
	if err != nil {
		log.Fatalln("error: unable to parse index.html " + err.Error())
	}

	// redis environments
	redisURL := os.Getenv("REDIS_URL")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	appPort := os.Getenv("APP_PORT")
	if redisURL == "" || redisPassword == "" || appPort == "" {
		log.Fatalln("APP_PORT, REDIS_URL or REDIS_PASSWORD is not set")
	}

	// init redis conn
	rContext := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisURL,
		Password: redisPassword,
		DB:       0,
	})

	// http handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// increment counter and return render result of t index.html with "Counter"
		counter, _ := rdb.Incr(rContext, "counter").Result()
		t.Execute(w, map[string]interface{}{
			"Counter": counter,
		})
	})

	// start server
	fmt.Println("Server is starting on port " + appPort)
	http.ListenAndServe(":"+appPort, nil)
}
