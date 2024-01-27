package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
)

type WebEntity struct {
	Url     string
	Body    string
	Headers string
}

type WebEntities struct {
	Entities []WebEntity
	Mutex    sync.Mutex
}

func formatHeaders(headers http.Header) string {
	var headerStr string
	for key, values := range headers {
		headerStr += key + ": " + strings.Join(values, ", ") + "\n"
	}
	return headerStr
}

// add a new entity
func (w1 *WebEntities) Add(entity WebEntity) {
	w1.Mutex.Lock()
	defer w1.Mutex.Unlock() // Corrected this line
	w1.Entities = append(w1.Entities, entity)
}

// print entities
func (w1 *WebEntities) Print() {
	defer w1.Mutex.Lock()
	for _, entity := range w1.Entities {
		fmt.Printf("URL: %s, Header: %v, Body: %v\n", entity.Url, entity.Headers, entity.Body)
	}
}

// global for list
var webEntities = WebEntities{}

func intercept(res http.ResponseWriter, req *http.Request) {

	// fmt.Println("Requested URL: ", req.URL.String())

	req.RequestURI = ""
	client := http.Client{}

	// forward the request
	resp, err := client.Do(req)
	if err != nil {
		http.Error(res, "Server Error", http.StatusInternalServerError)
		log.Println("Server error: ", err)
		return
	}
	defer resp.Body.Close()

	// read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(res, "Error reading response body", http.StatusInternalServerError)
		log.Println("Error reading response body: ", err)
		return
	}

	// print headers and body
	// for key, values := range resp.Header {
	// 	fmt.Printf("%s: %v\n", key, values)
	// }

	// fmt.Println("Response body: ", string(body))

	// write to the list
	if req.URL.Scheme == "http" {

		webEntities.Add(WebEntity{
			Url:     req.URL.String(),
			Body:    string(body),
			Headers: formatHeaders(resp.Header),
		})

		// webEntities.Print()
	}

	// copy the response headers and status code to the OG response
	for key, values := range resp.Header {
		for _, value := range values {
			res.Header().Add(key, value)
		}
	}
	res.WriteHeader(resp.StatusCode)

	// send the response back to the requester
	_, err = res.Write(body)
	if err != nil {
		log.Println("Error writing stream back to requester: ", err)
	}
}

func main() {
	port := 1234

	http.HandleFunc("/", intercept)

	fmt.Println("Starting listener on port: ", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))

}
