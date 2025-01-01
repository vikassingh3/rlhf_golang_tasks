package main  
import (  
    "fmt"
    "net/http"
    "sync"
    "time"
)

var (  
    // Global variable to hold the client
    client *http.Client
)

func init() {
    client = &http.Client{
        Timeout: time.Second * 5,
        Transport: &http.Transport{
            MaxIdleConnsPerHost: 100, // Adjust this based on your system capacity
        },
    }
}

func makeRequest(wg *sync.WaitGroup, url string) {
    defer wg.Done()
    resp, err := client.Get(url)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer resp.Body.Close()
    // Do some processing with the response
}

func main() {
    var wg sync.WaitGroup
    urls := []string{"http://example1.com", "http://example2.com", "http://example3.com"}

    wg.Add(len(urls))
    for _, url := range urls {
        go makeRequest(&wg, url)
    }
    wg.Wait()
}