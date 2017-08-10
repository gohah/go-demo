package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_dev/day9/chat/proto"
	"net"
)

func register(conn net.Conn, userId int, passwd string) (err error) {
	var msg proto.Message
	msg.Cmd = proto.UserRegister

	var registerCmd proto.RegisterCmd
	registerCmd.User.UserId = userId
	registerCmd.User.Nick = fmt.Sprintf("stu%d", userId)
	registerCmd.User.Sex = "man"
	registerCmd.User.Passwd = passwd
	registerCmd.User.Header = "http://baidu.com/header/1.jpg"

	data, err := json.Marshal(registerCmd)
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

	fmt.Println("packlen:", packLen)
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

	fmt.Println(msg)
	return
}
