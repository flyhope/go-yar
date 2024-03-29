package pack

import (
	"encoding/json"
)

// JSON处理器
type EncoderJson struct {
}

func (p *EncoderJson) Encode(request *Request) ([]byte, error) {
	return json.Marshal(request)
}

func (p *EncoderJson) Decode(body []byte, response *Response) error {
	response.Protocol = ProtocolJson
	return json.Unmarshal(body, response)
}

func (p *EncoderJson) ContentType() string {
	return "application/json"
}

func (p *EncoderJson) ShowProtocol() Protocol {
	return ProtocolJson
}
