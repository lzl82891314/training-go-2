package main

import (
	"fmt"
	tw "toy-web"
	_ "toy-web/internal/toyrouter/v1"
	"toy-web/internal/toyserver/builder"
)

func main() {
	builder.Use(func(next tw.HandlerFunc) tw.HandlerFunc {
		return func(ctx *tw.Context) {
			fmt.Println("hello middleware")
			next(ctx)
		}
	})
	server, err := builder.Build("testing_server")
	if err != nil {
		panic(err)
	}

	server.Map("/", "GET", func(ctx *tw.Context) {
		ctx.Response("hello, world", nil)
	})
	err = server.Start("localhost:8080")
	if err != nil {
		panic(err)
	}
}
