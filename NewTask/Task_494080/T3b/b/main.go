package main
import (
	"fmt"
	"time"
)

func dataProducer(dataCh chan<- int) {
	defer close(dataCh)
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * 1) // Simulate data arrival rate
		dataCh <- i
	}
}
func dataConsumer(dataCh <-chan int, processedCh chan<- string) {
	for data := range dataCh {
		processedData := fmt.Sprintf("Processed data: %d", data)
		processedCh <- processedData
	}
	close(processedCh)
}
func main() {
	dataCh := make(chan int)
	processedCh := make(chan string)

	go dataProducer(dataCh)
	go dataConsumer(dataCh, processedCh)

	// Print the processed data as it becomes available
	for processedData := range processedCh {
		fmt.Println(processedData)
	}
	fmt.Println("All data processed.")
}