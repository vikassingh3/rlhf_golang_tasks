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
	fileType string // The file extension/type
}

type Summary struct {
	fileType string // The type or extension of the file
	size     int64  // The size of the file
	count    int    // The count of files matching this type and size
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

func processDirectory(dirPath string) []FileInfo {
	var files []FileInfo

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		fileInfo := FileInfo{
			path:     path,
			size:     info.Size(),
			hash:     hashFile(path),
			fileType: getFileType(path),
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
	var duplicates []FileInfo

	for _, file := range files {
		if file.hash != currentHash {
			if len(duplicates) > 1 {
				fmt.Printf("\nDuplicates found with hash %s:\n", currentHash)
				for _, dup := range duplicates {
					fmt.Printf("  %s (%d bytes, %s)\n", dup.path, dup.size, dup.fileType)
				}
			}
			currentHash = file.hash
			duplicates = nil
		}

		duplicates = append(duplicates, file)
	}

	// Print the last group of duplicates, if any
	if len(duplicates) > 1 {
		fmt.Printf("\nDuplicates found with hash %s:\n", currentHash)
		for _, dup := range duplicates {
			fmt.Printf("  %s (%d bytes, %s)\n", dup.path, dup.size, dup.fileType)
		}
	}
}

func summarizeDuplicates(files []FileInfo) {
	var summary []Summary
	duplicateHashes := make(map[string]bool)

	for _, file := range files {
		if duplicateHashes[file.hash] {
			existing := findSummary(summary, file.fileType, file.size)
			if existing != nil {
				existing.count++
			} else {
				summary = append(summary, Summary{fileType: file.fileType, size: file.size, count: 1})
			}
		} else {
			duplicateHashes[file.hash] = true
		}
	}

	// Print the summary
	fmt.Println("\nSummary of duplicate files:")
	fmt.Println("File Type - Size - Count")
	for _, item := range summary {
		fmt.Printf("%s - %d bytes - %d\n", item.fileType, item.size, item.count)
	}
}

func findSummary(summary []Summary, fileType string, size int64) *Summary {
	for i, item := range summary {
		if item.fileType == fileType && item.size == size {
			return &summary[i]
		}
	}
	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run find_duplicates.go <directory>")
		return
	}

	dirPath := os.Args[1]
	files := processDirectory(dirPath)
	findDuplicates(files)
	summarizeDuplicates(files)
}
