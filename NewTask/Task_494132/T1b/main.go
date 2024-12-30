package main  
import (  
    "fmt"
    "reflect"
)  

// A function that returns different data types based on the input `i`
func GetType(i int) interface{} {  
    switch i {  
    case 0:  
        return "string"
    case 1:  
        return 42
    case 2:  
        return 3.14
    default:  
        return nil
    }  
}  

func main() {  
    for i := 0; i < 4; i++ {  
        value := GetType(i)
        // Use reflection to get the value's type and kind
        v := reflect.ValueOf(value)
        k := v.Kind()

        switch k {
        case reflect.String:
            fmt.Println("Type: string, Value:", v.String())
        case reflect.Int:
            fmt.Println("Type: int, Value:", v.Int())
        case reflect.Float64:
            fmt.Println("Type: float64, Value:", v.Float())
        default:
            fmt.Println("Type: nil")
        }
    }
}  