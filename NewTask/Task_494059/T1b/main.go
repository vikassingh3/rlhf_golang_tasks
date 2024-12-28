package main  
import (  
    "fmt"
    "os"
    "path/filepath"
)

// Define a File struct to store file metadata
type File struct {  
    ID       int
    Name     string
    Content  []byte // We'll use byte slices to store the file content
    OwnerID   int
}

// Define a User struct to store user information
type User struct {  
    ID       int
    Name     string
    Files    []*File // Slice to store pointers to files owned by the user
}

// Global slice to store all the users
var users []*User

// Global slice to store all the files
var files []*File

func main() {  
    // Create some users for demonstration purposes
    user1 := &User{ID: 1, Name: "Alice"}
    user2 := &User{ID: 2, Name: "Bob"}
    users = append(users, user1, user2)

    // Upload a file for Alice
    uploadFile(user1, "alice_photo.jpg")

    // Upload a file for Bob
    uploadFile(user2, "bob_video.mp4")

    // Share a file from Alice to Bob
    shareFile(user1, 1, user2)

    // Display shared files for Bob
    fmt.Println("Shared files for Bob:")
    listSharedFiles(user2)
}

// Function to upload a file
func uploadFile(user *User, filePath string) {  
    // Read the file content from the file system
    fileContent, err := os.ReadFile(filePath)
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

    // Generate a new file ID
    newFileID := len(files) + 1

    // Create a new File struct and add it to the global files slice
    newFile := &File{
        ID:       newFileID,
        Name:     filepath.Base(filePath),
        Content:  fileContent,
        OwnerID:   user.ID,
    }
    files = append(files, newFile)

    // Add the file pointer to the user's Files slice
    user.Files = append(user.Files, newFile)

    fmt.Printf("File '%s' uploaded by user %s (ID: %d)\n", newFile.Name, user.Name, user.ID)
}

// Function to share a file between users
func shareFile(owner *User, fileID int, receiver *User) {  
    for _, file := range owner.Files {
        if file.ID == fileID {
            // Found the file to share. Add it to the receiver's Files slice
            receiver.Files = append(receiver.Files, file)
            fmt.Printf("File '%s' shared from user %s (ID: %d) to user %s (ID: %d)\n", file.Name, owner.Name, owner.ID, receiver.Name, receiver.ID)
            return
        }
    }

    fmt.Printf("User %s (ID: %d) does not own file with ID: %d\n", owner.Name, owner.ID, fileID)
}

// Function to list shared files for a given user
func listSharedFiles(user *User) {  
    for _, file := range user.Files {
        if file.OwnerID != user.ID {
            fmt.Printf("File ID: %d, Name: %s, Owner: %s (ID: %d)\n", file.ID, file.Name, getUserNameByID(file.OwnerID), file.OwnerID)
        }
    }
}

// Helper function to get user name by ID
func getUserNameByID(userID int) string {  
    for _, user := range users {
        if user.ID == userID {
            return user.Name
        }
    }
    return "Unknown User"
}