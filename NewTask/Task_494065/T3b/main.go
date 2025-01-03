package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

func CreateAndReturnSlice() []int {
    slice := []int{1, 2, 3, 4, 5}
    // Unnecessary loop to add time consumption
    for i := 0; i < 100000000; i++ {
        slice = append(slice, i)
    }
    return slice
}

func main() {
    f, err := os.Create("profile.out")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    if err := pprof.StartCPUProfile(f); err != nil {
        log.Fatal(err)
    }
    defer pprof.StopCPUProfile()

    numbers := CreateAndReturnSlice()
    fmt.Println("The slice returned has", len(numbers), "elements")
}

// func CreateAndReturnSlice() []int {
//     sliceSize := 100000005
//     slice := make([]int, sliceSize)
//     for i := 0; i < sliceSize; i++ {
//         slice[i] = i
//     }
//     return slice
// }