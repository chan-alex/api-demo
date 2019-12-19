package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
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

func livenessProbe_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}

// Using Global variables here is the easier to do this.
var livenessProbeCount int = 0
var livenessProbeCountMax int = 5

func livenessProbe_failAtN_handler(w http.ResponseWriter, r *http.Request) {
	if livenessProbeCount < livenessProbeCountMax {
		fmt.Fprintln(w, "OK")
		livenessProbeCount++
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Return 500 liveness probe fail simulation")
	}
}

func main() {
	fmt.Println("Server is starting...")
	http.HandleFunc("/", handler)
	http.HandleFunc("/version", version_handler)
	http.HandleFunc("/id", id_handler)

	value, exists := os.LookupEnv("LIVEPROBE_FAIL_N")
	if exists {
		// Note Atoi() will parse nonsense text values too.
		i, err := strconv.Atoi(value)
		if err != nil {
			livenessProbeCountMax = i
			fmt.Println("Setting liveness probe to fail at ", value)
		}

		http.HandleFunc("/liveness", livenessProbe_failAtN_handler)
	} else {
		http.HandleFunc("/liveness", livenessProbe_handler)
	}
	log.Fatal(http.ListenAndServe(":8888", nil))
}
