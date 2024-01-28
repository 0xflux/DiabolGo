package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
)

func intercept(res http.ResponseWriter, req *http.Request) {

	fmt.Println("Requested URL: ", req.URL.String())

	// downgrade HTTPs for POC purposes
	if req.URL.Scheme == "https" {
		fmt.Println("HTTPS detected. Requested URL: ", req.URL.String())
		req.URL.Scheme = "http"

		// handle redirections
		req.Header.Del("Upgrade-Insecure-Requests")

		fmt.Println("HTTPS detected. Requested URL: ", req.URL.String())

	}

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

// handle TCP tunnelling
func HandleTunnelling(w http.ResponseWriter, r *http.Request) {
	destConn, err := net.Dial("tcp", r.Host)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	w.WriteHeader(http.StatusOK)
	hijacker, ok := w.(http.Hijacker)
	if !ok {
		http.Error(w, "Hijacking not supported", http.StatusInternalServerError)
		return
	}

	clientConn, _, err := hijacker.Hijack()
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	go transfer(destConn, clientConn)
	go transfer(clientConn, destConn)
}

func transfer(destination io.WriteCloser, source io.ReadCloser) {
	defer destination.Close()
	defer source.Close()
	io.Copy(destination, source)
}
