package main

import (
	_ "bookstore/internal/store"
	"bookstore/server"
	"bookstore/store/factory"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	s, err := factory.Create("mem")
	if err != nil {
		panic(err)
	}
	address := ":8080"
	srv := server.CreateStoreServer(address, s)

	errChan, err := srv.ListenAndServe()
	if err != nil {
		log.Println("web start failed with err:", err)
		return
	}
	log.Printf("web server start ok on address: %s\n", address)

	sglChan := make(chan os.Signal, 1)
	signal.Notify(sglChan, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err = <-errChan:
		log.Println("web server run failed: ", err)
		return
	case <-sglChan:
		{
			log.Println("bookstore is exiting...")
			ctx, cf := context.WithTimeout(context.Background(), time.Second)
			defer cf()
			err = srv.Shutdown(ctx)
		}
		if err != nil {
			log.Println("bookstore exit error: ", err)
			return
		}
		log.Println("bookstore exit ok")
	}
}
