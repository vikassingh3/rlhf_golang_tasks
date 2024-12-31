package main  
import (  
    "fmt"
    "math/rand"
    "sync"
    "time"
)

// FashionItem represents a fashion item in the dataset
type FashionItem struct {
    Color  string
    Style  string
    // Add other relevant attributes here
}

// Function to generate a slice of FashionItem structs (dataset)
func generateDataset() []FashionItem {
    rand.Seed(time.Now().UnixNano())
    colors := []string{"red", "blue", "green", "yellow", "black", "white"}
    styles := []string{"dress", "top", "bottom", "jumpers", "sunglasses", "hat"}

    dataset := make([]FashionItem, 10000) // Generate a large dataset
    for i := range dataset {
        dataset[i] = FashionItem{
            Color: colors[rand.Intn(len(colors))],
            Style: styles[rand.Intn(len(styles))],
        }
    }
    return dataset
}

// Function to calculate trending colors 
func analyzeTrendingColors(dataset []FashionItem, wg *sync.WaitGroup) {
    defer wg.Done()
    colorCount := make(map[string]int)
    for _, item := range dataset {
        colorCount[item.Color]++
    }
    // Find the most common color (trending color)
    var mostCommonColor string
    var maxCount int
    for color, count := range colorCount {
        if count > maxCount {
            mostCommonColor = color
            maxCount = count
        }
    }
    fmt.Printf("Trending color: %s (Count: %d)\n", mostCommonColor, maxCount)
}

// Function to calculate popular styles
func analyzePopularStyles(dataset []FashionItem, wg *sync.WaitGroup) {
    defer wg.Done()
    styleCount := make(map[string]int)
    for _, item := range dataset {
        styleCount[item.Style]++
    }

    // Find the top 3 popular styles
    popularStyles := make([]string, 3)
    var counts [3]int
    for style, count := range styleCount {
        for i := 2; i >= 0; i-- {
            if count > counts[i] {
                popularStyles[i+1] = popularStyles[i]
                counts[i+1] = counts[i]
                popularStyles[i] = style
                counts[i] = count
                break
            }
        }
    }
    fmt.Println("Popular styles:")
    for i, style := range popularStyles {
        fmt.Printf("%d. %s (Count: %d)\n", i+1, style, counts[i])
    }
}  

func main() {  
    dataset := generateDataset() // Generate the dataset

    // Create a WaitGroup to wait for all goroutines to complete
    var wg sync.WaitGroup

    // Process trending colors and popular styles in parallel
    wg.Add(1)
    go analyzeTrendingColors(dataset, &wg)
    wg.Add(1)
    go analyzePopularStyles(dataset, &wg)

    // Wait for all goroutines to finish
    wg.Wait()

    fmt.Println("\nAnalysis completed successfully!")
} 