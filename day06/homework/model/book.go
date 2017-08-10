package model

import "errors"

var (
	ErrStockNotEnough = errors.New("stock is not enough")
)

type Book struct {
	Name       string
	Total      int
	Author     string
	CreateTime string
}

func CreateBook(name string, total int, author, createTime string) (b *Book) {
	b = &Book{
		Name:       name,
		Total:      total,
		Author:     author,
		CreateTime: createTime,
	}

	return
}

func (b *Book) canBorrow(c int) bool {
	return b.Total >= c
}

func (b *Book) Borrow(c int) (err error) {
	if b.canBorrow(c) == false {
		err = ErrStockNotEnough
		return
	}

	b.Total -= c
	return
}

func (b *Book) Back(c int) (err error) {

	b.Total += c
	return
}
