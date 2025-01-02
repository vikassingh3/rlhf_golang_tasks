package main  
import (  
   "fmt"
   "sync"  
   "time"  
)  
const numElements = 10000000  
const numGoroutines = 8  
const paddingSize = 64 // Cache line size (on most platforms)

func main() {  
   var wg sync.WaitGroup  
   wg.Add(numGoroutines)  
   type PaddedInt struct {  
       value int32  
       padding [paddingSize - 4]byte  
   }  
   data := make([]PaddedInt, numElements)  
   // ... (Code removed for brevity)  
   start := time.Now()  
   for i := 0; i < numGoroutines; i++ {  
       go func(index int) {  
           defer wg.Done()  
           for j := index; j < numElements; j += numGoroutines {  
               data[j].value += 1  
           }  
       }(i)  
   }  
   wg.Wait()  
   elapsed := time.Since(start)  
   fmt.Println("Time taken:", elapsed)  
}  