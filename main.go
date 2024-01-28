package main

import (
	"fmt"
	"log"
	"net/http"
)

// global for list
var webEntities = WebEntities{}

func main() {
	port := 1234

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodConnect {
			HandleTunnelling(w, r)
		} else {
			intercept(w, r) // Existing HTTP handler
		}
	})

	http.HandleFunc("/logs", getLogs) // handler for viewing logs
	http.HandleFunc("/urls", getUrls) // handler for viewing urls visited

	fmt.Println("Starting listener on port: ", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))

}
