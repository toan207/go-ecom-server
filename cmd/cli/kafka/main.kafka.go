package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
)

var (
	kafkaProducer *kafka.Writer
)

const (
	kafkaURL   = "localhost:9092"
	kafkaTopic = "user_topic_test"
)

func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:        brokers,
		GroupID:        groupID,
		Topic:          topic,
		MinBytes:       10e3, // 10KB
		MaxBytes:       10e6, // 10MB
		CommitInterval: time.Second,
		StartOffset:    kafka.FirstOffset,
	})
}

func getKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

type StockInfo struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

func NewStockInfo(message, type_ string) *StockInfo {
	return &StockInfo{
		Message: message,
		Type:    type_,
	}
}

func actionStock(c *gin.Context) {
	s := NewStockInfo(c.Query("message"), c.Query("type"))
	body := map[string]interface{}{
		"action": "action",
		"info":   s,
	}
	jsonBody, _ := json.Marshal(body)

	msg := kafka.Message{
		Key:   []byte("action"),
		Value: []byte(string(jsonBody)),
	}

	err := kafkaProducer.WriteMessages(c, msg)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to write message"})
		return
	}
	c.JSON(200, gin.H{"message": "Message sent successfully"})
}

func RegisterConsumerATC(id int) {
	kafkaGroupId := fmt.Sprintf("group_%d", id)
	reader := getKafkaReader(kafkaURL, kafkaTopic, kafkaGroupId)
	defer reader.Close()

	fmt.Printf("Consumer %d started\n", id)
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Printf("Error reading message: %v\n", err)
			continue
		}
		fmt.Printf("Consumer %d: message=%s, topic=%s, partition=%d, offset=%d\n",
			id, string(m.Value), m.Topic, m.Partition, m.Offset)
	}
}

func main() {
	r := gin.Default()
	kafkaProducer = getKafkaWriter(kafkaURL, kafkaTopic)
	defer kafkaProducer.Close()
	r.POST("/action/stock", actionStock)
	go RegisterConsumerATC(1)
	go RegisterConsumerATC(2)
	r.Run(":8089")
}
