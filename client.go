package yar

import (
	"errors"
	"github.com/flyhope/go-yar/pack"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"
)

type Client struct {
	Request    *pack.Request
	Response   *pack.Response
	Http       *http.Request
	HttpClient *http.Client
	LogTrace   LogTrace
}

// 初始化一个客户端
func NewClient(addr string, method string, params interface{}) (*Client, error) {
	httpRequest, err := http.NewRequest(http.MethodPost, addr, nil)
	if err != nil {
		return nil, err
	}

	httpRequest.Header.Set("User-Agent", "Go Yar Rpc-0.1")
	c := &Client{
		Request:    pack.NewRequest(addr, method, params),
		Response:   new(pack.Response),
		Http:       httpRequest,
		HttpClient: &http.Client{Timeout: time.Second},
	}

	return c, nil
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

	Log.WithFields(logrus.Fields{"YAR": "Request"}).Debug(string(data))

	// 发送请求
	timeStart := time.Now()
	resp, err := c.HttpClient.Do(c.Http)

	// 通过接口记录跟踪日志
	if c.LogTrace != nil {
		timeEnd := time.Now()
		traceData := &LogTraceData{
			TimeStart: timeStart,
			TimeEnd:   timeEnd,
			Request:   c.Http,
			Err:       err,
		}
		c.LogTrace.Trace(traceData)
	}

	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)

	// 解析处理
	headerData := pack.NewHeaderWithBody(body, c.Request.Protocol)
	packHandler = pack.GetPackHandler(headerData.Packager)
	if packHandler == nil {
		return errors.New("can't unpack yar response")
	}
	bodyContent := body[pack.ProtocolLength+pack.PackagerLength:]
	err = packHandler.Decode(bodyContent, c.Response)

	if c.Response.Except != nil {
		Log.WithFields(logrus.Fields{"YAR": "Except"}).Debug(c.Response.Except)
	}

	Log.WithFields(logrus.Fields{"YAR": "BodyContent"}).Debug(string(bodyContent))

	if err != nil {
		return err
	}

	if c.Response.Except != nil {
		return c.Response.Except
	}

	return nil
}
