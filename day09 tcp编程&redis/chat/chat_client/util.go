package main

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"go_dev/day9/chat/proto"
	"net"
)

func readPackage(conn net.Conn) (msg proto.Message, err error) {
	var buf [8192]byte
	n, err := conn.Read(buf[0:4])
	if n != 4 {
		err = errors.New("read header failed")
		return
	}
	//fmt.Println("read package:", buf[0:4])

	var packLen uint32
	packLen = binary.BigEndian.Uint32(buf[0:4])

	//fmt.Printf("receive len:%d", packLen)
	n, err = conn.Read(buf[0:packLen])
	if n != int(packLen) {
		err = errors.New("read body failed")
		return
	}

	//fmt.Printf("receive data:%s\n", string(buf[0:packLen]))
	err = json.Unmarshal(buf[0:packLen], &msg)
	if err != nil {
		fmt.Println("unmarshal failed, err:", err)
	}
	return
}
