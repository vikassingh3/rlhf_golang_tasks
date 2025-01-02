package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

type DataMigrator interface {
	ReadData() ([]byte, error)
	WriteData(data []byte) error
}

type InMemoryStorage struct {
	data []byte
}

func (s *InMemoryStorage) ReadData() ([]byte, error) {
	return s.data, nil
}

func (s *InMemoryStorage) WriteData(data []byte) error {
	s.data = append(s.data, data...)
	return nil
}

type FileStorage struct {
	filename    string
	tempFilename string // Used for temporary file during migration
}

func (s *FileStorage) ReadData() ([]byte, error) {
	data, err := ioutil.ReadFile(s.filename)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *FileStorage) WriteData(data []byte) error {
	err := ioutil.WriteFile(s.tempFilename, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (s *FileStorage) Commit() error {
	// Replace the original file with the temp file
	err := os.Rename(s.tempFilename, s.filename)
	if err != nil {
		return err
	}
	return nil
}

func (s *FileStorage) Rollback() error {
	// Delete the temp file in case of rollback
	err := os.Remove(s.tempFilename)
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	return nil
}

type Transaction interface {
	Commit() error
	Rollback() error
}

type inMemoryTransaction struct {
	storage *InMemoryStorage
	data    []byte
}

func (t *inMemoryTransaction) Commit() error {
	t.storage.data = t.data
	return nil
}

func (t *inMemoryTransaction) Rollback() error {
	// Nothing to roll back in in-memory storage
	return nil
}

func startTransaction(storage DataMigrator) (Transaction, error) {
	switch s := storage.(type) {
	case *InMemoryStorage:
		return &inMemoryTransaction{storage: s, data: s.data}, nil
	case *FileStorage:
		return s, nil
	default:
		return nil, errors.New("unsupported storage for transaction")
	}
}

func MigrateData(source DataMigrator, destination DataMigrator, chunkSize int) error {
	// Start a transaction on the destination storage
	tx, err := startTransaction(destination)
	if err != nil {
		return fmt.Errorf("error starting transaction: %w", err)
	}

	defer func() {
		if r := recover(); r != nil {
			// Rollback changes if a panic occurs
			tx.Rollback()
			panic(r)
		}
	}()

	// Perform migration
	err = migrateDataWithTransaction(source, destination, chunkSize)
	if err != nil {
		// Rollback changes if an error occurs
		tx.Rollback()
		return fmt.Errorf("error migrating data: %w", err)
	}

	// Commit the transaction if everything succeeds
	err = tx.Commit()
	if err != nil {
		// Rollback changes if an error occurs during commit
		tx.Rollback()
		return fmt.Errorf("error committing transaction: %w", err)
	}

	return nil
}

func migrateDataWithTransaction(source DataMigrator, destination DataMigrator, chunkSize int) error {
	data, err := source.ReadData()
	if err != nil {
		return fmt.Errorf("error reading data from source: %w", err)
	}

	for i := 0; i < len(data); i += chunkSize {
		end := i + chunkSize
		if end > len(data) {
			end = len(data)
		}
		chunk := data[i:end]
		err := destination.WriteData(chunk)
		if err != nil {
			return fmt.Errorf("error writing data to destination: %w", err)
		}
	}

	return nil
}

func main() {
	inMemorySource := &InMemoryStorage{data: []byte("This is some test data to migrate.")}
	fileDestination := &FileStorage{filename: "output.txt", tempFilename: "output_temp.txt"}

	chunkSize := 10
	err := MigrateData(inMemorySource, fileDestination, chunkSize)
	if err != nil {
		fmt.Println("Error migrating data:", err)
	} else {
		fmt.Println("Data migration successful.")
	}

	// Verify the written file
	data, err := ioutil.ReadFile("output.txt")
	if err == nil {
		fmt.Println("Written data:", string(data))
	} else {
		fmt.Println("Error reading migrated file:", err)
	}
}
