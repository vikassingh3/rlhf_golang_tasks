package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/IBM/sarama"                              // Correct Kafka consumer package
	"github.com/tensorflow/tensorflow/go"                // TensorFlow Go package
	"github.com/tensorflow/tensorflow/go/core/framework" // TensorFlow core framework
	"github.com/tensorflow/tensorflow/go/tensor"         // TensorFlow tensor support
)

// Data struct representing the streaming data
type Data struct {
	Feature1 float32 `json:"feature1"`
	Feature2 float32 `json:"feature2"`
}

func main() {
	// Kafka consumer configuration
	config := sarama.NewConfig()
	config.Consumer.Group.ID = "anomaly-detection-group"
	config.Metadata.RefreshFrequency = 60 * 5 * time.Second
	config.Offset.Initial = sarama.OffsetNewest

	// Connect to Kafka
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatal(err)
	}
	defer consumer.Close()

	// Load TensorFlow SavedModel
	sess, err := tf.NewSession(nil)
	if err != nil {
		log.Fatal(err)
	}
	defer sess.Close()

	err = sess.LoadSavedModel("./path/to/your/saved_model", []string{"serve"})
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll("./path/to/your/saved_model")

	// Run the Kafka consumer
	go func() {
		for {
			select {
			case msg := <-consumer.Messages():
				var data Data
				err = json.Unmarshal(msg.Value, &data)
				if err != nil {
					log.Printf("Error unmarshaling message: %v", err)
					continue
				}

				// Prepare data for TensorFlow model
				input := []*framework.TensorProto{{
					Dtype: framework.DT_FLOAT,
					TensorShape: &framework.TensorShapeProto{
						Dim: []*framework.TensorShapeProto_Dim{{Size: 1}},
					},
					FloatVal: []float32{data.Feature1, data.Feature2},
				}}

				// Run inference
				outputs, err := sess.Run(context.Background(), nil, []string{"anomaly_score"}, input)
				if err != nil {
					log.Printf("Error running inference: %v", err)
					continue
				}

				anomalyScore, err := tensor.Float64(outputs[0])
				if err != nil {
					log.Printf("Error converting output to float: %v", err)
					continue
				}

				// Check for anomaly
				if anomalyScore > 0.5 { // Adjust threshold as needed
					log.Printf("Anomaly detected! Score: %.2f, Data: %v", anomalyScore, data)
				} else {
					log.Printf("Normal data. Score: %.2f, Data: %v", anomalyScore, data)
				}
			}
		}
	}()

	// Keep the program running
	select {}
}
