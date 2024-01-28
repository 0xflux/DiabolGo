package main

import (
	"fmt"
	"html"
	"net/http"
)

// function to print out full logs of HTTP intercepted
func getLogs(res http.ResponseWriter, req *http.Request) {
	webEntities.Mutex.Lock()
	defer webEntities.Mutex.Unlock()

	fmt.Fprint(res, "<html><body><pre>")

	for _, entity := range webEntities.Entities {
		fmt.Fprint(res, "******************************************************************************************************\n")
		safeUrl := html.EscapeString(entity.Url)
		safeHeaders := html.EscapeString(entity.Headers)
		safeBody := html.EscapeString(entity.Body)

		fmt.Fprintf(res, "URL: %s\nHeader: %v\nBody: %v\n\n", safeUrl, safeHeaders, safeBody)
	}

	fmt.Fprint(res, "</pre></body></html>")
}

// function to print out only the urls over HTTP intercepted
func getUrls(res http.ResponseWriter, req *http.Request) {
	webEntities.Mutex.Lock()
	defer webEntities.Mutex.Unlock()

	fmt.Fprint(res, "<html><body><pre>")

	for _, entity := range webEntities.Entities {
		safeUrl := html.EscapeString(entity.Url)

		fmt.Fprintf(res, "URL: %s\n", safeUrl)
	}

	fmt.Fprint(res, "</pre></body></html>")
}
