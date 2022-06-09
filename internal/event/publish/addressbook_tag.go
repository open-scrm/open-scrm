package publish

import (
	"context"
	"github.com/open-scrm/open-scrm/configs"
)

type AddressBookTagPublisher struct {
	*Producer
}

// NewAddressBookTagPublisher 一个分区就行.
func NewAddressBookTagPublisher(ctx context.Context, conf *configs.Config) (*AddressBookTagPublisher, error) {
	p, err := newProducer(conf.Kafka.Address, conf.Kafka.Topics.TagChangeEvent)
	if err != nil {
		return nil, err
	}
	return &AddressBookTagPublisher{Producer: p}, nil
}
