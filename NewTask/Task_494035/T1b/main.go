package main  
import (  
    "fmt"
    "reflect"
)  
// User struct  
type User struct {  
    Name string `json:"name"`
    Age  int    `json:"age"`
    City string `json:"city"`
}  

// main function  
func main() {  
    user := User{Name: "John Doe", Age: 30, City: "New York"}  
    formatStruct(user)  
}  

// formatStruct function  
func formatStruct(s interface{}) {  
    v := reflect.ValueOf(s)  
    t := v.Type()  
    format := ""  
    values := make([]interface{}, 0)  
    for i := 0; i < v.NumField(); i++ {  
        field := t.Field(i)  
        value := v.Field(i).Interface()  
        format += fmt.Sprintf("%s: %%v\n", field.Name)  
        values = append(values, value)  
    }  
    fmt.Printf(format, values...)  
}