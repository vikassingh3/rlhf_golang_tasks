package main

import (
    "fmt"
)

type Result struct {
    Quotient, Remainder float64
}

func calculate(x, y float64) Result {
    quotient := x / y
    remainder := x - quotient*y
    return Result{quotient, remainder}
}

func main() {
    x := 10.0
    y := 3.0
    result := calculate(x, y)
    fmt.Printf("Quotient: %.2f, Remainder: %.2f\n", result.Quotient, result.Remainder)
}
