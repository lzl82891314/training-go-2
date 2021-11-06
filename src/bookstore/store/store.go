package store

import "errors"

var (
	ErrExist    = errors.New("this book has existed")
	ErrNotFound = errors.New("can not found this book")
)

type Book struct {
	ISBN      string   `json:"ISBN"`
	Name      string   `json:"name"`
	Authors   []string `json:"authors"`
	Publisher string   `json:"publisher"`
	Price     float32  `json:"price"`
}

type Store interface {
	Insert(book *Book) error
	Remove(key string) error
	Modify(book *Book) error
	Query(key string) (Book, error)
	QueryAll() ([]Book, error)
}
