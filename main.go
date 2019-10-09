package main

import (
	"os"
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) { 
	fmt.Fprintln(w, "OK") }

func version_handler(w http.ResponseWriter, r *http.Request) { 
	fmt.Fprintln(w, "0.1") }

func id_handler(w http.ResponseWriter, r *http.Request) { 
	fmt.Fprintln(w, os.Getenv("NODE_ID")) }


func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/version", version_handler)
	http.HandleFunc("/id", id_handler)
	log.Fatal(http.ListenAndServe(":8888", nil))
}
