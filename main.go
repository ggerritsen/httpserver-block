package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {
	fmt.Printf("Start of demonstration.\n")

	http.HandleFunc("/", index)
	http.HandleFunc("/records/", serveRecords)

	c := make(chan error, 1)
	go func() {
		c <- http.ListenAndServe(":8081", http.DefaultServeMux)
	}()

	// check Index page
	req, err := http.NewRequest("GET", "http://localhost:8081/", nil)
	if err != nil {
		log.Fatal(err)
	}
	if err := checkResponse(req, http.StatusOK, "Index"); err != nil {
		log.Fatal(err)
	}

	// check retrieving records
	req, err = http.NewRequest("GET", "http://localhost:8081/records/1", nil)
	if err != nil {
		log.Fatal(err)
	}
	if err := checkResponse(req, http.StatusOK, "{\"id\":1,\"name\":\"testRecord\"}\n"); err != nil {
		log.Fatal(err)
	}

	// check creating records
	v := url.Values{"name": []string{"test"}}
	resp, err := http.PostForm("http://localhost:8081/records/", v)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		log.Fatalf("Unexpected status code: %d", resp.StatusCode)
	}

	close(c)
	err = <-c
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("End of demonstration.\n")
}

func checkResponse(req *http.Request, expectedStatus int, expectedResponse string) error {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != expectedStatus {
		return fmt.Errorf("Unexpected status code: %d for request %s %s", resp.StatusCode, req.Method, req.URL.String())
	}

	bb, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if b := string(bb); b != expectedResponse {
		return fmt.Errorf("Unexpected response: %q", b)
	}

	return nil
}
