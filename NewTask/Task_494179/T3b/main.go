package main

import (
    "fmt"
)

func dividePtr(x, y float64, quotient, remainder *float64) {
    *quotient = x / y
    *remainder = x - *quotient*y
}

func main() {
    x := 10.0
    y := 3.0
    var quotient, remainder float64

    dividePtr(x, y, &quotient, &remainder)

    fmt.Printf("Quotient: %.2f, Remainder: %.2f\n", quotient, remainder)
}
