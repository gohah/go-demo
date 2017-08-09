package model

import (
	"errors"
	"time"
)

const (
	RoleAdmin = 1
	RoleUser  = 2
)

var (
	ErrNotFoundSn = errors.New("Not Found Found Sn")
)

type BookRecord struct {
	start time.Time
	end   time.Time
	days  int
	book  *Book
}

type User struct {
	username    string
	passwd      string
	grade       string
	sex         string
	age         int
	id          string
	role        int
	bookRecords map[string]*BookRecord
}

func NewUser(username, passwd, grade, sex string, age int, id string) *User {
	return &User{
		username:    username,
		passwd:      passwd,
		grade:       grade,
		sex:         sex,
		age:         age,
		id:          id,
		bookRecords: make(map[string]*BookRecord, 16),
	}
}

func (p *User) SetAdmin() {
	p.role = RoleAdmin
}

func (p *User) IsAdmin() bool {
	return p.role == RoleAdmin
}

func (p *User) GetBookRecords() map[string]*BookRecord {
	return p.bookRecords
}

func (p *User) BorrowBook(book *Book, interval time.Duration) {

	now := time.Now()
	record := &BookRecord{
		start: now,
		end:   now.Add(interval),
		book:  book,
	}
	p.bookRecords[book.Sn] = record
}

func (p *User) BackBook(sn string) (err error) {
	_, ok := p.bookRecords[sn]
	if !ok {
		err = ErrNotFoundSn
		return
	}

	delete(p.bookRecords, sn)
	return
}
