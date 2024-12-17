// data_processing.go
package main

import (
	// "errors"
	"fmt"
	"log"

	"github.com/pkg/errors"
)

// ProcessData processes data chunks and returns processed data or an error
func ProcessData(chunk []byte) ([]byte, error) {
	if len(chunk) == 0 {
		return nil, errors.New("empty data chunk")
	}

	processed, err := transformData(chunk)
	if err != nil {
		return nil, errors.Wrap(err, "failed to transform data")
	}

	return processed, nil
}

func transformData(data []byte) ([]byte, error) {
	for i := range data {
		if data[i] == 'e' {
			return nil, errors.New("data contains prohibited character 'e'")
		}
	}
	return data, nil
}

func main() {
	chunk := []byte("hello")
	processed, err := ProcessData(chunk)
	if err != nil {
		log.Fatalf("processing failed: %v", err)
	}
	fmt.Println(string(processed))

	chunk = []byte("invalid")
	processed, err = ProcessData(chunk)
	if err != nil {
		log.Fatalf("processing failed: %v", err)
	}
	fmt.Println(string(processed))
}
