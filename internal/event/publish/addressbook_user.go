package publish

import (
	"context"
	"github.com/open-scrm/open-scrm/configs"
)

type AddressBookUserPublisher struct {
	p *Producer
}

// NewAddressBookUserPublisher 一个分区就行.
func NewAddressBookUserPublisher(ctx context.Context, conf *configs.Config) (*AddressBookUserPublisher, error) {
	p, err := newProducer(conf.Kafka.Address, conf.Kafka.Topics.UserChangeEvent)
	if err != nil {
		return nil, err
	}
	return &AddressBookUserPublisher{p: p}, nil
}

func (p *AddressBookUserPublisher) PublishOne(ctx context.Context, data interface{}) error {
	return p.p.sendMessage(ctx, data)
}

func (p *AddressBookUserPublisher) Close() error {
	return p.p.Close()
}
