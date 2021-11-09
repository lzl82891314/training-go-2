package toy_web

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type Context struct {
	Req       *http.Request
	RspWriter http.ResponseWriter
}

type ResponseDto struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	Timestamp int64       `json:"timestamp"`
}

func CreateContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Req:       req,
		RspWriter: w,
	}
}

func (ctx *Context) Request(data interface{}) error {
	body, err := io.ReadAll(ctx.Req.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, data)
}

func (ctx *Context) Response(v interface{}, err error) {
	resp := &ResponseDto{
		Timestamp: time.Now().UnixMilli(),
	}
	if err == nil {
		resp.Code = http.StatusOK
		resp.Message = "ok"
		resp.Data = v
	} else {
		resp.Code = http.StatusBadRequest
		resp.Message = err.Error()
		resp.Data = nil
	}
	ctx.doResponse(resp)
}

func (ctx *Context) NotFoundResponse(message string) {
	resp := &ResponseDto{
		Code:    http.StatusNotFound,
		Message: message,
	}
	ctx.doResponse(resp)
}

func (ctx *Context) doResponse(resp *ResponseDto) {
	data, err := json.Marshal(resp)
	if err != nil {
		http.Error(ctx.RspWriter, err.Error(), http.StatusInternalServerError)
		return
	}
	ctx.RspWriter.Header().Set("Content-Type", "application/json")
	if _, err := ctx.RspWriter.Write(data); err != nil {
		http.Error(ctx.RspWriter, err.Error(), http.StatusInternalServerError)
		return
	}
	ctx.RspWriter.WriteHeader(resp.Code)
}
