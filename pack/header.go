package pack

import (
	"bytes"
	"encoding/binary"
)

// 头信息
const (
	MagicNumber = 0x80DFEC60
)

const (
	ProtocolLength = 82
	PackagerLength = 8
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
	Packager    Protocol
}

func NewHeader(packager Protocol) *Header {
	proto := new(Header)
	proto.MagicNumber = MagicNumber
	proto.Packager = packager
	return proto

}

func NewHeaderWithBody(body []byte, packager Protocol) *Header {
	payload := bytes.NewBuffer(body[0 : ProtocolLength+ProtocolLength])
	p := NewHeader(packager)
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
