package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	// IMPORTANT -- put your actual borrow_details_id here
	borrowDetailID := 11

	url := fmt.Sprintf("http://localhost:8080/borrowdetails/%d/return", borrowDetailID)

	client := &http.Client{Timeout: 5 * time.Second}

	var wg sync.WaitGroup
	wg.Add(20)

	success := 0
	fail := 0
	var mu sync.Mutex

	for i := 0; i < 20; i++ {
		go func(i int) {
			defer wg.Done()

			req, _ := http.NewRequest("PUT", url, nil)
			resp, err := client.Do(req)
			if err != nil {
				fmt.Printf("[%d] ERROR: %v\n", i, err)
				return
			}

			fmt.Printf("[%d] STATUS: %d\n", i, resp.StatusCode)

			mu.Lock()
			if resp.StatusCode == 200 {
				success++
			} else {
				fail++
			}
			mu.Unlock()

			resp.Body.Close()
		}(i)
	}

	wg.Wait()

	fmt.Println("All goroutines completed.")
	fmt.Println("Success:", success, "Failures:", fail)
}
