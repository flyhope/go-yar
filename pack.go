package yar

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack"
)

type Pack interface {
	Encode(*Request) ([]byte, error)
	ContentType() string
	//Decode([]byte) *Response
}

// JSON处理器
type PackJson struct {
}

func (p *PackJson) Encode(request *Request) ([]byte, error) {
	return json.Marshal(request)
}

func (p *PackJson) ContentType() string {
	return "application/application/json"
}

// msgpack处理器
type PackMsgpack struct {
}

func (p *PackMsgpack) Encode(request *Request) ([]byte, error) {
	return msgpack.Marshal(request)
}

func (p *PackMsgpack) ContentType() string {
	return "application/x-msgpack"
}

// 根据协议获取编码、解码器
func getPackHandler(protocol Protocol) Pack {
	var result Pack

	switch protocol {
	case ProtocolJson:
		result = new(PackJson)
	case ProtocolMsgpack:
		result = new(PackMsgpack)
	}

	return result
}
