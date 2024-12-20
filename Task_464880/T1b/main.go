
package main  
import (  
   "fmt"
   "time"
)
func main() {  
   //Example 1: Performance with range loop
   start := time.Now()
   slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
   sum := 0
   for _, value := range slice1 {
       sum += value
   }
   elapsed := time.Since(start)
   fmt.Printf("Sum using range loop: %d\n", sum)
   fmt.Printf("Time taken: %s\n", elapsed)
   
   //Example 2: Performance with for loop
   start = time.Now()
   slice2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
   sum = 0
   for i := 0; i < len(slice2); i++ {
       sum += slice2[i]
   }
   elapsed = time.Since(start)
   fmt.Printf("Sum using for loop: %d\n", sum)
   fmt.Printf("Time taken: %s\n", elapsed)
 }
 