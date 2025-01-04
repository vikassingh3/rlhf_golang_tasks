package main

import (
	"bufio"
	"log"
	"os"
	"testing"
)

func BenchmarkStandardLogging(b *testing.B) {
	file, err := os.OpenFile("standard_log.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		b.Fatal(err)
	}
	defer file.Close()

	for i := 0; i < b.N; i++ {
		log.SetOutput(file)
		log.Println("Standard log message")
	}
}

func BenchmarkBufferedLogging(b *testing.B) {
	file, err := os.OpenFile("buffered_log.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		b.Fatal(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	for i := 0; i < b.N; i++ {
		_, err = writer.WriteString("Buffered log message\n")
		if err != nil {
			b.Fatal(err)
		}
		if err := writer.Flush(); err != nil {
			b.Fatal(err)
		}
	}
}
