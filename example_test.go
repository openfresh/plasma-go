package plasma_client_test

import (
	"encoding/json"
	"log"

	"github.com/openfresh/plasma-go"
	"github.com/openfresh/plasma-go/config"
	"github.com/openfresh/plasma-go/event"
)

func ExampleNew() {
	config := config.Config{
		Type: plasma_client.TypeRedis,
	}

	client, err := plasma_client.New(config)
	if err != nil {
		log.Fatal(err)
	}

	payload := event.Payload{
		Meta: event.MetaData{
			Type: "video:1234:views",
		},
		Data: json.RawMessage(`{"data":55301}`),
	}
	err = client.Publish(payload)
	if err != nil {
		log.Fatal(err)
	}

}
