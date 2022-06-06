package publish

import (
	"context"
	"encoding/json"
	"github.com/open-scrm/open-scrm/configs"
	"github.com/segmentio/kafka-go"
)

type AddressBookPublisher struct {
	w *kafka.Writer
}

// NewAddressBookDeptPublisher 一个分区就行.
func NewAddressBookDeptPublisher(ctx context.Context, conf *configs.Config) (*AddressBookPublisher, error) {
	w := &kafka.Writer{
		Addr:     kafka.TCP(conf.Kafka.Address),
		Topic:    conf.Kafka.Topics.DepartmentChangeEvent,
		Balancer: &kafka.LeastBytes{},
	}
	return &AddressBookPublisher{w: w}, nil
}

func (p *AddressBookPublisher) PublishOne(ctx context.Context, data interface{}) error {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return p.w.WriteMessages(ctx, kafka.Message{
		Value: dataBytes,
	})
}

func (p *AddressBookPublisher) Close() error {
	return p.w.Close()
}
