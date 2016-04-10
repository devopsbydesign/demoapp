package main

import (
	"fmt"
	"net/http"
    "time"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	fmt.Fprintf(w, "This message was provided by the host %s at %s", hostname, time.Now())
}


func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8800", nil)
}
