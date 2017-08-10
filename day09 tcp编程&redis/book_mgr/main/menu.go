package main

import (
	"fmt"
)

func showAdminMenu() {

	fmt.Println()
	fmt.Println("1. 添加书籍")
	fmt.Println("2. 添加用户")
	fmt.Println("3. 书籍列表")
	fmt.Println("4. 返回")

	for {
		fmt.Scanf("%d", &curSel)
		switch curSel {
		case 1:
			addBook()
		case 2:
			addUser()
		case 3:
			listBook()
		case 4:
			return
		default:
		}
	}

}

func showUserMenu() {
	fmt.Println()
	fmt.Println("1. 借书")
	fmt.Println("2. 还书")
	fmt.Println("3. 已借列表")
	fmt.Println("4. 返回")

	for {
		fmt.Scanf("%d", &curSel)
		switch curSel {
		case 1:
			borrowBook()
		case 2:
			backBook()
		case 3:
			listBorrowedBook()
		case 4:
			return
		default:
		}
	}

}

func addBook() {
}

func addUser() {
}

func listBook() {
}

func borrowBook() {
}

func backBook() {
}

func listBorrowedBook() {
}
