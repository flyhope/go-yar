package request

import (
	"context"
	"github.com/flyhope/go-yar/logger"
	"io/ioutil"
	"net/http"
	"time"
)

type Http struct {
	Abstract
	Client *http.Client
}

func (h *Http) Do(ctx context.Context, req *http.Request) ([]byte, error) {
	// 发送请求
	timeStart := time.Now()
	resp, err := h.Client.Do(req.WithContext(ctx))

	// 通过接口记录跟踪日志
	if h.LogTrace != nil {
		timeEnd := time.Now()
		traceData := &logger.LogTraceData{
			TimeStart: timeStart,
			TimeEnd:   timeEnd,
			Request:   req,
			Err:       err,
		}
		h.LogTrace.Trace(traceData)
	}

	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(resp.Body)
}
