package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
)

// FileInfo represents file metadata including its hash for duplication checking.
type FileInfo struct {
	Name    string
	Size    int64
	ModTime string
	Hash    string
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run find_duplicates.go <directory>")
		return
	}

	rootDir := os.Args[1]
	fileInfos, err := findDuplicates(rootDir)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	if len(fileInfos) > 0 {
		fmt.Println("Duplicate Files Found:")
		sort.Slice(fileInfos, func(i, j int) bool {
			return fileInfos[i].Name < fileInfos[j].Name
		})
		for _, fi := range fileInfos {
			fmt.Printf("%s (%d bytes, modified: %s) -> Hash: %s\n", fi.Name, fi.Size, fi.ModTime, fi.Hash)
		}
	} else {
		fmt.Println("No duplicate files found.")
	}
}

// findDuplicates walks through the directory to gather file information and find duplicates.
func findDuplicates(rootDir string) ([]FileInfo, error) {
	var fileInfos []FileInfo
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Read the file's content to calculate its hash
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return fmt.Errorf("error reading file %s: %v", path, err)
		}

		hash := calculateHash(data)

		// Create a new FileInfo struct and append it to the slice
		fi := FileInfo{
			Name:    filepath.Base(path),
			Size:    info.Size(),
			ModTime: info.ModTime().String(),
			Hash:    hash,
		}
		fileInfos = append(fileInfos, fi)
		return nil
	})

	if err != nil {
		return nil, err
	}

	// Find duplicates by comparing hashes
	duplicateFileInfos := findDuplicatesByHash(fileInfos)
	return duplicateFileInfos, nil
}

// calculateHash calculates a SHA-1 hash of the given data.
func calculateHash(data []byte) string {
	hash := sha1.Sum(data)
	return hex.EncodeToString(hash[:])
}

// findDuplicatesByHash identifies duplicate files based on their hash values.
func findDuplicatesByHash(fileInfos []FileInfo) []FileInfo {
	// Create a map to group files by hash
	hashToFiles := make(map[string][]FileInfo)
	for _, fi := range fileInfos {
		hashToFiles[fi.Hash] = append(hashToFiles[fi.Hash], fi)
	}

	// Collect duplicates
	var duplicateFileInfos []FileInfo
	for _, files := range hashToFiles {
		if len(files) > 1 {
			duplicateFileInfos = append(duplicateFileInfos, files...)
		}
	}
	return duplicateFileInfos
}
