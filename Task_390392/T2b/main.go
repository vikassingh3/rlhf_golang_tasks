package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/tensorflow/tensorflow/lite/golang/interpreter"
)

// Global variable for anomaly threshold
var anomalyThreshold float64

// Load the TensorFlow Lite model
func loadModel(modelPath string) (*interpreter.Interpreter, error) {
	modelBuffer, err := os.ReadFile(modelPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read model file: %w", err)
	}

	interpreter := interpreter.New(modelBuffer)
	err = interpreter.AllocateTensors()
	if err != nil {
		return nil, fmt.Errorf("failed to allocate tensors: %w", err)
	}

	return interpreter, nil
}

// Preprocess data (normalize and reshape)
func preprocessData(rawData []byte) ([]float32, error) {
	data := make([]float32, len(rawData))
	for i, b := range rawData {
		data[i] = float32(b) / 255.0 // Normalize from [0, 255] to [0, 1]
	}

	// Example: Adjust shape if model requires specific dimensions
	return data, nil
}

// Perform inference on the model
func infer(model *interpreter.Interpreter, data []float32) (float32, error) {
	inputTensor := model.Input(0)
	outputTensor := model.Output(0)

	// Copy input data to the model's input tensor
	if err := inputTensor.Copy(data); err != nil {
		return 0, fmt.Errorf("failed to copy input data: %w", err)
	}

	// Run inference
	if err := model.Invoke(); err != nil {
		return 0, fmt.Errorf("failed to invoke model: %w", err)
	}

	// Read output data
	output := make([]float32, outputTensor.Shape()[0])
	if err := outputTensor.Read(output); err != nil {
		return 0, fmt.Errorf("failed to read output tensor: %w", err)
	}

	// Return anomaly score (assuming scalar output)
	return output[0], nil
}

// Consume messages from Kafka and evaluate anomalies
func consumeMessages(consumer *kafka.Consumer, model *interpreter.Interpreter) {
	defer consumer.Close()

	for {
		msg, err := consumer.ReadMessage(-1)
		if err != nil {
			log.Printf("Error reading message: %v", err)
			continue
		}

		// Preprocess data
		data, err := preprocessData(msg.Value)
		if err != nil {
			log.Printf("Error preprocessing data: %v", err)
			continue
		}

		// Perform inference
		anomalyScore, err := infer(model, data)
		if err != nil {
			log.Printf("Error during inference: %v", err)
			continue
		}

		// Evaluate anomaly
		if anomalyScore > float32(anomalyThreshold) {
			fmt.Printf("Anomaly detected at %v: Score: %.2f, Data: %v\n", time.Now(), anomalyScore, data)
		} else {
			fmt.Printf("Normal data at %v: Score: %.2f\n", time.Now(), anomalyScore)
		}
	}
}

func main() {
	// Parse command-line flags
	broker := flag.String("broker", "localhost:9092", "Kafka broker address")
	topic := flag.String("topic", "anomaly_data", "Kafka topic")
	modelPath := flag.String("model", "./model.tflite", "Path to TensorFlow Lite model")
	flag.Float64Var(&anomalyThreshold, "threshold", 0.5, "Anomaly score threshold")
	flag.Parse()

	// Load the TensorFlow Lite model
	model, err := loadModel(*modelPath)
	if err != nil {
		log.Fatalf("Failed to load model: %v", err)
	}

	// Set up Kafka consumer
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": *broker,
		"group.id":          "anomaly_detector",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		log.Fatalf("Failed to create Kafka consumer: %v", err)
	}

	// Subscribe to Kafka topic
	if err := consumer.Subscribe(*topic, nil); err != nil {
		log.Fatalf("Failed to subscribe to topic: %v", err)
	}

	// Start consuming messages and evaluating anomalies
	consumeMessages(consumer, model)
}
