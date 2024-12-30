package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type ConnectionPool struct {
	// Implement resource pooling logic here
}

func getData(ctx context.Context, client *http.Client, url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := &http.Client{}
	// pool := &ConnectionPool{}

	// Process multiple URLs
	urls := []string{"http://example.com/api/data1", "http://example.com/api/data2"}

	for _, url := range urls {
		batchCtx, batchCancel := context.WithTimeout(ctx, 2*time.Second)
		defer batchCancel()

		select {
		case <-batchCtx.Done():
			fmt.Println("Timeout for URL:", url)
		default:
			_, err := getData(batchCtx, client, url)
			if err != nil {
				log.Println("Error fetching data from", url, ":", err)
			} else {
				fmt.Println("Data fetched from", url)
			}
		}
	}
}