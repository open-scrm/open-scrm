package publish

import (
	"context"
	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/pkg/errors"
)

type Producer struct {
	topic string
	sarama.SyncProducer
}

func newProducer(address []string, topic string) (*Producer, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll // Wait for all in-sync replicas to ack the message
	config.Producer.Retry.Max = 3                    // Retry up to 10 times to produce the message
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(address, config)
	if err != nil {
		return nil, errors.Wrap(err, "创建producer失败")
	}
	return &Producer{SyncProducer: producer, topic: topic}, nil
}

func (p *Producer) sendMessage(ctx context.Context, v interface{}) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	_, _, err = p.SyncProducer.SendMessage(&sarama.ProducerMessage{
		Topic: p.topic,
		Value: sarama.ByteEncoder(data),
	})
	return err
}

func (p *Producer) PublishOne(ctx context.Context, data interface{}) error {
	return p.sendMessage(ctx, data)
}

func (p *Producer) Close() error {
	return p.SyncProducer.Close()
}
