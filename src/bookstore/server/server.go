package server

import (
	"bookstore/server/middleware"
	"bookstore/store"
	"context"
	"encoding/json"
	"errors"
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
	router.HandleFunc("/*", srv.rootHandler).Methods("GET").Methods("POST")
	router.NotFoundHandler = http.HandlerFunc(srv.notfoundHandler)
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

func (s *StoreServer) rootHandler(w http.ResponseWriter, req *http.Request) {
	response(w, "hello, world", nil)
}

func (s *StoreServer) notfoundHandler(w http.ResponseWriter, req *http.Request) {
	resp := &ResponseDto{
		Code:      http.StatusNotFound,
		Message:   "service not found",
		Data:      nil,
		Timestamp: time.Now().UnixMilli(),
	}
	doResponse(w, resp)
}

func (s *StoreServer) insertHandler(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var book store.Book
	if err := decoder.Decode(&book); err != nil {
		response(w, nil, err)
		return
	}

	err := s.store.Insert(&book)
	response(w, "insert ok", err)
}

func (s *StoreServer) removeHandler(w http.ResponseWriter, req *http.Request) {
	key, ok := mux.Vars(req)["key"]
	if !ok {
		response(w, nil, errors.New(fmt.Sprintf("the book not found by key: %s", key)))
		return
	}
	err := s.store.Remove(key)
	response(w, "remove ok", err)
}

func (s *StoreServer) modifyHandler(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var book store.Book
	if err := decoder.Decode(&book); err != nil {
		response(w, nil, err)
		return
	}
	err := s.store.Modify(&book)
	response(w, "modify ok", err)
}

func (s *StoreServer) queryHandler(w http.ResponseWriter, req *http.Request) {
	key, ok := mux.Vars(req)["key"]
	if !ok {
		response(w, nil, errors.New(fmt.Sprintf("the book not found by key: %s", key)))
		return
	}
	book, err := s.store.Query(key)
	response(w, book, err)
}

func (s *StoreServer) queryAllHandler(w http.ResponseWriter, req *http.Request) {
	books, err := s.store.QueryAll()
	response(w, books, err)
}

type ResponseDto struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	Timestamp int64       `json:"timestamp"`
}

func response(w http.ResponseWriter, v interface{}, err error) {
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
	doResponse(w, resp)
}

func doResponse(w http.ResponseWriter, resp *ResponseDto) {
	data, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
