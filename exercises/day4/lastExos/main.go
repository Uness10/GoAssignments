package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

func fetchURL(ctx context.Context, url string, wg *sync.WaitGroup) {
	defer wg.Done()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		fmt.Printf("Error creating request for %s: %v\n", url, err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		select {
		case <-ctx.Done():
			fmt.Printf("Fetch canceled or timed out for URL: %s\n", url)
		default:
			fmt.Printf("Error fetching URL %s: %v\n", url, err)
		}
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response for URL %s: %v\n", url, err)
		return
	}

	fmt.Printf("Fetched URL: %s\nContent: %s\n", url, string(body))
}

func main() {
	urls := []string{
		"http://example.com",
		"http://httpbin.org/delay/2",
		"http://httpbin.org/delay/5",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go fetchURL(ctx, url, &wg)
	}

	wg.Wait()
	fmt.Println("All fetches complete or timed out")
}
