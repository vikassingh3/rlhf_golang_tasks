package main

import (
	"bufio"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
)

var (
	fileSizes  = make(map[string]int64)
	fileHashes = make(map[string][]string)
	mu         sync.Mutex
)

func hashFile(filePath string) string {
	hash := sha1.New()

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", filePath, err)
		return ""
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	_, err = io.Copy(hash, reader)
	if err != nil {
		fmt.Printf("Error copying file %s: %v\n", filePath, err)
		return ""
	}

	return hex.EncodeToString(hash.Sum(nil))
}

func processFile(filePath string) {
	mu.Lock()
	defer mu.Unlock()

	// Get file info
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		fmt.Printf("Error getting file size for %s: %v\n", filePath, err)
		return
	}

	// Get file size
	size := fileInfo.Size()
	fileSizes[filePath] = size

	// Compute file hash
	fileHash := hashFile(filePath)
	if fileHash == "" {
		return
	}

	// Store file hash and associated file path
	fileHashes[fileHash] = append(fileHashes[fileHash], filePath)
}

func findDuplicates(rootDir string) {
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		processFile(path)
		return nil
	})
	if err != nil {
		fmt.Printf("Error walking directory %s: %v\n", rootDir, err)
	}
}

func printDuplicates() {
	for fileHash, filePaths := range fileHashes {
		if len(filePaths) > 1 {
			fmt.Printf("Duplicates found with hash %s:\n", fileHash)
			for _, path := range filePaths {
				fmt.Printf("  %s (%d bytes)\n", path, fileSizes[path])
			}
		}
	}
}

func removeDuplicates() {
	for _, filePaths := range fileHashes {
		if len(filePaths) > 1 {
			for _, path := range filePaths[1:] {
				err := os.Remove(path)
				if err != nil {
					fmt.Printf("Error removing file %s: %v\n", path, err)
				}
			}
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run find_redundant_files.go <directory>")
		return
	}

	rootDir := os.Args[1]

	findDuplicates(rootDir)
	printDuplicates()

	// Ask user if they want to remove duplicates
	fmt.Println("Do you want to remove the duplicates? (y/n)")
	var response string
	fmt.Scan(&response)
	if response == "y" {
		removeDuplicates()
	}
}
