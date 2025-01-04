package main  
import (  
    "encoding/json"  
    "fmt"  
    "net/http"  
)  
// APIResponse is a struct that represents the response from the API.  
type APIResponse struct {  
    Message string `json:"message"`  
}  
// getJSON makes a GET request to the specified URL and decodes the response into the target struct.  
func getJSON(url string, target interface{}) error {  
    // Make the GET request.  
    resp, err := http.Get(url)  
    if err != nil {  
        return err  
    }  
    defer resp.Body.Close()  
// Check the status code.  
    if resp.StatusCode != http.StatusOK {  
        return fmt.Errorf("got status code %d", resp.StatusCode)  
    }  
// Decode the response into the target struct.  
    if err := json.NewDecoder(resp.Body).Decode(target); err != nil {  
        return err  
    }  
    return nil  
}  
func main() {  
    var response APIResponse  
    // Make the GET request and decode the response.  
    err := getJSON("https://jsonplaceholder.typicode.com/posts/1", &response)  
    if err != nil {  
        fmt.Printf("Error: %v\n", err)  
        return  
    }  
    // Print the decoded response.  
    fmt.Println(response)  
}  