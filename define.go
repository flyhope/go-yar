package yar

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"math/rand"
)

// 头信息
const (
	MagicNumber = 0x80DFEC60
)

type Header struct {
	Id          uint32
	Version     uint16
	MagicNumber uint32
	Reserved    uint32
	Provider    [28]byte
	Encrypt     uint32
	Token       [32]byte
	BodyLength  uint32
	Packager    [8]byte
}

func NewHeader() *Header {
	proto := new(Header)
	proto.MagicNumber = MagicNumber
	return proto

}

func NewHeaderWithBytes(payload *bytes.Buffer) *Header {
	p := NewHeader()
	p.Init(payload)
	return p
}

func (h *Header) Init(payload *bytes.Buffer) bool {

	binary.Read(payload, binary.BigEndian, &h.Id)
	binary.Read(payload, binary.BigEndian, &h.Version)
	binary.Read(payload, binary.BigEndian, &h.MagicNumber)
	binary.Read(payload, binary.BigEndian, &h.Reserved)
	binary.Read(payload, binary.BigEndian, &h.Provider)
	binary.Read(payload, binary.BigEndian, &h.Encrypt)
	binary.Read(payload, binary.BigEndian, &h.Token)
	binary.Read(payload, binary.BigEndian, &h.BodyLength)
	binary.Read(payload, binary.BigEndian, &h.Packager)
	return true
}

func (h *Header) Bytes() *bytes.Buffer {

	buffer := new(bytes.Buffer)
	err := binary.Write(buffer, binary.BigEndian, h)

	if err != nil {
		return nil
	}
	return buffer
}

// 请求协议
type Protocol string

const (
	ProtocolJson    Protocol = "json\000\000\000\000"
	ProtocolMsgpack Protocol = "msgpack\000"
)

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
		Protocol: ProtocolMsgpack,
	}
	return request
}

// 响应
type StatusType int

const (
	StatusOkey       StatusType = 0x0
	ErrPackager      StatusType = 0x1
	ErrProtocol      StatusType = 0x2
	ErrRequest       StatusType = 0x4
	ErrOutput        StatusType = 0x8
	ErrTransport     StatusType = 0x10
	ErrForbidden     StatusType = 0x20
	ErrException     StatusType = 0x40
	ErrEmptyResponse StatusType = 0x80
)

type exceptionDefine struct {
	Type    string `json:"_type"`
	Code    int    `json:"code"`
	File    string `json:"file"`
	Line    uint   `json:"line"`
	Message string `json:"message"`
}

type Exception struct {
	exceptionDefine
}

func (e *Exception) Error() string {
	return e.Message
}

func (e *Exception) UnmarshalJSON(b []byte) (err error) {
	err = json.Unmarshal(b, &e.exceptionDefine)
	if errType, ok := err.(*json.UnmarshalTypeError); ok {
		if errType.Value == "string" {
			err = json.Unmarshal(b, &e.Message)
		}
	}
	return err
}

type Response struct {
	Protocol Protocol    `json:"-" msgpack:"-"`
	Id       uint32      `json:"i" msgpack:"i"`
	Except   *Exception  `json:"e" msgpack:"e"`
	Out      string      `json:"o" msgpack:"o"`
	Status   StatusType  `json:"s" msgpack:"s"`
	Retval   interface{} `json:"r" msgpack:"r"`
}
