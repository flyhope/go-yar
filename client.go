package yar

import (
	"context"
	"errors"
	"fmt"
	"github.com/flyhope/go-yar/logger"
	"github.com/flyhope/go-yar/pack"
	"github.com/flyhope/go-yar/request"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"
)

type Client struct {
	ctx           context.Context
	Request       *pack.Request
	Response      *pack.Response
	Http          *http.Request
	RequestClient request.Handler
	logTrace      logger.LogTrace
}

// 初始化一个客户端
func NewClient(ctx context.Context, addr string, method string, params interface{}) (*Client, error) {
	handler := &request.Http{Client: &http.Client{Timeout: time.Second}}
	return NewWithHandler(ctx, handler, addr, method, params)
}

// 初始化一个客户端
func NewWithHandler(ctx context.Context, requestClient request.Handler, addr string, method string, params interface{}) (*Client, error) {
	httpRequest, err := http.NewRequest(http.MethodPost, addr, nil)
	if err != nil {
		return nil, err
	}

	httpRequest.Header.Set("User-Agent", "Go Yar Rpc-0.1")
	c := &Client{
		ctx:           ctx,
		Request:       pack.NewRequest(addr, method, params),
		Response:      new(pack.Response),
		Http:          httpRequest,
		RequestClient: requestClient,
	}

	return c, nil
}

// 设置日志追踪方法
func (c *Client) SetLogTrace(logTrace logger.LogTrace) {
	c.logTrace = logTrace
	c.RequestClient.SetLog(logTrace)
}

// 设置返回值结构体
func (c *Client) SetResponseRetStruct(retVal interface{}) *Client {
	c.Response.Retval = retVal
	return c
}

// 开始发送请求数据
func (c *Client) Send() error {
	packHandler := pack.GetPackHandler(c.Request.Protocol)
	data, err := packHandler.Encode(c.Request)
	if err != nil {
		return err
	}

	// 拼接body
	header := pack.NewHeader(c.Request.Protocol)
	buffer := header.Bytes()
	buffer.Write(data)


	c.Http.Body = ioutil.NopCloser(buffer)
	c.Http.Header.Set("Content-Type", packHandler.ContentType())
	c.Http.Header.Add("Content-Length", fmt.Sprintf("%d", buffer.Len()))

	logger.Log.WithFields(logrus.Fields{"YAR": "Request"}).Debug(string(data))

	// 发送请求
	body, err := c.RequestClient.Do(c.ctx, c.Http)
	if err != nil {
		return err
	}

	// 解析处理
	headerData := pack.NewHeaderWithBody(body, c.Request.Protocol)
	packHandler = pack.GetPackHandler(headerData.Packager)
	if packHandler == nil {
		return errors.New("can't unpack yar response")
	}
	bodyContent := body[pack.ProtocolLength+pack.PackagerLength:]
	err = packHandler.Decode(bodyContent, c.Response)

	if c.Response.Except != nil {
		logger.Log.WithFields(logrus.Fields{"YAR": "Except"}).Debug(c.Response.Except)
	}

	logger.Log.WithFields(logrus.Fields{"YAR": "BodyContent"}).Debug(string(bodyContent))

	if err != nil {
		return err
	}

	if c.Response.Except != nil {
		return c.Response.Except
	}

	return nil
}
