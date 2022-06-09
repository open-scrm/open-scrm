package publish

import (
	"context"
	"github.com/open-scrm/open-scrm/configs"
)

type AddressBookDeptPublisher struct {
	*Producer
}

// NewAddressBookDeptPublisher 一个分区就行.
func NewAddressBookDeptPublisher(ctx context.Context, conf *configs.Config) (*AddressBookDeptPublisher, error) {
	p, err := newProducer(conf.Kafka.Address, conf.Kafka.Topics.DepartmentChangeEvent)
	if err != nil {
		return nil, err
	}
	return &AddressBookDeptPublisher{Producer: p}, nil
}
