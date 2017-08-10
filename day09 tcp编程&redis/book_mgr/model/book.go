package model

import (
	"errors"
)

var (
	ErrInvalidStock      = errors.New("Invalid Stock")
	ErrInvalidSn         = errors.New("Invalid Sn")
	ErrInvalidBookStatus = errors.New("Invalid book status")
	ErrAlreadyBorrowed   = errors.New("book already borrowed")
)

const (
	ItemFlagBorrowed = 1
	ItemFlagFree     = 2
)

type Book struct {
	Sn   string
	Flag int

	Name    string
	Publish string
	Date    string
	Author  string
}

func NewBook(sn, name, publish, date, author string) *Book {
	return &Book{
		Sn:      sn,
		Flag:    ItemFlagFree,
		Name:    name,
		Publish: publish,
		Date:    date,
		Author:  author,
	}
}

func (p *Book) Borrow() (sn string, err error) {
	if p.Flag == ItemFlagBorrowed {
		err = ErrAlreadyBorrowed
		return
	}

	p.Flag = ItemFlagBorrowed
	return
}

func (p *Book) Back() (err error) {

	if p.Flag != ItemFlagBorrowed {
		err = ErrInvalidBookStatus
		return
	}

	p.Flag = ItemFlagFree
	return
}
