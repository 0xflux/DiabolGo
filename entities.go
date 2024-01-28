package main

import (
	"fmt"
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
