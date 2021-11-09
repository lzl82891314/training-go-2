package main

import (
	"toy-web"
	"toy-web/internal/toyserver"
)

func main() {
	builder := toyserver.CreateToyBuilder()
	builder.UseRoute("home/index", "GET", func(ctx *toy_web.Context) {
		ctx.Response("hello, world", nil)
	})
	build, err := builder.Build("testing_server")
	if err != nil {
		panic(err)
	}
	err = build.Start("localhost:8080")
	if err != nil {
		panic(err)
	}
}
