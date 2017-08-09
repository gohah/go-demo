package main

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"go_dev/day9/chat/proto"
	"net"
)

type Client struct {
	conn   net.Conn
	userId int
	buf    [8192]byte
}

func (p *Client) readPackage() (msg proto.Message, err error) {

	n, err := p.conn.Read(p.buf[0:4])
	if n != 4 {
		err = errors.New("read header failed")
		return
	}
	fmt.Println("read package:", p.buf[0:4])

	var packLen uint32
	packLen = binary.BigEndian.Uint32(p.buf[0:4])

	fmt.Printf("receive len:%d", packLen)
	n, err = p.conn.Read(p.buf[0:packLen])
	if n != int(packLen) {
		err = errors.New("read body failed")
		return
	}

	fmt.Printf("receive data:%s\n", string(p.buf[0:packLen]))
	err = json.Unmarshal(p.buf[0:packLen], &msg)
	if err != nil {
		fmt.Println("unmarshal failed, err:", err)
	}
	return
}

func (p *Client) writePackage(data []byte) (err error) {

	packLen := uint32(len(data))

	binary.BigEndian.PutUint32(p.buf[0:4], packLen)
	n, err := p.conn.Write(p.buf[0:4])
	if err != nil {
		fmt.Println("write data  failed")
		return
	}

	n, err = p.conn.Write(data)
	if err != nil {
		fmt.Println("write data  failed")
		return
	}

	if n != int(packLen) {
		fmt.Println("write data  not finished")
		err = errors.New("write data not fninshed")
		return
	}

	return
}

func (p *Client) Process() (err error) {

	for {
		var msg proto.Message
		msg, err = p.readPackage()
		if err != nil {
			//	fmt.Println("")
			clientMgr.DelClient(p.userId)
			//TODO:通知所有在线用户，该用户已经下线
			return err
		}

		err = p.processMsg(msg)
		if err != nil {
			fmt.Println("process msg failed, err:", err)
			continue
			//return
		}
	}
}

func (p *Client) processMsg(msg proto.Message) (err error) {

	switch msg.Cmd {
	case proto.UserLogin:
		err = p.login(msg)
	case proto.UserRegister:
		err = p.register(msg)
	case proto.UserSendMessageCmd:
		err = p.proccessUserSendMessage(msg)
	default:
		err = errors.New("unsupport message")
		return
	}
	return
}

func (p *Client) SendMessageToUser(userId int, text string) {

	var respMsg proto.Message
	respMsg.Cmd = proto.UserRecvMessageCmd

	var recvMsg proto.UserRecvMessageReq
	recvMsg.UserId = userId
	recvMsg.Data = text

	data, err := json.Marshal(recvMsg)
	if err != nil {
		fmt.Println("marshal failed, ", err)
		return
	}

	respMsg.Data = string(data)
	data, err = json.Marshal(respMsg)
	if err != nil {
		fmt.Println("marshal failed, ", err)
		return
	}
	err = p.writePackage(data)
	if err != nil {
		fmt.Println("send failed, ", err)
		return
	}
}

func (p *Client) proccessUserSendMessage(msg proto.Message) (err error) {
	var userReq proto.UserSendMessageReq
	err = json.Unmarshal([]byte(msg.Data), &userReq)
	if err != nil {
		fmt.Println("unmarshal failed, err:", err)
		return
	}

	users := clientMgr.GetAllUsers()
	for id, client := range users {
		if id == userReq.UserId {
			continue
		}

		client.SendMessageToUser(userReq.UserId, userReq.Data)
	}
	return
}

func (p *Client) loginResp(err error) {
	var respMsg proto.Message
	respMsg.Cmd = proto.UserLoginRes

	var loginRes proto.LoginCmdRes
	loginRes.Code = 200

	userMap := clientMgr.GetAllUsers()
	for userId, _ := range userMap {
		loginRes.User = append(loginRes.User, userId)
	}

	if err != nil {
		loginRes.Code = 500
		loginRes.Error = fmt.Sprintf("%v", err)
	}

	data, err := json.Marshal(loginRes)
	if err != nil {
		fmt.Println("marshal failed, ", err)
		return
	}

	respMsg.Data = string(data)
	data, err = json.Marshal(respMsg)
	if err != nil {
		fmt.Println("marshal failed, ", err)
		return
	}
	err = p.writePackage(data)
	if err != nil {
		fmt.Println("send failed, ", err)
		return
	}
}

func (p *Client) login(msg proto.Message) (err error) {
	defer func() {
		p.loginResp(err)
	}()

	fmt.Printf("recv user login request, data:%v", msg)
	var cmd proto.LoginCmd
	err = json.Unmarshal([]byte(msg.Data), &cmd)
	if err != nil {
		fmt.Println("unmarshal failed, err:", err)
		return
	}

	_, err = mgr.Login(cmd.Id, cmd.Passwd)
	if err != nil {
		return
	}

	clientMgr.AddClient(cmd.Id, p)
	p.userId = cmd.Id

	p.NotifyOthersUserOnline(cmd.Id)
	return
}

func (p *Client) NotifyOthersUserOnline(userId int) {
	users := clientMgr.GetAllUsers()
	for id, client := range users {
		if id == userId {
			continue
		}

		client.NotifyUserOnline(userId)
	}
}

func (p *Client) NotifyUserOnline(userId int) {

	var respMsg proto.Message
	respMsg.Cmd = proto.UserStatusNotifyRes

	var noitfyRes proto.UserStatusNotify
	noitfyRes.UserId = userId
	noitfyRes.Status = proto.UserOnline

	data, err := json.Marshal(noitfyRes)
	if err != nil {
		fmt.Println("marshal failed, ", err)
		return
	}

	respMsg.Data = string(data)
	data, err = json.Marshal(respMsg)
	if err != nil {
		fmt.Println("marshal failed, ", err)
		return
	}
	err = p.writePackage(data)
	if err != nil {
		fmt.Println("send failed, ", err)
		return
	}
}

func (p *Client) register(msg proto.Message) (err error) {
	var cmd proto.RegisterCmd
	err = json.Unmarshal([]byte(msg.Data), &cmd)
	if err != nil {
		return
	}

	err = mgr.Register(&cmd.User)
	if err != nil {
		return
	}

	return
}
