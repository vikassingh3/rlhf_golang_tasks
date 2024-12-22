package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
)

func main() {
	// Read command line arguments
	if len(os.Args) != 3 || os.Args[1] != "-directory" {
		fmt.Println("Usage: go run main.go -directory=<directory path>")
		return
	}
	dirPath := os.Args[2]

	// Find redundant files
	redundantFiles, err := findRedundantFiles(dirPath)
	if err != nil {
		log.Fatalf("Error finding redundant files: %v", err)
	}

	// Prompt user before removing redundant files
	fmt.Printf("Found %d redundant files. Do you want to remove them? (yes/no): ", len(redundantFiles))
	var confirm string
	fmt.Scanln(&confirm)
	if confirm == "yes" {
		for _, filePath := range redundantFiles {
			if err := os.Remove(filePath); err != nil {
				log.Printf("Error removing file %s: %v", filePath, err)
			}
		}
		fmt.Println("Redundant files removed successfully.")
	} else {
		fmt.Println("Operation canceled. Redundant files not removed.")
	}
}

// findRedundantFiles finds and returns the paths of redundant files in the given directory.
func findRedundantFiles(dirPath string) ([]string, error) {
	var fileSizes []struct {
		size  int64
		paths []string
	}

	// Walk through the directory and collect file sizes
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		fileSizes = append(fileSizes, struct {
			size  int64
			paths []string
		}{size: info.Size(), paths: []string{path}})
		return nil
	})
	if err != nil {
		return nil, err
	}

	// Sort file sizes in ascending order
	sort.Slice(fileSizes, func(i, j int) bool {
		return fileSizes[i].size < fileSizes[j].size
	})

	var redundantFiles []string
	for i := 1; i < len(fileSizes); i++ {
		if fileSizes[i].size == fileSizes[i-1].size {
			redundantFiles = append(redundantFiles, fileSizes[i].paths...)
		}
	}

	return redundantFiles, nil
}
