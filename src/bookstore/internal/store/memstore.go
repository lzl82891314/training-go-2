package store

import (
	"bookstore/store"
	"bookstore/store/factory"
	"sync"
)

func init() {
	factory.Register("mem", &MemStore{
		books: make(map[string]*store.Book),
	})
}

type MemStore struct {
	books map[string]*store.Book
	sync.RWMutex
}

func (mem *MemStore) Insert(book *store.Book) error {
	mem.Lock()
	defer mem.Unlock()

	if _, ok := mem.books[book.ISBN]; ok {
		return store.ErrExist
	}

	newBook := *book
	mem.books[book.ISBN] = &newBook
	return nil
}

func (mem *MemStore) Remove(key string) error {
	mem.Lock()
	defer mem.Unlock()

	if _, ok := mem.books[key]; !ok {
		return store.ErrNotFound
	}
	delete(mem.books, key)
	return nil
}

func (mem *MemStore) Modify(book *store.Book) error {
	mem.Lock()
	defer mem.Unlock()

	p, ok := mem.books[book.ISBN]
	if !ok {
		return store.ErrNotFound
	}
	cur := *p
	if book.Name != "" {
		cur.Name = book.Name
	}
	if book.Authors != nil {
		cur.Authors = book.Authors
	}
	if book.Publisher != "" {
		cur.Publisher = book.Publisher
	}
	if book.Price != 0.0 {
		cur.Price = book.Price
	}
	mem.books[book.ISBN] = &cur
	return nil
}

func (mem *MemStore) Query(key string) (store.Book, error) {
	mem.Lock()
	defer mem.Unlock()

	p, ok := mem.books[key]
	if !ok {
		return store.Book{}, store.ErrNotFound
	}
	return *p, nil
}

func (mem *MemStore) QueryAll() ([]store.Book, error) {
	mem.Lock()
	defer mem.Unlock()

	books := make([]store.Book, 0, len(mem.books))
	for _, book := range mem.books {
		books = append(books, *book)
	}
	return books, nil
}
