package plasma_client

import (
	"encoding/json"

	"github.com/openfresh/plasma-go/config"
	"github.com/openfresh/plasma-go/event"
	"github.com/pkg/errors"
	"gopkg.in/redis.v5"
)

const TypeRedis = "redis"

type Redis struct {
	client  *redis.Client
	config  config.Redis
	channel string
}

func newRedis(config config.Config) (Publisher, error) {
	redisConf := config.Redis
	addr := redisConf.Addr
	opt := &redis.Options{
		Addr:     addr,
		Password: redisConf.Password,
		DB:       redisConf.DB,
	}

	client := redis.NewClient(opt)
	return &Redis{
		client:  client,
		config:  redisConf,
		channel: config.Redis.Channel,
	}, nil
}

// NOTE: If Go version less than 1.8, RawMessage marshals as base64
// https://groups.google.com/forum/#!topic/Golang-Nuts/38ShOlhxAYY
type internalPayload struct {
	Meta event.MetaData   `json:"meta"`
	Data *json.RawMessage `json:"data"`
}

// Publish sends payload to the redis channel
func (r *Redis) Publish(payload event.Payload) error {
	eventType := payload.Meta.Type

	p := internalPayload{
		Meta: payload.Meta,
		Data: &payload.Data,
	}

	message, err := json.Marshal(p)
	if err != nil {
		return errors.Wrap(err, "failed to marshal json")
	}
	if err := r.client.Publish(r.channel, string(message)).Err(); err != nil {
		return errors.Wrapf(err, "failed to publish %s:%s", eventType, message)
	}
	return nil
}
