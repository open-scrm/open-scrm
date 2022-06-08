package publish

import (
	"context"
	"encoding/json"
	"github.com/open-scrm/open-scrm/configs"
	"github.com/segmentio/kafka-go"
)

type AddressBookBatchJobResultPublisher struct {
	w *kafka.Writer
}

// NewAddressBookBatchJobResultPublisher 一个分区就行.
func NewAddressBookBatchJobResultPublisher(ctx context.Context, conf *configs.Config) (*AddressBookBatchJobResultPublisher, error) {
	w := &kafka.Writer{
		Addr:     kafka.TCP(conf.Kafka.Address),
		Topic:    conf.Kafka.Topics.BatchJobResult,
		Balancer: &kafka.LeastBytes{},
	}
	return &AddressBookBatchJobResultPublisher{w: w}, nil
}

func (p *AddressBookBatchJobResultPublisher) PublishOne(ctx context.Context, data interface{}) error {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return p.w.WriteMessages(ctx, kafka.Message{
		Value: dataBytes,
	})
}

func (p *AddressBookBatchJobResultPublisher) Close() error {
	return p.w.Close()
}
