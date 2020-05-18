package pack

import "github.com/vmihailenco/msgpack"

// msgpack处理器
type EncoderMsgpack struct {
}

func (p *EncoderMsgpack) Encode(request *Request) ([]byte, error) {
	return msgpack.Marshal(request)
}

func (p *EncoderMsgpack) Decode(body []byte, response *Response) error{
	response.Protocol = ProtocolMsgpack
	return msgpack.Unmarshal(body, response)
}

func (p *EncoderMsgpack) ContentType() string {
	return "application/x-msgpack"
}

