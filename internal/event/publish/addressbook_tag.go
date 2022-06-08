package publish

import (
	"context"
	"encoding/json"
	"github.com/open-scrm/open-scrm/configs"
	"github.com/segmentio/kafka-go"
)

type AddressBookTagPublisher struct {
	w *kafka.Writer
}

// NewAddressBookTagPublisher 一个分区就行.
func NewAddressBookTagPublisher(ctx context.Context, conf *configs.Config) (*AddressBookTagPublisher, error) {
	w := &kafka.Writer{
		Addr:     kafka.TCP(conf.Kafka.Address),
		Topic:    conf.Kafka.Topics.TagChangeEvent,
		Balancer: &kafka.LeastBytes{},
	}
	return &AddressBookTagPublisher{w: w}, nil
}

func (p *AddressBookTagPublisher) PublishOne(ctx context.Context, data interface{}) error {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return p.w.WriteMessages(ctx, kafka.Message{
		Value: dataBytes,
	})
}

func (p *AddressBookTagPublisher) Close() error {
	return p.w.Close()
}
