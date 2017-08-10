package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	UserName string `json:"username"`
	NickName string `json:"nickname"`
	Age      int
	Birthday string
	Sex      string
	Email    string
	Phone    string
}

func testStruct() {
	user1 := &User{
		UserName: "user1",
		NickName: "上课看似",
		Age:      18,
		Birthday: "2008/8/8",
		Sex:      "男",
		Email:    "mahuateng@qq.com",
		Phone:    "110",
	}

	data, err := json.Marshal(user1)
	if err != nil {
		fmt.Printf("json.marshal failed, err:", err)
		return
	}

	fmt.Printf("%s\n", string(data))
}

func testInt() {
	var age = 100
	data, err := json.Marshal(age)
	if err != nil {
		fmt.Printf("json.marshal failed, err:", err)
		return
	}

	fmt.Printf("%s\n", string(data))
}

func testMap() {
	var m map[string]interface{}
	m = make(map[string]interface{})
	m["username"] = "user1"
	m["age"] = 18
	m["sex"] = "man"

	data, err := json.Marshal(m)
	if err != nil {
		fmt.Printf("json.marshal failed, err:", err)
		return
	}

	fmt.Printf("%s\n", string(data))
}

func testSlice() {
	var m map[string]interface{}
	var s []map[string]interface{}
	m = make(map[string]interface{})
	m["username"] = "user1"
	m["age"] = 18
	m["sex"] = "man"

	s = append(s, m)

	m = make(map[string]interface{})
	m["username"] = "user2"
	m["age"] = 29
	m["sex"] = "female"
	s = append(s, m)

	data, err := json.Marshal(s)
	if err != nil {
		fmt.Printf("json.marshal failed, err:", err)
		return
	}

	fmt.Printf("%s\n", string(data))
}

func main() {
	//testStruct()
	//testInt()
	//testMap()
	testSlice()
}
