package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func processData(ctx context.Context, data []byte) error {
	// Simulate processing
	time.Sleep(2 * time.Second)
	return nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Read a large file
	file, err := os.Open("./large_file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	// Process data in batches
	batchSize := 100000
	for i := 0; i < len(data); i += batchSize {
		// Ensure we don't go out of bounds by adjusting batch size at the end of the data slice
		end := i + batchSize
		if end > len(data) {
			end = len(data)
		}
		batch := data[i:end]

		// Create a new context for each batch to handle potential timeouts
		batchCtx, batchCancel := context.WithTimeout(ctx, 3*time.Second)
		defer batchCancel()

		select {
		case <-batchCtx.Done():
			fmt.Println("Batch processing canceled.")
			return
		default:
			go func() {
				if err := processData(batchCtx, batch); err != nil {
					log.Println("Error processing batch:", err)
				} else {
					fmt.Println("Batch processed successfully.")
				}
			}()
		}
	}

	fmt.Println("All batches processed.")
}
