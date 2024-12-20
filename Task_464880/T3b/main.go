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
	"strings"
)

type FileInfo struct {
	path      string
	size      int64
	hash      string
	extension string
}

type DuplicateGroup struct {
	Hash      string
	Files     []FileInfo
	TotalSize int64
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
		fmt.Printf("Error hashing file %s: %v\n", filePath, err)
		return ""
	}
	return hex.EncodeToString(hash.Sum(nil))
}

func processDirectory(dirPath string) []FileInfo {
	var files []FileInfo
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error accessing file %s: %v\n", path, err)
			return nil
		}
		if info.IsDir() {
			return nil
		}
		ext := filepath.Ext(path)
		fileInfo := FileInfo{
			path:      path,
			size:      info.Size(),
			hash:      hashFile(path),
			extension: strings.TrimPrefix(ext, "."),
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

	var duplicateGroups []DuplicateGroup
	var currentHash string
	var duplicates []FileInfo

	for _, file := range files {
		if file.hash != currentHash {
			if len(duplicates) > 1 {
				totalSize := int64(0)
				for _, dup := range duplicates {
					totalSize += dup.size
				}
				duplicateGroups = append(duplicateGroups, DuplicateGroup{
					Hash:      currentHash,
					Files:     duplicates,
					TotalSize: totalSize,
				})
			}
			currentHash = file.hash
			duplicates = nil
		}
		duplicates = append(duplicates, file)
	}

	// Process the last duplicate group
	if len(duplicates) > 1 {
		totalSize := int64(0)
		for _, dup := range duplicates {
			totalSize += dup.size
		}
		duplicateGroups = append(duplicateGroups, DuplicateGroup{
			Hash:      currentHash,
			Files:     duplicates,
			TotalSize: totalSize,
		})
	}

	// Print duplicate groups
	if len(duplicateGroups) > 0 {
		fmt.Println("Duplicate Groups Found:")
		for _, group := range duplicateGroups {
			fmt.Printf("\nHash: %s\n", group.Hash)
			fmt.Printf("Total Size: %d bytes\n", group.TotalSize)
			fmt.Println("Files:")
			for _, file := range group.Files {
				fmt.Printf(" - %s (%d bytes, extension: %s)\n", file.path, file.size, file.extension)
			}
		}
	} else {
		fmt.Println("No duplicate groups found.")
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run find_duplicates.go <directory>")
		return
	}

	dirPath := os.Args[1]
	files := processDirectory(dirPath)
	findDuplicates(files)
}
