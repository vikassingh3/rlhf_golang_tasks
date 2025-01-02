
package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

// File struct represents a file entry with metadata
type File struct {
	ID      int
	Name    string
	UserID  int
	Data    []byte // Actual file data can be stored differently (e.g., in a database or filesystem)
	Shared  bool
}

var files []File

// AddFile handler to add a file
func AddFile(c *gin.Context) {
	var newFile File
	// In a real-world scenario, you'd read the file data from an HTTP request.
	// Here, we'll create some dummy data for simplicity.
	newFile.Name = c.Query("filename")
	newFile.UserID = atoi(c.Query("user_id"))
	newFile.Data = []byte(`This is the file data.`) // This is just a placeholder
	newFile.Shared = false

	// Increment the ID
	newFile.ID = len(files) + 1

	files = append(files, newFile)

	c.JSON(http.StatusOK, gin.H{
		"message": "File added successfully.",
		"file":    newFile,
	})
}

// ListFiles handler to list files of a given user
func ListFiles(c *gin.Context) {
	userID := atoi(c.Query("user_id"))
	var userFiles []File

	for _, file := range files {
		if file.UserID == userID {
			userFiles = append(userFiles, file)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"files": userFiles,
	})
}

// ShareFile handler to share a file
func ShareFile(c *gin.Context) {
	fileID := atoi(c.Query("file_id"))

	for i, file := range files {
		if file.ID == fileID {
			files[i].Shared = true
			break
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "File shared successfully.",
	})
}

// ListSharedFiles handler to list shared files
func ListSharedFiles(c *gin.Context) {
	var sharedFiles []File

	for _, file := range files {
		if file.Shared {
			sharedFiles = append(sharedFiles, file)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"shared_files": sharedFiles,
	})
}

func atoi(str string) int {
	if i, err := strconv.Atoi(str); err == nil {
		return i
	}
	panic("Invalid integer input")
}

func main() {
	r := gin.Default()
	r.POST("/upload", AddFile)
	r.GET("/files", ListFiles)
	r.GET("/share", ShareFile)
	r.GET("/shared", ListSharedFiles)

	if err := r.Run(":8080"); err != nil {
		fmt.Println("Server error:", err)
		os.Exit(1)
	}
}


