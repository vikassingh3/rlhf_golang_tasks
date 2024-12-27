package main

import (
    "fmt"
    "sync"
)

var counter int
var wg = &sync.WaitGroup{}

func increment() {
    for i := 0; i < 100000; i++ {
        counter++
    }
    wg.Done()
}

func main() {
    wg.Add(10)
    for i := 0; i < 10; i++ {
        go increment()
    }
    wg.Wait()
    fmt.Println("Final counter value:", counter)
}