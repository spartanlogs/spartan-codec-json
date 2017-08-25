package codecs

import (
	"encoding/json"

	"github.com/spartanlogs/spartan/codecs"
	"github.com/spartanlogs/spartan/event"
	"github.com/spartanlogs/spartan/utils"
)

// The JSONCodec encodes/decodes an event as JSON.
type JSONCodec struct {
	codecs.BaseCodec
}

func init() {
	codecs.Register("json", newJSONCodec)
}

func newJSONCodec(options utils.InterfaceMap) (codecs.Codec, error) {
	return &JSONCodec{}, nil
}

// Encode Event as JSON object.
func (c *JSONCodec) Encode(e *event.Event) []byte {
	data := e.Data()
	j, _ := json.Marshal(data)
	return j
}
