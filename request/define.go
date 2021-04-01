package request

import (
	"context"
	"github.com/flyhope/go-yar/logger"
	"net/http"
)

type Handler interface {
	SetLog(logger.LogTrace)
	Do(ctx context.Context, req *http.Request) ([]byte, error)
}

type Abstract struct {
	logger.LogTrace
}

func (a *Abstract) SetLog(l logger.LogTrace) {
	a.LogTrace = l
}
