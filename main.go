package main

import (
	"fmt"
	"net/http"
	"time"
)

var count int64
var t0 time.Time

// CounterHandler shows such statistics number of visits and the time since server start
func CounterHandler(w http.ResponseWriter, r *http.Request) {
	count++
	t1 := time.Now()
	fmt.Fprintf(w, "Time since server start %v.\n", t1.Sub(t0))
	fmt.Fprintf(w, "Visits: %d\n", count)
}

// NullHandler for requests ignoring
func NullHandler(w http.ResponseWriter, r *http.Request) {
}

func main() {
	t0 = time.Now()
	http.HandleFunc("/", CounterHandler)
	http.HandleFunc("/favicon.ico", NullHandler)
	http.ListenAndServe(":8080", nil)
}
