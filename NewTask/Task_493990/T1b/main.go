
package main  
import (  
    "encoding/csv"  
    "encoding/json"  
    "encoding/xml"  
    "fmt"  
    "io/ioutil"  
    "log"  
    "os"  
)  
// JSON data structure  
type Person struct {  
    Name string `json:"name"`  
    Age  int    `json:"age"`  
}  
// XML data structure  
type Book struct {  
    Title   string `xml:"title"`  
    Author  string `xml:"author"`  
    Genre   string `xml:"genre"`  
    Price   float64 `xml:"price"`  
    PublishDate string `xml:"publish_date"`  
}  
func main() {  
    // Parse JSON file  
    parseJSON("./data.json")  
    // Parse XML file  
    parseXML("./data.xml")  
    // Parse CSV file  
    parseCSV("./data.csv")  
}  
func parseJSON(filePath string) {  
    file, err := os.Open(filePath)  
    if err != nil {  
        log.Fatalf("Error opening JSON file: %v", err)  
    }  
    defer file.Close()  
    data, err := ioutil.ReadAll(file)  
    if err != nil {  
        log.Fatalf("Error reading JSON file: %v", err)  
    }  
    var person Person  
    if err := json.Unmarshal(data, &person); err != nil {  
        log.Fatalf("Error parsing JSON: %v", err)  
    }  
    fmt.Println("JSON Parsing Result:")  
    fmt.Println("Name:", person.Name)  
    fmt.Println("Age:", person.Age)  
}  
func parseXML(filePath string) {  
    file, err := os.Open(filePath)  
    if err != nil {  
        log.Fatalf("Error opening XML file: %v", err)  
    }  
    defer file.Close()  
    data, err := ioutil.ReadAll(file)  
    if err != nil {  
        log.Fatalf("Error reading XML file: %v", err)  
    }  
    var book Book  
    if err := xml.Unmarshal(data, &book); err != nil {  
        log.Fatalf("Error parsing XML: %v", err)  
    }  
    fmt.Println("\nXML Parsing Result:")  
    fmt.Println("Title:", book.Title)  
    fmt.Println("Author:", book.Author)  
    fmt.Println("Genre:", book.Genre)  
    fmt.Println("Price:", book.Price)  
    fmt.Println("Publish Date:", book.PublishDate)  
}  
func parseCSV(filePath string) {  
    file, err := os.Open(filePath)  
    if err != nil {  
        log.Fatalf("Error opening CSV file: %v", err)  
    }  
    defer file.Close()  
    reader := csv.NewReader(file)  
    records, err := reader.ReadAll()  
    if err != nil {  
        log.Fatalf("Error reading CSV file: %v", err)  
    }  
    fmt.Println("\nCSV Parsing Result:")  
    for _, record := range records {  
        fmt.Println(record)  
    }  
}  
