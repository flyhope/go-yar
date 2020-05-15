package yar

import (
	"io/ioutil"
	"net/http"
)

type client struct {
	Request *Request
}

func Call(addr string, method string, params interface{}) ([]byte, error) {
	c := &client{
		Request: NewRequest(addr, method, params),
	}
	return c.send()
}

func (c *client) send() ([]byte, error) {
	packHandler := getPackHandler(c.Request.Protocol)
	data, err := packHandler.Encode(c.Request)
	if err != nil {
		return nil, err
	}

	// 拼接body
	var p [8]byte
	for i, s := range []byte(c.Request.Protocol) {
		p[i] = s
	}

	header := NewHeader()
	header.Packager = p
	buffer := header.Bytes()
	buffer.Write(data)

	resp, err := http.Post(c.Request.Addr, packHandler.ContentType(), buffer)
	if err != nil {
		return nil, err
	}

	body , err := ioutil.ReadAll(resp.Body)
	return body, err
}