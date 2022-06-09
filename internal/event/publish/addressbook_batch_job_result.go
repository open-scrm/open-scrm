package publish

import (
	"context"
	"github.com/open-scrm/open-scrm/configs"
)

type AddressBookBatchJobResultPublisher struct {
	*Producer
}

// NewAddressBookBatchJobResultPublisher 一个分区就行.
func NewAddressBookBatchJobResultPublisher(ctx context.Context, conf *configs.Config) (*AddressBookBatchJobResultPublisher, error) {
	p, err := newProducer(conf.Kafka.Address, conf.Kafka.Topics.BatchJobResult)
	if err != nil {
		return nil, err
	}
	return &AddressBookBatchJobResultPublisher{Producer: p}, nil
}
