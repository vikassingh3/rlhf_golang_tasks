package main  
import (  
    "encoding/json"  
    "fmt"  
    "net/http"  
)  
// APIResponse is a struct that represents the response from the API.  
type APIResponse struct {  
    Message string `json:"message,omitempty"`  
    ID      int    `json:"id,omitempty"`  
}  
func getJSON(url string, target interface{}) error {  
    resp, err := http.Get(url)  
    if err != nil {  
        return err  
    }  
    defer resp.Body.Close()  
    if resp.StatusCode != http.StatusOK {  
        return fmt.Errorf("got status code %d", resp.StatusCode)  
    }  
    if err := json.NewDecoder(resp.Body).Decode(target); err != nil {  
        return err  
    }  
    return nil  
}  
func main() {  
    var response APIResponse  
    // The JSON data returned by the API might have missing fields.  
    err := getJSON("https://example.com/api/data", &response)  
    if err != nil {  
        fmt.Printf("Error: %v\n", err)  
        return  
    }  
    fmt.Println("Decoded Response:")  
    fmt.Printf("Message: %s\n", response.Message)  
    fmt.Printf("ID: %d\n", response.ID)  
}  