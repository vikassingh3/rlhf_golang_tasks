package main

import (
	"fmt"
	"io/ioutil"
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
    filename string  
}  
func (s *FileStorage) ReadData() ([]byte, error) {  
    // Implement file reading logic here  
    data, err := ioutil.ReadFile(s.filename)  
    if err != nil {  
        return nil, err  
    }  
    return data, nil  
}  
func (s *FileStorage) WriteData(data []byte) error {  
    // Implement file writing logic here  
    err := ioutil.WriteFile(s.filename, data, 0644)  
    if err != nil {  
        return err  
    }  
    return nil  
}  
func MigrateData(source DataMigrator, destination DataMigrator) error {  
    for {  
        chunk, err := source.ReadData()  
        if err != nil {  
            return err  
        }  
        if len(chunk) == 0 {  
            break  
        }  
        err = destination.WriteData(chunk)  
        if err != nil {  
            return err  
        }  
    }  
    return nil  
}  
func main() {  
    // Example usage:  
    inMemorySource := &InMemoryStorage{data: []byte("Hello, World!")}  
    fileDestination := &FileStorage{filename: "output.txt"}  
    err := MigrateData(inMemorySource, fileDestination)  
    if err != nil {  
        fmt.Println("Error migrating data:", err)  
    } else {  
        fmt.Println("Data migration successful.")  
    }  
}  