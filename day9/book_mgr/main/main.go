package main

import (
	"fmt"
	"model"
)

var (
	curPos int = TopMenu
	curSel int
	exit   bool
	mgr    = model.NewBookMgr()
	user   *model.User
)

const (
	TopMenu = 1
)

func mainMenu() {
	fmt.Println()
	fmt.Println("1. 管理员登陆")
	fmt.Println("2. 用户登陆")
	fmt.Println("3. 退出\n\n")

	fmt.Scanf("%d", &curSel)
	switch curSel {
	case 1:
		u, err := adminLogin()
		if err != nil {
			fmt.Printf("admin login failed, err:%v\n", err)
			return
		}
		user = u
		showAdminMenu()
	case 2:
		u, err := userLogin()
		if err != nil {
			fmt.Printf("login failed, err:%v\n", err)
			return
		}
		user = u
		showUserMenu()
	case 3:
		exit = true
		return
	}
}

func main() {
	for !exit {
		if curPos == TopMenu {
			mainMenu()
		}
	}
}
