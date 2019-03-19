package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Printf("Start of demonstration.")

	http.HandleFunc("/", index)
	http.HandleFunc("/records/", serveRecords)

	log.Fatal(http.ListenAndServe(":8081", http.DefaultServeMux))
}
