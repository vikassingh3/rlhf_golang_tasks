
package main  
import (  
    "fmt"
    "strconv"
    "bufio"
    "os"
    "strings"
)

func main() {  
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter your name: ")
    name, _ := reader.ReadString('\n')
    name = strings.TrimSpace(name)

    fmt.Print("Enter your age: ")
    ageStr, _ := reader.ReadString('\n')
    ageStr = strings.TrimSpace(ageStr)
    age, _ := strconv.Atoi(ageStr)

    fmt.Printf("Hello, my name is %s and I am %d years old.\n", name, age)
}  
  