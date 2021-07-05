package yar

import (
	"errors"
	"github.com/flyhope/go-yar/pack"
	"net/http"
)

type serverFunc func(params interface{}) (interface{}, error)

type httpRegistry struct {
	executer map[string]serverFunc
}

// Method add a server method executer
func (h *httpRegistry) Method(name string, method serverFunc) {
	h.executer[name] = method
}

// Http handle a http request and response json data
func (h *httpRegistry) Http(request *http.Request) []byte {
	reader, err := request.GetBody()
	if err != nil {
		return serverError(pack.ErrPackager, err)
	}

	packager := make([]byte, 0, pack.PackagerLength)
	length, err :=  reader.Read(packager)
	if err != nil {
		return serverError(pack.ErrPackager, err)
	}

	if length != pack.PackagerLength {
		return serverError(pack.ErrPackager, errors.New("packager length bad"))
	}

	return nil
}

// response server error
func serverError(typ pack.StatusType ,err error) []byte {
	// @todo
	return nil
}

// NewHttpRegistry create a http registry
func NewHttpRegistry() *httpRegistry {
	return &httpRegistry{
		executer: map[string]serverFunc{},
	}
}
