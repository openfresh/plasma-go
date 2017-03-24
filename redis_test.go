package plasma_client

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
	"time"

	redis "gopkg.in/redis.v5"

	"github.com/openfresh/plasma-go/config"
	"github.com/openfresh/plasma-go/event"
)

func receive(config config.Config) (string, error) {
	redisConf := config.Redis
	opt := &redis.Options{
		Addr:     redisConf.Addr,
		Password: redisConf.Password,
		DB:       redisConf.DB,
	}
	client := redis.NewClient(opt)

	ps, err := client.Subscribe(redisConf.Channel)
	if err != nil {
		return "", err
	}

	msg, err := ps.ReceiveMessage()
	if err != nil {
		return "", err
	}

	return msg.Payload, nil
}

func TestRedisPublish(t *testing.T) {
	conf := config.Config{
		Redis: config.Redis{
			Addr:    "localhost:6379",
			DB:      0,
			Channel: "plasma",
		},
	}
	pub, err := newRedis(conf)
	if err != nil {
		t.Fatal(err)
	}

	payload := event.Payload{
		Meta: event.MetaData{
			Type: "test",
		},
		Data: json.RawMessage(`{"data":"test message"}`),
	}

	msgChan := make(chan string)
	go func() {
		msg, err := receive(conf)
		if err != nil {
			t.Fatal(err)
		}
		msgChan <- msg
	}()

	time.Sleep(10 * time.Millisecond)
	if err := pub.Publish(payload); err != nil {
		log.Fatal(err)
	}

	msg := <-msgChan

	p := event.Payload{}
	if err := json.Unmarshal([]byte(msg), &p); err != nil {
		log.Fatal(err)
	}

	if p.Meta.Type != payload.Meta.Type {
		t.Error(fmt.Sprintf("Expected: %s, Actual: %s", payload.Meta.Type, p.Meta.Type))
	}

	if string(p.Data) != string(payload.Data) {
		t.Error(fmt.Sprintf("Expected: %s, Actual: %s", payload.Data, p.Data))
	}
}
