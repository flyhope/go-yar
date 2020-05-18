package pack

import "math/rand"

// 请求结构体
type Request struct {
	Protocol Protocol    `json:"-" msgpack:"-"`
	Addr     string      `json:"-" msgpack:"-"`
	Id       uint32      `json:"i" msgpack:"i"`
	Method   string      `json:"m" msgpack:"m"`
	Params   interface{} `json:"p" msgpack:"p"`
}

func NewRequest(addr string, method string, params interface{}) (request *Request) {
	request = &Request{
		Id:       rand.Uint32(),
		Addr:     addr,
		Method:   method,
		Params:   params,
		Protocol: ProtocolJson,
	}
	return request
}

