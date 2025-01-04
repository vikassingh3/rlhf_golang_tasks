package main

import (
	"errors"
	"fmt"
	"strings"
)

type PipelineError struct {
	Stage string
	Err   error
}

func (p *PipelineError) Error() string {
	return fmt.Sprintf("Pipeline error at stage '%s': %v", p.Stage, p.Err)
}

// Stage1 is the first step in the pipeline.
func Stage1(data string) (string, error) {
	if strings.Contains(data, "!") {
		return data, errors.New("Unexpected '!' in data.")
	}
	return strings.TrimSpace(data), nil
}

// Stage2 is the second step in the pipeline.
func Stage2(data string) (string, error) {
	if data == "" {
		return data, errors.New("Data is empty after trimming.")
	}
	return data, nil
}

// Stage3 is the third step in the pipeline.
func Stage3(data string) (string, error) {
	if strings.Contains(data, "@") {
		return data, errors.New("Unexpected '@' in data.")
	}
	return data, nil
}

// ProcessData processes data through a pipeline with error handling.
func ProcessData(data string) (string, error) {
	var (
		err      error
		processed string = data
	)

	processed, err = Stage1(processed)
	if err != nil {
		return "", &PipelineError{"Stage1", err}
	}

	processed, err = Stage2(processed)
	if err != nil {
		return "", &PipelineError{"Stage2", err}
	}

	processed, err = Stage3(processed)
	if err != nil {
		return "", &PipelineError{"Stage3", err}
	}

	return processed, nil
}

func main() {
	data := "   hello! world @ "
	processed, err := ProcessData(data)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Processed data: %s\n", processed)
	}
}