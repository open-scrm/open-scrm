package publish

import (
	"context"
	"encoding/json"
	"github.com/open-scrm/open-scrm/configs"
	"github.com/segmentio/kafka-go"
)

type AddressBookUserPublisher struct {
	w *kafka.Writer
}

// NewAddressBookDeptPublisher 一个分区就行.
func NewAddressBookUserPublisher(ctx context.Context, conf *configs.Config) (*AddressBookUserPublisher, error) {
	w := &kafka.Writer{
		Addr:     kafka.TCP(conf.Kafka.Address),
		Topic:    conf.Kafka.Topics.UserChangeEvent,
		Balancer: &kafka.LeastBytes{},
	}
	return &AddressBookUserPublisher{w: w}, nil
}

func (p *AddressBookUserPublisher) PublishOne(ctx context.Context, data interface{}) error {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return p.w.WriteMessages(ctx, kafka.Message{
		Value: dataBytes,
	})
}

func (p *AddressBookUserPublisher) Close() error {
	return p.w.Close()
}
