package services

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/gin-gonic/gin/json"
	log "github.com/sirupsen/logrus"
)

const KafkaKey = "kafka"

func GetKafka(c context.Context) KafkaInterface {
	return c.Value(KafkaKey).(KafkaInterface)
}

type KafkaInterface interface {
	SendValue(message interface{}, topic string) error
}

/* MOCK */
type KafkaMock struct{}

func (f *KafkaMock) SendValue(message interface{}, topic string) error {
	return nil
}

type Kafka struct {
	Producer sarama.SyncProducer
}

func NewKafka(producer sarama.SyncProducer) KafkaInterface {
	return &Kafka{Producer: producer}
}

func (f *Kafka) SendValue(message interface{}, topic string) error {
	// We are not setting a message key, which means that all messages will
	// be distributed randomly over the different partitions.
	messageJson, err := json.Marshal(message)
	if err != nil {
		return err
	}

	partition, offset, err := f.Producer.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(messageJson),
	})
	if err != nil {
		return err
	}

	log.Info("Data sent to topic: ", topic, ", partition: ", partition, ", offset: ", offset)

	return nil
}
