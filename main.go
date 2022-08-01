package main

import (
	_ "github.com/go-redis/redis"
	"net/http"

	"fmt"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8899", nil)
}
