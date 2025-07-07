package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

type FetchResult struct {
	URL    string
	Status string
	Error  error
}

func fetchUrl(url string, resultCh chan<- FetchResult, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("\nFetching %s", url)

	resp, err := http.Get(url)
	if err != nil {
		resultCh <- FetchResult{URL: url, Error: err}
		return
	}

	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		resultCh <- FetchResult{URL: url, Status: resp.Status, Error: nil}
		return
	}

	resultCh <- FetchResult{URL: url, Status: resp.Status, Error: nil}
}

func main() {
	urls := []string{
		"https://www.google.com",
		"https://www.bing.com",
		"https://www.yahoo.com",
		"https://invalid-url-example.com",
	}

	resultCh := make(chan FetchResult, len(urls))
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		fetchUrl(url, resultCh, &wg)
	}

	wg.Wait()
	close(resultCh)

	fmt.Printf("\n---- Fetch Results -----\n")
	for result := range resultCh {
		if result.Error != nil {
			fmt.Printf("URL: %s, Error: %v\n", result.URL, result.Error)
		} else {
			fmt.Printf("URL: %s, Status: %s\n", result.URL, result.Status)
		}
	}
}
