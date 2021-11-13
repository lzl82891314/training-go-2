package main

import (
	"fmt"
	tw "toy-web"
	_ "toy-web/internal/toyrouter/v1"
	"toy-web/internal/toyserver"
)

func main() {
	toyserver.Use(func(next tw.Action) tw.Action {
		return func(ctx *tw.Context) {
			fmt.Println("hello middleware")
			next(ctx)
		}
	})
	server, err := toyserver.Build("testing_server")
	if err != nil {
		panic(err)
	}

	server.Map("/", "GET", func(ctx *tw.Context) {
		ctx.Ok("hello, world")
	})
	server.Map("hello/*", "GET", func(ctx *tw.Context) {
		ctx.Ok("hello, *")
	})
	server.Map("hello/jeffery", "GET", func(ctx *tw.Context) {
		ctx.Ok("hello, jeffery")
	})
	err = server.Start("localhost:8080")
	if err != nil {
		panic(err)
	}
}
