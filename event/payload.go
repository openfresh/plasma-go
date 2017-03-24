package event

import "encoding/json"

type MetaData struct {
	Type string `json:"type"`
}

type Payload struct {
	Meta MetaData        `json:"meta"`
	Data json.RawMessage `json:"data"`
}
