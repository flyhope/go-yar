package pack

import (
	"bytes"
	"github.com/vmihailenco/msgpack"
)

// msgpack处理器，兼容json tag定义
type EncoderMsgpack struct {
}

func (p *EncoderMsgpack) Encode(request *Request) ([]byte, error) {
	var buf bytes.Buffer
	encoder := msgpack.NewEncoder(&buf)
	encoder.UseJSONTag(true)
	err := encoder.Encode(request)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), err
}

func (p *EncoderMsgpack) Decode(body []byte, response *Response) error{
	reader := bytes.NewReader(body)
	decoder := msgpack.NewDecoder(reader)
	decoder.UseJSONTag(true)
	return decoder.Decode(response)
}

func (p *EncoderMsgpack) ContentType() string {
	return "application/x-msgpack"
}

