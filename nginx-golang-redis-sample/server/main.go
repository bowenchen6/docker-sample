package main

import (
    "net/http"
    "log"
    "fmt"
    "github.com/gorilla/mux"
	"github.com/garyburd/redigo/redis"
)

var redisCon redis.Conn

func MyHandler(w http.ResponseWriter, r *http.Request) {
    counter, err := redisCon.Do("GET", "counter")
    if err != nil {
        fmt.Fprintf(w, "Could not connect to Redis with error: %s\n", err)
    }
    if counter == nil {
        counter = "0"
        redisCon.Do("set", "counter", counter)
    }
    redisCon.Do("incr", "counter")
    fmt.Fprintf(w, "This page has been viewed %s times!\n", counter)
}

func main() {
    var err error
    r := mux.NewRouter()
    redisCon, err = redis.Dial("tcp", "redis:6379")
    if err != nil {
		log.Fatalf("Could not connect to Redis with error: %s", err)
	}
	defer redisCon.Close()
    r.HandleFunc("/", MyHandler)
    log.Fatal(http.ListenAndServe(":8080", r))
}