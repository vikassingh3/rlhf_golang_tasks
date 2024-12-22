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
)

type FileInfo struct {
	path     string
	size     int64
	hash     string
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

func processDirectory(dirPath string) []FileInfo {
	var files []FileInfo

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		fileInfo := FileInfo{
			path:     path,
			size:     info.Size(),
			hash:     hashFile(path),
		}
		files = append(files, fileInfo)
		return nil
	})
	if err != nil {
		fmt.Printf("Error walking directory %s: %v\n", dirPath, err)
	}

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
	var duplicates []string

	for _, file := range files {
		if file.hash != currentHash {
			currentHash = file.hash
			duplicates = nil
		}

		duplicates = append(duplicates, file.path)
		if len(duplicates) > 1 {
			fmt.Printf("Duplicates found with hash %s:\n", currentHash)
			for _, path := range duplicates {
				fmt.Printf("  %s (%d bytes)\n", path, file.size)
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
	files := processDirectory(dirPath)
	findDuplicates(files)
}