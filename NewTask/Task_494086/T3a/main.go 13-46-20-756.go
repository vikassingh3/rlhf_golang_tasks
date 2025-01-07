package main

import (
	"fmt"
	"sync"
	"time"
)

func download(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	// Simulate downloading a file
	time.Sleep(time.Duration(id) * time.Second)
	fmt.Println("Downloaded file", id)
}

func processDownloadedData(files []string) {
	fmt.Println("Processing downloaded data...")
	// Simulate data processing
	time.Sleep(2 * time.Second)
	fmt.Println("Data processing completed.")
}

func main() {
	var wg sync.WaitGroup
	files := []string{}

	// Simulate downloading multiple files
	wg.Add(3)
	go download(1, &wg)
	go download(2, &wg)
	go download(3, &wg)

	// Wait for all downloads to complete
	wg.Wait()

	// Process downloaded data
	processDownloadedData(files)
}