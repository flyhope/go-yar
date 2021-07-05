package pack

// Protocol request protocol
type Protocol [8]byte

var (
	ProtocolJson    = Protocol{'J', 'S', 'O', 'N', '\000', 'Y', 'A', 'R'}
	ProtocolMsgpack = Protocol{'M', 'S', 'G', 'P', 'A', 'C', 'K'}
)

// Pack define encode/decode interface
type Pack interface {
	Encode(*Request) ([]byte, error)
	ContentType() string
	Decode([]byte, *Response) error
}

// GetPackHandler 根据协议获取编码、解码器
func GetPackHandler(protocol Protocol) Pack {
	var result Pack

	switch protocol {
	case ProtocolJson:
		result = new(EncoderJson)
	case ProtocolMsgpack:
		result = new(EncoderMsgpack)
	}

	return result
}
