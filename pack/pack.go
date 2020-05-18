package pack

// 请求协议
type Protocol [8]byte

var (
	ProtocolJson    = Protocol{'J', 'S', 'O', 'N', '\000', 'Y', 'A', 'R'}
	ProtocolMsgpack = Protocol{'M', 'S', 'G', 'P', 'A', 'C', 'K'}
)

// 定义数据编码解码接口
type Pack interface {
	Encode(*Request) ([]byte, error)
	ContentType() string
	Decode([]byte, *Response) error
}

// 根据协议获取编码、解码器
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
