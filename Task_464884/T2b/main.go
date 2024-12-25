package main

import (
	"context"
	"io"
	"os"
	"sync"
	"time"
)

// Define a struct to store metadata
type FileMetadata struct {
	StartTime     time.Time
	EndTime       time.Time
	BytesRead     int64
	BytesWritten  int64
	Size          int64
	Status        string
}

// A goroutine-safe interface to update file metadata
type MetadataUpdater interface {
	IncrementBytesRead(int64)
	IncrementBytesWritten(int64)
	SetStatus(string)
	Finish(error)
}

type fileMetadataUpdater struct {
	metadata *FileMetadata
	mux      sync.Mutex
}

func (u *fileMetadataUpdater) IncrementBytesRead(n int64) {
	u.mux.Lock()
	defer u.mux.Unlock()
	u.metadata.BytesRead += n
}

func (u *fileMetadataUpdater) IncrementBytesWritten(n int64) {
	u.mux.Lock()
	defer u.mux.Unlock()
	u.metadata.BytesWritten += n
}

func (u *fileMetadataUpdater) SetStatus(status string) {
	u.mux.Lock()
	defer u.mux.Unlock()
	u.metadata.Status = status
}

func (u *fileMetadataUpdater) Finish(err error) {
	u.mux.Lock()
	defer u.mux.Unlock()
	u.metadata.EndTime = time.Now()
	if err != nil {
		u.metadata.Status = err.Error()
	}
}

// Create a new MetadataUpdater
func newMetadataUpdater(metadata *FileMetadata) MetadataUpdater {
	return &fileMetadataUpdater{metadata: metadata}
}

// A wrapper type for a custom writer function
type WriterFunc func(p []byte) (n int, err error)

// Implement the Write method for WriterFunc
func (f WriterFunc) Write(p []byte) (n int, err error) {
	return f(p)
}

func readFileWithMetadata(ctx context.Context, filename string) (data []byte, metadata FileMetadata, err error) {
	metadata.StartTime = time.Now()
	updater := newMetadataUpdater(&metadata)

	// Open the file for reading
	file, err := os.Open(filename)
	if err != nil {
		updater.Finish(err)
		return nil, metadata, err
	}
	defer file.Close()

	// Read the file content in chunks and update metadata
	reader := io.TeeReader(file, WriterFunc(func(p []byte) (n int, err error) {
		n = len(p)
		updater.IncrementBytesRead(int64(n))
		return n, nil
	}))
	data, err = io.ReadAll(reader)
	if err != nil {
		updater.Finish(err)
		return nil, metadata, err
	}

	// Retrieve file information and update size
	fileInfo, err := file.Stat()
	if err != nil {
		updater.Finish(err)
		return nil, metadata, err
	}

	metadata.Size = fileInfo.Size()
	updater.SetStatus("success")
	updater.Finish(nil)

	return data, metadata, nil
}

func writeFileWithMetadata(ctx context.Context, filename string, data []byte) (metadata FileMetadata, err error) {
	metadata.StartTime = time.Now()
	updater := newMetadataUpdater(&metadata)

	// Open the file for writing
	file, err := os.Create(filename)
	if err != nil {
		updater.Finish(err)
		return metadata, err
	}
	defer file.Close()

	// Write the file content in chunks and update metadata
	writer := io.MultiWriter(file, WriterFunc(func(p []byte) (n int, err error) {
		n = len(p)
		updater.IncrementBytesWritten(int64(n))
		return n, nil
	}))
	_, err = writer.Write(data)
	if err != nil {
		updater.Finish(err)
		return metadata, err
	}

	metadata.Size = int64(len(data))
	updater.SetStatus("success")
	updater.Finish(nil)

	return metadata, nil
}

func main() {
	// Example usage
	ctx := context.Background()
	filename := "example.txt"
	content := []byte("Hello, Go!")

	// Write file
	metadata, err := writeFileWithMetadata(ctx, filename, content)
	if err != nil {
		panic(err)
	}
	println("Write Metadata:", metadata.BytesWritten, "bytes written.")

	// Read file
	data, metadata, err := readFileWithMetadata(ctx, filename)
	if err != nil {
		panic(err)
	}
	println("Read Metadata:", metadata.BytesRead, "bytes read.")
	println("File Content:", string(data))
}
