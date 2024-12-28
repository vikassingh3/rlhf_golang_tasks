package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/xuri/excelize/v2"
)

type ExcelRecord struct {
	Name     string
	Age      int
	Email    string
	PhoneNum string
}

func parseExcel(filePath string) {
	file, err := excelize.OpenFile(filePath)
	if err != nil {
		log.Fatalf("Error opening Excel file: %v", err)
	}

	// Get all rows from the "Sheet1"
	rows, err := file.GetRows("Sheet1")
	if err != nil {
		log.Fatalf("Error reading rows from Excel file: %v", err)
	}

	for rowIndex, row := range rows {
		// Skip the header row
		if rowIndex == 0 {
			continue
		}

		record := ExcelRecord{}
		for colIndex, cell := range row {
			switch colIndex {
			case 0:
				record.Name = cell
			case 1:
				if cell == "" {
					record.Age = 0
				} else {
					record.Age, _ = strconv.Atoi(cell)
				}
			case 2:
				record.Email = cell
			case 3:
				record.PhoneNum = cell
			}
		}

		fmt.Println("Excel Parsing Result:")
		fmt.Printf("Name: %s, Age: %d, Email: %s, Phone Num: %s\n", record.Name, record.Age, record.Email, record.PhoneNum)
	}
}

func main() {
	parseExcel("./people.zip")
}
