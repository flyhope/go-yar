package pack

import (
	"encoding/json"
)

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

// 基础异常结构体
type exceptionDefine struct {
	Type    string `json:"_type"`
	Code    int32  `json:"code"`
	File    string `json:"file"`
	Line    uint   `json:"line"`
	Message string `json:"message"`
}

// 供JSON解析及外部使用的异常结构体
type Exception struct {
	exceptionDefine
}

func (e *Exception) Error() string {
	return e.Message
}

func (e *Exception) GetCode() int32 {
	return e.Code
}

func (e *Exception) GetMeta() (result map[string]interface{}) {
	result = make(map[string]interface{})
	result["Type"] = e.Type
	result["File"] = e.File
	result["Line"] = e.Line
	return result
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

// 响应结构体
type Response struct {
	Protocol Protocol    `json:"-" msgpack:"-"`
	Id       uint32      `json:"i" msgpack:"i"`
	Except   *Exception  `json:"e" msgpack:"e"`
	Out      string      `json:"o" msgpack:"o"`
	Status   StatusType  `json:"s" msgpack:"s"`
	Retval   interface{} `json:"r" msgpack:"r"`
}
