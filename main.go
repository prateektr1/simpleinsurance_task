package main

import (
	"encoding/gob"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

const (
	dataFileName = "request_counter.gob"
	windowSize   = 60 * time.Second
)

type RequestCounter struct {
	Counter int
	LastTS  time.Time
	sync.Mutex
}

func (rc *RequestCounter) Increment() {
	rc.Lock()
	defer rc.Unlock()
	rc.Counter++
	rc.LastTS = time.Now()
}

func (rc *RequestCounter) Count() int {
	rc.Lock()
	defer rc.Unlock()
	return rc.Counter
}

func (rc *RequestCounter) Cleanup() {
	rc.Lock()
	defer rc.Unlock()
	if time.Since(rc.LastTS) > windowSize {
		rc.Counter = 0
	}
}

func loadCounter() *RequestCounter {
	file, err := os.Open(dataFileName)
	if err != nil {
		return &RequestCounter{}
	}
	defer file.Close()

	var rc RequestCounter
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&rc)
	if err != nil {
		return &RequestCounter{}
	}

	rc.Cleanup()
	return &rc
}

func saveCounter(rc *RequestCounter) {
	file, err := os.Create(dataFileName)
	if err != nil {
		fmt.Println("Error saving data:", err)
		return
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(rc)
	if err != nil {
		fmt.Println("Error encoding data:", err)
	}
}

func main() {
	rc := loadCounter()
	go func() {
		for {
			time.Sleep(time.Second)
			rc.Cleanup()
			saveCounter(rc)
		}
	}()

	http.HandleFunc("/numberOfRequests", func(w http.ResponseWriter, r *http.Request) {
		rc.Increment()
		fmt.Fprintf(w, "Total requests in the last 60 seconds: %d", rc.Count())
	})

	http.ListenAndServe(":8080", nil)
}
