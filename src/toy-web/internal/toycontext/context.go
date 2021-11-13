package toycontext

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
	tw "toy-web"
)

type Context struct {
	req *http.Request
	w   http.ResponseWriter
	ctx context.Context
}

var _ tw.IContext = &Context{}

func New(w http.ResponseWriter, req *http.Request) tw.IContext {
	return &Context{
		req: req,
		w:   w,
		ctx: req.Context(),
	}
}

// #region Response

func (ctx *Context) Json(s int, v interface{}, m string) error {
	res := &responseDto{
		Code:      s,
		Message:   m,
		Data:      v,
		Timestamp: time.Now().UnixMilli(),
	}

	data, err := json.Marshal(res)
	if err != nil {
		ctx.w.WriteHeader(http.StatusInternalServerError)
		return err
	}

	ctx.w.Header().Set("Content-Type", "application/json")
	ctx.w.WriteHeader(s)
	_, err = ctx.w.Write(data)
	return err
}

func (ctx *Context) Ok(v interface{}) error {
	return ctx.Json(http.StatusOK, v, "ok")
}

func (ctx *Context) NotFound(m string) error {
	return ctx.Json(http.StatusNotFound, nil, m)
}

// #endregion

// #region Param

func (ctx *Context) QueryAll() map[string][]string {
	if ctx.req != nil {
		return ctx.req.URL.Query()
	}
	return map[string][]string{}
}

func (ctx *Context) QueryInt(key string, def int) int {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		l := len(vals)
		if l > 0 {
			if val, err := strconv.Atoi(vals[l-1]); err == nil {
				return val
			}
		}
	}
	return def
}

func (ctx *Context) QueryStr(key string, def string) string {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		l := len(vals)
		if l > 0 {
			return vals[l-1]
		}
	}
	return def
}

func (ctx *Context) QueryArr(key string, def []string) []string {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		return vals
	}
	return def
}

func (ctx *Context) FormAll() map[string][]string {
	if ctx.req != nil {
		return ctx.req.PostForm
	}
	return map[string][]string{}
}

func (ctx *Context) FormInt(key string, def int) int {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		l := len(vals)
		if l > 0 {
			if val, err := strconv.Atoi(vals[l-1]); err == nil {
				return val
			}
		}
	}
	return def
}

func (ctx *Context) FormStr(key string, def string) string {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		l := len(vals)
		if l > 0 {
			return vals[l-1]
		}
	}
	return def
}

func (ctx *Context) FormArr(key string, def []string) []string {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		return vals
	}
	return def
}

// #endregion
