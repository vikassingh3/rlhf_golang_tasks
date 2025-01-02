package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/gofrs/flock"
)

type FileData struct {
	Content string `json:"content"`
	Version string `json:"version"`
}

func main() {
	filePath := "./example.txt"
	lockFilePath := filePath + ".lock"
	
	// Load initial file data with a version
	fileData, err := loadFileData(filePath)
	if err != nil {
		fmt.Printf("Error loading file: %v\n", err.Error())
		return
	}
	
	// Simulate user editing the file
	fileData.Content += " Edited by User1"
	fileData.Version = generateVersion(fileData.Version)
	
	// Attempt to save changes
	saved, err := saveFileData(filePath, lockFilePath, fileData)
	if err != nil {
		fmt.Printf("Error saving file: %v\n", err)
		return
	}
	
	if !saved {
		fmt.Println("Conflict detected. Please try again.")
		return
	}
	
	fmt.Println("File saved successfully.")
}

func loadFileData(filePath string) (*FileData, error) {
	file, err := os.OpenFile(filePath, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	
	var fileData FileData
	if err := json.NewDecoder(file).Decode(&fileData); err != nil {
		if os.IsNotExist(err) {
			return &FileData{Content: "", Version: generateVersion("")}, nil
		}
		return nil, err
	}
	
	return &fileData, nil
}

func saveFileData(filePath, lockFilePath string, fileData *FileData) (bool, error) {
	lock := flock.New(lockFilePath) // Create a new flock instance for locking
	locked, err := lock.TryLock()
	if err != nil {
		return false, err
	}
	defer lock.Unlock()
	
	if !locked {
		return false, fmt.Errorf("could not acquire lock")
	}
	
	currentData, err := loadFileData(filePath)
	if err != nil {
		return false, err
	}
	
	if currentData.Version != fileData.Version {
		return false, fmt.Errorf("conflict detected: version mismatch")
	}
	
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return false, err
	}
	defer file.Close()
	
	if err := json.NewEncoder(file).Encode(fileData); err != nil {
		return false, err
	}
	
	return true, nil
}

func generateVersion(currentVersion string) string {
	return fmt.Sprintf("%s-%d", currentVersion, time.Now().UnixNano())
}
