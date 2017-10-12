package plasma_client

import (
	"github.com/openfresh/plasma-go/config"
	"github.com/openfresh/plasma-go/event"
	"github.com/pkg/errors"
)

type Publisher interface {
	Publish(payload event.Payload) error
}

// New returns a publisher by Config
func New(config config.Config) (Publisher, error) {
	var publisher Publisher
	var err error

	switch config.Type {
	case TypeRedis:
		publisher, err = newRedis(config)
		if err != nil {
			return nil, errors.Wrap(err, "failed to newRedis")
		}
	default:
		return nil, errors.Wrapf(err, "can't get such %s type publisher", config.Type)
	}

	return publisher, nil
}
