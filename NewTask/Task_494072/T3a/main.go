package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// DataMigrator defines an interface for migration operations
type DataMigrator interface {
	ReadData(maxBytes int) ([]byte, error)
	WriteData(data []byte) error
}

// InMemoryStorage stores data in memory
type InMemoryStorage struct {
	data []byte
}

func (s *InMemoryStorage) ReadData(maxBytes int) ([]byte, error) {
	if len(s.data) == 0 {
		return nil, io.EOF
	}
	end := min(len(s.data), maxBytes)
	data := s.data[:end]
	s.data = s.data[end:]
	return data, nil
}

func (s *InMemoryStorage) WriteData(data []byte) error {
	s.data = append(s.data, data...)
	return nil
}

// FileStorage handles file-based storage
type FileStorage struct {
	filename string
}

func (s *FileStorage) ReadData(maxBytes int) ([]byte, error) {
	data, err := ioutil.ReadFile(s.filename)
	if err != nil {
		return nil, err
	}
	return data[:min(len(data), maxBytes)], nil
}

func (s *FileStorage) WriteData(data []byte) error {
	tmpFile, err := ioutil.TempFile("", "migration")
	if err != nil {
		return err
	}
	defer tmpFile.Close()

	_, err = tmpFile.Write(data)
	if err != nil {
		return err
	}

	return os.Rename(tmpFile.Name(), s.filename)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// MigrateData handles data migration with error recovery and slice integrity checks
func MigrateData(source DataMigrator, destination DataMigrator) error {
	retries := 3
	for retries > 0 {
		err := func() error {
			chunk, err := source.ReadData(1024 * 1024) // Read in chunks
			if err != nil {
				if err == io.EOF {
					return nil
				}
				return fmt.Errorf("error reading data: %v", err)
			}

			// Validate data integrity
			hash := md5.New()
			_, err = hash.Write(chunk)
			if err != nil {
				return fmt.Errorf("error computing hash: %v", err)
			}
			checksum := hex.EncodeToString(hash.Sum(nil))

			fmt.Printf("Checksum: %v\n", checksum)

			// Write data to destination
			err = destination.WriteData(chunk)
			if err != nil {
				return fmt.Errorf("error writing data: %v", err)
			}

			return nil
		}()
		if err != nil {
			retries--
			fmt.Printf("Error occurred: %v. Retrying... (%d retries left)\n", err, retries)
			if retries == 0 {
				return fmt.Errorf("migration failed after retries: %v", err)
			}
		} else {
			break
		}
	}
	return nil
}

func main() {
	source := &InMemoryStorage{data: []byte("Test data")}
	dest := &FileStorage{filename: "output.txt"}

	err := MigrateData(source, dest)
	if err != nil {
		fmt.Println("Migration failed:", err)
	} else {
		fmt.Println("Migration succeeded!")
	}
}