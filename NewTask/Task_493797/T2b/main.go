
package main

import (
	"fmt"
	"os"
)

// File struct with Version field
type File struct {
	Content []byte
	Version int
}

// openFile function to open and read a file along with its version
func openFile(filename string) (*File, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Read the version number from the first line of the file
	var version int
	_, err = fmt.Fscanf(file, "%d\n", &version)
	if err != nil {
		return nil, err
	}

	return &File{Content: content, Version: version}, nil
}

// writeFile function to write a file with the updated content and version
func writeFile(filename string, file *File) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	// Write the updated version number and content to the file
	_, err = fmt.Fprintf(f, "%d\n", file.Version)
	if err != nil {
		return err
	}
	_, err = f.Write(file.Content)
	return err
}

func resolveConflict(serverFile *File, userFile *File) *File {
	// Implement your conflict resolution strategy here (e.g., use merge algorithm)
	// In this simple example, we'll just overwrite the server file with the user file
	return userFile
}

func main() {
	serverFilename := "./server_file.txt"
	userFilename := "./user_file.txt"

	// Load the files from disk
	serverFile, err := openFile(serverFilename)
	if err != nil {
		fmt.Printf("Error opening server file: %v\n", err)
		return
	}
	userFile, err := openFile(userFilename)
	if err != nil {
		fmt.Printf("Error opening user file: %v\n", err)
		return
	}

	// Simulate a conflict by editing both files separately
	serverFile.Content = append(serverFile.Content, []byte(" Server change")...)
	serverFile.Version++

	userFile.Content = append(userFile.Content, []byte(" User change")...)
	userFile.Version++

	if serverFile.Version != userFile.Version {
		fmt.Println("Conflict detected!")
		// Resolve the conflict (e.g., using versioning, merge, etc.)
		updatedFile := resolveConflict(serverFile, userFile)

		// Write the updated file back to disk (overwriting the server file for this example)
		err = writeFile(serverFilename, updatedFile)
		if err != nil {
			fmt.Printf("Error writing file: %v\n", err)
		} else {
			fmt.Println("File resolved and updated successfully.")
		}
	} else {
		fmt.Println("No conflicts detected.")
	}
}
