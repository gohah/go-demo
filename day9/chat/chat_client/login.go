package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_dev/day9/chat/common"
	"go_dev/day9/chat/proto"
	"net"
	"os"
)

func login(conn net.Conn, userId int, passwd string) (err error) {
	var msg proto.Message
	msg.Cmd = proto.UserLogin

	var loginCmd proto.LoginCmd
	loginCmd.Id = userId
	loginCmd.Passwd = passwd

	data, err := json.Marshal(loginCmd)
	if err != nil {
		return
	}

	msg.Data = string(data)
	data, err = json.Marshal(msg)
	if err != nil {
		return
	}

	var buf [4]byte
	packLen := uint32(len(data))

	//fmt.Println("packlen:", packLen)
	binary.BigEndian.PutUint32(buf[0:4], packLen)

	n, err := conn.Write(buf[:])
	if err != nil || n != 4 {
		fmt.Println("write data  failed")
		return
	}

	_, err = conn.Write([]byte(data))
	if err != nil {
		return
	}

	msg, err = readPackage(conn)
	if err != nil {
		fmt.Println("read package failed, err:", err)
	}

	var loginResp proto.LoginCmdRes
	json.Unmarshal([]byte(msg.Data), &loginResp)
	if loginResp.Code == 500 {
		fmt.Println("user not register, start register")
		register(conn, userId, passwd)
		os.Exit(0)
	}

	//fmt.Println("online user list:")
	for _, v := range loginResp.User {
		if v == userId {
			continue
		}
		fmt.Println("user logined:", v)
		user := &common.User{UserId: v}
		onlineUserMap[user.UserId] = user
	}
	return
}
