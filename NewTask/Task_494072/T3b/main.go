package main

import (
	"encoding/binary"
	"fmt"
	"hash/crc32"
	"sync"
)

// DataMigrator defines the methods required for data migration.
type DataMigrator interface {
	ReadData(chunkSize int) ([]byte, error)
	WriteData(data []byte) error
}

// InMemoryStorage is an in-memory implementation of DataMigrator.
type InMemoryStorage struct {
	data []byte
	mu   sync.Mutex
}

// ReadData reads a chunk of data from the in-memory storage.
func (s *InMemoryStorage) ReadData(chunkSize int) ([]byte, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.data) == 0 {
		return nil, fmt.Errorf("EOF") // No more data to read
	}

	// Ensure chunkSize does not exceed available data
	if chunkSize > len(s.data) {
		chunkSize = len(s.data)
	}

	chunk := s.data[:chunkSize]
	s.data = s.data[chunkSize:] // Remove the chunk from storage
	return chunk, nil
}

// WriteData writes data to the in-memory storage.
func (s *InMemoryStorage) WriteData(data []byte) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data = append(s.data, data...)
	return nil
}

// CalculateChecksum computes a checksum for a given byte slice.
func calculateChecksum(data []byte) uint32 {
	return crc32.ChecksumIEEE(data)
}

// MigrateData migrates data from a source to a destination with chunking and checksum validation.
func MigrateData(source DataMigrator, destination DataMigrator, chunkSize int) error {
	for {
		// Read chunk from source
		chunk, err := source.ReadData(chunkSize)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return fmt.Errorf("error reading data: %w", err)
		}

		// Calculate checksum
		checksum := calculateChecksum(chunk)

		// Append checksum to the chunk
		checksumBytes := make([]byte, 4)
		binary.LittleEndian.PutUint32(checksumBytes, checksum)
		chunkWithChecksum := append(chunk, checksumBytes...)

		// Write chunk with checksum to destination
		err = destination.WriteData(chunkWithChecksum)
		if err != nil {
			return fmt.Errorf("error writing data: %w", err)
		}
	}

	// Validate data integrity during reading from destination
	for {
		// Read chunk with checksum from destination
		chunkWithChecksum, err := destination.ReadData(chunkSize + 4)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return fmt.Errorf("error reading data: %w", err)
		}

		// Separate data and checksum
		receivedChecksum := binary.LittleEndian.Uint32(chunkWithChecksum[len(chunkWithChecksum)-4:])
		data := chunkWithChecksum[:len(chunkWithChecksum)-4]

		// Validate checksum
		calculatedChecksum := calculateChecksum(data)
		if receivedChecksum != calculatedChecksum {
			return fmt.Errorf("data corruption detected in chunk")
		}
	}

	return nil
}

func main() {
	// Source with initial data
	source := &InMemoryStorage{
		data: []byte("Hello, Go! This is a test of data migration."),
	}

	// Destination with no data initially
	destination := &InMemoryStorage{}

	// Migrate data with chunk size 10 bytes
	err := MigrateData(source, destination, 10)
	if err != nil {
		fmt.Println("Migration failed:", err)
	} else {
		fmt.Println("Migration successful. Destination data:", string(destination.data))
	}
}
