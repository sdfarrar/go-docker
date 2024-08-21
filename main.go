package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "Hello world\n")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Print("Starting server at http://localhost:8888\n")
	log.Fatal(http.ListenAndServe(":8888", nil))
}
