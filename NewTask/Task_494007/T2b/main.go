package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"runtime"
	"runtime/pprof"
)

func main() {  
    runtime.GOMAXPROCS(runtime.NumCPU())

    // Start the CPU profile
    f, err := os.Create("cpu.prof")
    if err != nil {
        log.Fatal(err)
    }
    pprof.StartCPUProfile(f)
    defer pprof.StopCPUProfile()

    // Your program's main code to be profiled
    expensiveFunction()
}

func expensiveFunction() {  
    sum := 0.0
    for i := 0; i < 10000000; i++ {  
        sum += math.Sqrt(float64(i))
    }  
    fmt.Println("Sum:", sum)
}  


// package main  
// import (  
//     "fmt"
//     "os"
//     "runtime"
//     "runtime/pprof"
// )  
// func main() {  
//     f, err := os.Create("heap.prof")
//     if err != nil {
//         log.Fatal(err)
//     }
//     defer f.Close()

//     pprof.WriteHeapProfile(f)

//     // Rest of the program
//     expensiveFunction()
// }