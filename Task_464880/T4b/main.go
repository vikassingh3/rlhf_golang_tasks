package main

import (
	"bufio"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"sync"
)

const (
	minFileSize = 1024 * 1024 // Process files greater than 1MB only
	bufferSize  = 1024 * 1024 // 1MB buffer
)

type FileInfo struct {
	path     string
	size     int64
	hash     string
	fileType string
}

func getFileType(path string) string {
	return filepath.Ext(path)
}

func hashFile(filePath string, wg *sync.WaitGroup, ch chan<- FileInfo) {
	defer wg.Done()

	hash := sha1.New()
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", filePath, err)
		return
	}
	defer file.Close()

	reader := bufio.NewReaderSize(file, bufferSize)
	if _, err := io.Copy(hash, reader); err != nil {
		fmt.Printf("Error copying file %s: %v\n", filePath, err)
		return
	}

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("Error getting file info %s: %v\n", filePath, err)
		return
	}

	c := FileInfo{
		path:     filePath,
		size:     fileInfo.Size(),
		hash:     hex.EncodeToString(hash.Sum(nil)),
		fileType: getFileType(filePath),
	}

	// Send FileInfo into the channel
	ch <- c
}

func processDirectory(dirPath string) []FileInfo {
	var files []FileInfo
	var wg sync.WaitGroup
	ch := make(chan FileInfo)

	// Goroutine to collect files from channel
	go func() {
		for fileInfo := range ch {
			files = append(files, fileInfo)
		}
	}()

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error walking directory %s: %v\n", dirPath, err)
			return nil
		}

		if info.IsDir() {
			return nil
		}

		if info.Size() < minFileSize {
			// Skip small files
			return nil
		}

		// Add a new file to be processed concurrently
		wg.Add(1)
		go hashFile(path, &wg, ch)
		return nil
	})
	if err != nil {
		fmt.Printf("Error walking directory %s: %v\n", dirPath, err)
	}

	// Wait for all goroutines to finish
	wg.Wait()
	// Close the channel after all processing is done
	close(ch)
	return files
}

func findDuplicates(files []FileInfo) {
	if len(files) == 0 {
		fmt.Println("No files found in the directory.")
		return
	}

	// Sort files by hash
	sort.Slice(files, func(i, j int) bool {
		return files[i].hash < files[j].hash
	})

	var currentHash string
	var duplicates []FileInfo

	for _, file := range files {
		if file.hash != currentHash {
			if len(duplicates) > 1 {
				// Print the duplicates for the current hash
				fmt.Printf("Duplicates found with hash %s:\n", currentHash)
				for _, dup := range duplicates {
					fmt.Printf("  %s (%d bytes, %s)\n", dup.path, dup.size, dup.fileType)
				}
			}
			// Reset for the next hash
			currentHash = file.hash
			duplicates = nil
		}

		duplicates = append(duplicates, file)
	}

	// Handle the last set of duplicates
	if len(duplicates) > 1 {
		fmt.Printf("Duplicates found with hash %s:\n", currentHash)
		for _, dup := range duplicates {
			fmt.Printf("  %s (%d bytes, %s)\n", dup.path, dup.size, dup.fileType)
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run find_duplicates.go <directory>")
		return
	}

	dirPath := os.Args[1]
	files := processDirectory(dirPath)
	findDuplicates(files)
}
