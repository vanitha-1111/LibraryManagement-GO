package main

import (
	"bytes"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	url := "http://localhost:8080/borrowdetails"

	// IMPORTANT: set these to the IDs created during your manual setup
	borrowID := 10 // replace with actual borrow_id
	bookID := 49   // replace with actual book_id

	body := []byte(fmt.Sprintf(`{"borrow_id": %d, "book_id": %d}`, borrowID, bookID))

	var wg sync.WaitGroup
	wg.Add(20)

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	for i := 0; i < 20; i++ {
		go func(i int) {
			defer wg.Done()

			resp, err := client.Post(url, "application/json", bytes.NewReader(body))
			if err != nil {
				fmt.Printf("[%d] ERROR: %v\n", i, err)
				return
			}

			fmt.Printf("[%d] STATUS: %d\n", i, resp.StatusCode)
		}(i)
	}

	wg.Wait()

	fmt.Println("All goroutines completed.")
}
