package services

import (
	"github.com/Shopify/sarama/mocks"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSendValueSucceed(t *testing.T) {
	producerMock := mocks.NewSyncProducer(t, nil)
	producerMock.ExpectSendMessageAndSucceed()

	kafkaImpl := NewKafka(producerMock)

	type foo struct {
		Bar string `json:"bar"`
	}

	err := kafkaImpl.SendValue(foo{}, "bar")
	assert.Nil(t, err)
}

func TestSendValueFail(t *testing.T) {
	producerMock := mocks.NewSyncProducer(t, nil)
	producerMock.ExpectSendMessageAndFail(errors.New("Failure"))

	kafkaImpl := NewKafka(producerMock)

	type foo struct {
		Bar string `json:"bar"`
	}

	err := kafkaImpl.SendValue(foo{}, "bar")
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "Failure")
}

func TestSendValueJsonMarshallFail(t *testing.T) {
	producerMock := mocks.NewSyncProducer(t, nil)
	kafkaImpl := NewKafka(producerMock)

	err := kafkaImpl.SendValue(make(chan int), "bar")
	assert.NotNil(t, err)
}
