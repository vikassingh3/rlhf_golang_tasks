package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

// fileMetadata contains file related metadata
type fileMetadata struct {
	size    int64
	status  string
	err     error
	duration time.Duration
}

// readFileWithMetadata reads a file and tracks its size and status using a context.
func readFileWithMetadata(ctx context.Context, filePath string) (data []byte, metadata fileMetadata) {
	startTime := time.Now()
	defer func() {
		metadata.duration = time.Since(startTime)
	}()

	file, err := os.Open(filePath)
	if err != nil {
		return nil, fileMetadata{
			status:  fmt.Sprintf("Error opening file: %v", err),
			err:     err,
			duration: metadata.duration,
		}
	}
	defer file.Close()

	select {
	case <-ctx.Done():
		return nil, fileMetadata{
			status:  "Canceled",
			err:     ctx.Err(),
			duration: metadata.duration,
		}
	default:
	}

	buf := make([]byte, 1024)
	var readSize int64
	for {
		n, err := file.Read(buf)
		if n > 0 {
			readSize += int64(n)
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fileMetadata{
				status:  fmt.Sprintf("Error reading file: %v", err),
				err:     err,
				duration: metadata.duration,
			}
		}
		select {
		case <-ctx.Done():
			return nil, fileMetadata{
				status:  "Canceled",
				err:     ctx.Err(),
				duration: metadata.duration,
			}
		default:
		}
	}

	metadata.size = readSize
	metadata.status = "Success"
	return buf[:readSize], metadata
}

func main() {
	filePath := "example.txt"

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, metadata := readFileWithMetadata(ctx, filePath)

	if metadata.err != nil {
		log.Printf("Error reading file: %v", metadata.err)
		return
	}

	fmt.Printf("File read successfully. Size: %d bytes, Duration: %v\n", metadata.size, metadata.duration)
}