package yar

import (
	"github.com/flyhope/go-yar/pack"
	"io/ioutil"
	"net/http"
	"time"
)

type client struct {
	Request *pack.Request
	Response *pack.Response
	Http *http.Client
}

// 初始化一个客户端
func Client(addr string, method string, params interface{}) *client {
	c := &client{
		Request: pack.NewRequest(addr, method, params),
		Response: new(pack.Response),
		Http:  &http.Client{Timeout: time.Second},
	}

	return c
}

// 设置返回值结构体
func (c *client) SetResponseRetStruct(retVal interface{}) *client {
	c.Response.Retval = retVal
	return c
}

// 开始发送请求数据
func (c *client) Send() error {
	packHandler := pack.GetPackHandler(c.Request.Protocol)
	data, err := packHandler.Encode(c.Request)
	if err != nil {
		return err
	}

	// 拼接body
	header := pack.NewHeader(c.Request.Protocol)
	buffer := header.Bytes()
	buffer.Write(data)

	// 发送请求
	resp, err := c.Http.Post(c.Request.Addr, packHandler.ContentType(), buffer)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)

	// 解析处理
	headerData := pack.NewHeaderWithBody(body, c.Request.Protocol)
	packHandler = pack.GetPackHandler(headerData.Packager)
	bodyContent := body[pack.ProtocolLength + pack.PackagerLength:]
	return packHandler.Decode(bodyContent, c.Response)
}
