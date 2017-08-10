package model

import (
	"errors"
	"time"
)

var (
	ErrNotFoundUser = errors.New("Not Found User")
	ErrNotFoundBook = errors.New("Not Found User")
)

type BookMgr struct {
	//userid --> User
	users      map[string]*User
	adminUsers map[string]*User
	//book sn ---> Book
	books map[string]*Book
	//all books
	allBooks []*Book
	allUsers []*User
	allAdmin []*User
}

func NewBookMgr() *BookMgr {
	mgr := &BookMgr{
		users:      make(map[string]*User),
		adminUsers: make(map[string]*User),
		books:      make(map[string]*Book),
	}

	user := NewUser("admin", "admin", "1", "man", 28, "273")
	user.SetAdmin()

	mgr.adminUsers["admin"] = user
	mgr.allAdmin = append(mgr.allAdmin, user)
	return mgr
}

func (p *BookMgr) Borrow(id, sn string, interval time.Duration) (book *Book, err error) {
	u, ok := p.users[id]
	if !ok {
		err = ErrNotFoundUser
		return
	}

	book, ok = p.books[sn]
	if !ok {
		err = ErrNotFoundBook
		return
	}

	_, err = book.Borrow()
	if err != nil {
		return
	}

	u.BorrowBook(book, interval)
	return
}

func (p *BookMgr) BackBook(id, sn string) (err error) {

	u, ok := p.users[id]
	if !ok {
		err = ErrNotFoundUser
		return
	}

	book, ok := p.books[sn]
	if !ok {
		err = ErrNotFoundBook
		return
	}

	err = u.BackBook(sn)
	if err != nil {
		return
	}

	err = book.Back()
	if err != nil {
		return
	}

	return
}

func (p *BookMgr) ListAllBook(offset, size int) (books []*Book, err error) {
	if offset < 0 || size <= 0 || offset >= len(p.users) {
		err = errors.New("invalid parameter, offset and size")
		return
	}

	books = p.allBooks[offset : offset+size]
	return
}

func (p *BookMgr) AddBook(book *Book) (err error) {
	if book == nil {
		err = errors.New("invalid book, nil")
		return
	}

	p.allBooks = append(p.allBooks, book)
	p.books[book.Sn] = book
	return
}

func (p *BookMgr) AddUser(user *User) (err error) {
	if user == nil {
		err = errors.New("invalid user")
		return
	}

	p.users[user.id] = user
	return
}

func (p *BookMgr) ListUser(offset, size int) (users []*User, err error) {
	if offset < 0 || size <= 0 || offset >= len(p.users) {
		err = errors.New("invalid parameter, offset and size")
		return
	}

	users = p.allUsers[offset : offset+size]
	return
}

func (p *BookMgr) UserLogin(username, passwd string) (user *User, err error) {

	v, ok := p.users[username]
	if !ok {
		err = errors.New("user not exists")
		return
	}

	if v.username != username || v.passwd != passwd {
		err = errors.New("username or passwd not right")
		return
	}

	user = v
	return
}

func (p *BookMgr) AdminLogin(username, passwd string) (user *User, err error) {

	v, ok := p.adminUsers[username]
	if !ok {
		err = errors.New("user not exists")
		return
	}

	if v.username != username || v.passwd != passwd {
		err = errors.New("username or passwd not right")
		return
	}

	if v.IsAdmin() == false {
		err = errors.New("Not Admin User")
		return
	}

	user = v
	return
}
