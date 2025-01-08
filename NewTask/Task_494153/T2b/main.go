package main

import (
	"errors"
	"fmt"
)

func calculateRectangle(length, width float64) (float64, float64, error) {
	if length <= 0 || width <= 0 {
		return 0, 0, errors.New("length and width must be positive")
	}
	area := length * width
	perimeter := 2 * (length + width)
	return area, perimeter, nil
}

func main() {
	// Test the function using test cases
	testCases := []struct {
		length           float64
		width            float64
		expectedArea     float64
		expectedPerimeter float64
		expectedErr      string
	}{
		{
			length:           3.0,
			width:            4.0,
			expectedArea:     12.0,
			expectedPerimeter: 14.0,
			expectedErr:      "",
		},
		{
			length:           -2.0,
			width:            3.0,
			expectedArea:     0.0,
			expectedPerimeter: 0.0,
			expectedErr:      "length and width must be positive",
		},
	}

	for _, test := range testCases {
		area, perimeter, err := calculateRectangle(test.length, test.width)

		// Check if the error matches the expected error
		if err != nil {
			if err.Error() != test.expectedErr {
				fmt.Printf("Test failed for length=%f and width=%f: Expected error: %q, Got error: %q\n",
					test.length, test.width, test.expectedErr, err.Error())
			} else {
				fmt.Printf("Test passed for length=%f and width=%f: Got expected error: %q\n",
					test.length, test.width, err.Error())
			}
			continue
		}

		// If no error, check area and perimeter
		if area != test.expectedArea {
			fmt.Printf("Test failed for length=%f and width=%f: Expected area: %f, Got area: %f\n",
				test.length, test.width, test.expectedArea, area)
			continue
		}

		if perimeter != test.expectedPerimeter {
			fmt.Printf("Test failed for length=%f and width=%f: Expected perimeter: %f, Got perimeter: %f\n",
				test.length, test.width, test.expectedPerimeter, perimeter)
			continue
		}

		fmt.Printf("Test passed for length=%f and width=%f\n", test.length, test.width)
	}
}
