package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	addr := ":8080"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from Slipway! commit=%s\n", os.Getenv("GIT_SHA"))
	})
	log.Println("listening on", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
