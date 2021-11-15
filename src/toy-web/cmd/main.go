package main

import (
	"fmt"
	"time"
	tw "toy-web"
	_ "toy-web/internal/toyrouter/v1"
	_ "toy-web/internal/toyrouter/v2"
	_ "toy-web/internal/toyrouter/v3"
	"toy-web/internal/toyserver"
)

func main() {

	server, err := toyserver.New("toy")
	if err != nil {
		panic(err)
	}
	server.Use(func(next tw.Action) tw.Action {
		return func(ctx tw.IContext) {
			fmt.Println("hello middleware1")
			next(ctx)
		}
	})

	server.Use(func(next tw.Action) tw.Action {
		return func(ctx tw.IContext) {
			fmt.Println("hello middleware2")
			next(ctx)
		}
	})

	server.Map("/", "GET", func(ctx tw.IContext) {
		ctx.Ok("hello, world")
	})
	server.Map("hello/*", "GET", func(ctx tw.IContext) {
		ctx.Ok("hello, *")
	})
	server.Map("hello/jeffery", "GET", func(ctx tw.IContext) {
		ctx.Ok("hello, jeffery")
	})
	server.Map("mid/test", "GET", func(ctx tw.IContext) {
		time.Sleep(time.Second * 10)
		ctx.Ok("sleep over")
	})
	err = server.Start("localhost:8080")
	if err != nil {
		panic(err)
	}
}
