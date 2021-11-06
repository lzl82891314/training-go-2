package server

import (
	"bookstore/server/middleware"
	"bookstore/store"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func CreateStoreServer(addr string, s store.Store) *StoreServer {
	srv := &StoreServer{
		store: s,
		server: &http.Server{
			Addr: addr,
		},
	}

	router := mux.NewRouter()
	router.HandleFunc("/", srv.rootHandler)
	router.HandleFunc("/book", srv.insertHandler).Methods("POST")
	router.HandleFunc("/book/{key}", srv.removeHandler).Methods("POST")
	router.HandleFunc("/book", srv.modifyHandler).Methods("PUT")
	router.HandleFunc("/book/{key}", srv.queryHandler).Methods("GET")
	router.HandleFunc("/book", srv.queryAllHandler).Methods("GET")

	srv.server.Handler = middleware.Logging(router)
	return srv
}

type StoreServer struct {
	store  store.Store
	server *http.Server
}

func (s *StoreServer) ListenAndServe() (<-chan error, error) {
	errChan := make(chan error)
	var err error
	go func() {
		err = s.server.ListenAndServe()
		errChan <- err
	}()

	select {
	case err = <-errChan:
		return nil, err
	case <-time.After(time.Second):
		return errChan, nil
	}
}

func (s *StoreServer) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func (s StoreServer) rootHandler(w http.ResponseWriter, req *http.Request) {
	response(w, "hello, world")
}

func (s *StoreServer) insertHandler(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var book store.Book
	if err := decoder.Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := s.store.Insert(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (s *StoreServer) removeHandler(w http.ResponseWriter, req *http.Request) {
	key, ok := mux.Vars(req)["key"]
	if !ok {
		http.Error(w, fmt.Sprintf("the book not found by key: %s", key), http.StatusBadRequest)
	}
	if err := s.store.Remove(key); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (s *StoreServer) modifyHandler(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var book store.Book
	if err := decoder.Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := s.store.Modify(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (s *StoreServer) queryHandler(w http.ResponseWriter, req *http.Request) {
	key, ok := mux.Vars(req)["key"]
	if !ok {
		http.Error(w, fmt.Sprintf("the book not found by key: %s", key), http.StatusBadRequest)
		return
	}
	book, err := s.store.Query(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response(w, book)
}

func (s *StoreServer) queryAllHandler(w http.ResponseWriter, req *http.Request) {
	books, err := s.store.QueryAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response(w, books)
}

func response(w http.ResponseWriter, v interface{}) {
	data, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	if _, err := w.Write(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
