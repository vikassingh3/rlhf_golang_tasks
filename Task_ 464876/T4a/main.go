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

type FileInfo struct {
	path     string
	size     int64
	hash     string
	fileType string
}

func getFileType(path string) string {
	return filepath.Ext(path)
}

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

func processDirectory(dirPath string, wg *sync.WaitGroup) {
	defer wg.Done()

	fileHashes := make(map[string]map[string]int) // map of hash -> fileType -> count

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		fileType := getFileType(path)
		fileHash := hashFile(path)

		if fileHash == "" {
			return nil
		}

		if fileHashes[fileHash] == nil {
			fileHashes[fileHash] = make(map[string]int)
		}

		fileHashes[fileHash][fileType]++
		return nil
	})
	if err != nil {
		fmt.Printf("Error walking directory %s: %v\n", dirPath, err)
	}

	// Print duplicates
	for fileHash, fileTypes := range fileHashes {
		if len(fileTypes) > 1 {
			fmt.Printf("Duplicates found with hash %s:\n", fileHash)
			for fileType, count := range fileTypes {
				suffix := ""
				if count > 1 {
					suffix = "s"
				}
				fmt.Printf("  %d file%s of type %s\n", count, suffix, fileType)
			}
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run find_duplicates.go <directory>")
		return
	}

	dirPath := os.Args[1]
	var wg sync.WaitGroup

	// Process directory concurrently
	wg.Add(1)
	go processDirectory(dirPath, &wg)

	wg.Wait()
}