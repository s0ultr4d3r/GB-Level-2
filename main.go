package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

var mu sync.Mutex

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		prot := "https://" + scanner.Text()
		lines = append(lines, prot)
	}
	return lines, scanner.Err()
}
func takeMyUrls() []string {

	lines, err := readLines("urls.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	return lines
}

func main() {
	mu.Lock()
	defer mu.Unlock()
	var wg sync.WaitGroup
	var urls = takeMyUrls()
	for _, url := range urls {
		// Increment the WaitGroup counter.
		wg.Add(1)
		// Launch a goroutine to fetch the URL.
		go func(url string) {
			// Decrement the counter when the goroutine completes.
			defer wg.Done()
			// Fetch the URL.
			_, err := http.Get(url)
			if err != nil {
				fmt.Printf("error %v \n", err)
			}
		}(url)
	}
	// Wait for all HTTP fetches to complete.
	wg.Wait()
	return
}
