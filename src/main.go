package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request from ", r.RemoteAddr)
	hostname, _ := os.Hostname()
	fmt.Fprintln(w, "You've hit ", hostname)
}

func version_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "0.1")
}

func id_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, os.Getenv("NODE_ID"))
}

func main() {
	fmt.Println("Server is starting...")
	http.HandleFunc("/", handler)
	http.HandleFunc("/version", version_handler)
	http.HandleFunc("/id", id_handler)
	log.Fatal(http.ListenAndServe(":8888", nil))
}
