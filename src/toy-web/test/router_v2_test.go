package test

import (
	"testing"
	tw "toy-web"
	"toy-web/internal/toyrouter/factory"
	_ "toy-web/internal/toyrouter/v2"
)

func TestTreeGenerator(t *testing.T) {
	router, _ := factory.New("v2")
	path := "/home/hello/world/jeffery/lee"
	router.Map(path, "GET", func(ctx *tw.Context) {
		ctx.Response("hello, world", nil)
	})
	_, b := router.Find(path, "GET")
	if !b {
		t.Error(router)
	}
}
