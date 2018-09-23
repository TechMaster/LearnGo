package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)
func handler(w http.ResponseWriter, r *http.Request) {
	println("--->", os.Args[1], r.URL.String())
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}


func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s <port>", os.Args[0])
	}
	if _, err := strconv.Atoi(os.Args[1]); err != nil {
		log.Fatalf("Invalid port: %s (%s)\n", os.Args[1], err)
	}

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":"+os.Args[1], nil))
}