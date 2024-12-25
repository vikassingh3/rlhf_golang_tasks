package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

// define a struct to store metadata
type FileMetadata struct {
    Size int64
    Status string
}

func readFileWithMetadata(ctx context.Context, filename string) (data []byte, metadata FileMetadata, err error) {
    // create a child context with a cancel function
    ctx, cancel := context.WithCancel(ctx)
    defer cancel() // ensure the cancel function is called even if an error occurs

    // open the file for reading
    file, err := os.Open(filename)
    if err != nil {
        return nil, FileMetadata{}, err
    }
    defer file.Close()

    // read the file content
    data, err = ioutil.ReadAll(file)
    if err != nil {
        return nil, FileMetadata{}, err
    }

    // retrieve file information
    fileInfo, err := file.Stat()
    if err != nil {
        return nil, FileMetadata{}, err
    }

    // update the metadata struct
    metadata = FileMetadata{
        Size:   fileInfo.Size(),
        Status: "success",
    }

    return data, metadata, nil
}

func writeFileWithMetadata(ctx context.Context, filename string, data []byte) (metadata FileMetadata, err error) {
    // create a child context with a cancel function
    ctx, cancel := context.WithCancel(ctx)
    defer cancel() // ensure the cancel function is called even if an error occurs

    // open the file for writing
    file, err := os.Create(filename)
    if err != nil {
        return FileMetadata{}, err
    }
    defer file.Close()

    // write the file content
    _, err = file.Write(data)
    if err != nil {
        return FileMetadata{}, err
    }

    // retrieve file information
    fileInfo, err := file.Stat()
    if err != nil {
        return FileMetadata{}, err
    }

    // update the metadata struct
    metadata = FileMetadata{
        Size:   fileInfo.Size(),
        Status: "success",
    }

    return metadata, nil
}

func main() {  
    // create a context with a timeout
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel() // ensure the cancel function is called even if the main function exits

    // perform file reading
    data, metadata, err := readFileWithMetadata(ctx, "example.txt")
    if err != nil {
        fmt.Printf("Error reading file: %v\n", err)
        return
    }

    fmt.Printf("File Content: %s\n", string(data))
    fmt.Printf("File Size: %d bytes\n", metadata.Size)
    fmt.Printf("File Operation Status: %s\n", metadata.Status)

    // perform file writing
    newData := []byte("Hello, world!")
    metadata, err = writeFileWithMetadata(ctx, "example_output.txt", newData)
    if err != nil {
        fmt.Printf("Error writing file: %v\n", err)
        return
    }

    fmt.Printf("File Size: %d bytes\n", metadata.Size)
    fmt.Printf("File Operation Status: %s\n", metadata.Status)
}
