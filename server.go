package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
)

type Record struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Index\n")
}

func serveRecords(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		readRecord(w, r)
		return
	}

	if r.Method == http.MethodPost {
		saveRecord(w, r)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

var p = regexp.MustCompile("/records/([0-9]+)")

func readRecord(w http.ResponseWriter, r *http.Request) {
	m := p.FindStringSubmatch(r.URL.Path)

	// no match found (m[0] matches the whole string)
	if len(m) < 2 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Printf("Found match: %q\n", m[1])
	i, err := strconv.ParseInt(m[1], 10, 0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id := int(i)

	rec := &Record{ID: id, Name: "testRecord"}
	if err := json.NewEncoder(w).Encode(rec); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	return
}

func saveRecord(w http.ResponseWriter, r *http.Request) {

}
