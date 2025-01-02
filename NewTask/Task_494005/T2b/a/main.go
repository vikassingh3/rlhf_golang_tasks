package main  
import (  
  "fmt"
  "sync"
  "time"
)  
const numElements = 10000000
const numGoroutines = 8

func main() {  
  var wg sync.WaitGroup  
  wg.Add(numGoroutines)

  data := make([]int, numElements)  
  // Fill the data slice with random values
  // ... (Code removed for brevity)

  start := time.Now()

  for i := 0; i < numGoroutines; i++ {  
      go func(index int) {  
          defer wg.Done()  
          for j := index; j < numElements; j += numGoroutines {  
              // Sequential access to adjacent memory locations
              data[j] += 1  
          }  
      }(i)  
  }  
  wg.Wait()  
  elapsed := time.Since(start)  
  fmt.Println("Time taken:", elapsed)  
}  