package shell

import "encoding/json"

type Serializable interface {
	Marshal() ([]byte, error)
	Unmarshal(data []byte) error
}

type SerializationWrap struct {
	Name string
	Data json.RawMessage
}
