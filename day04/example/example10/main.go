package main

import "fmt"

func testMap() {
	var a map[string]string = map[string]string{
		"key": "value",
	}
	//a := make(map[string]string, 10)
	a["abc"] = "efg"
	a["abc"] = "efg"
	a["abc1"] = "efg"

	fmt.Println(a)
}
func testMap2() {

	a := make(map[string]map[string]string, 100)
	a["key1"] = make(map[string]string)
	a["key1"]["key2"] = "abc"
	a["key1"]["key3"] = "abc"
	a["key1"]["key4"] = "abc"
	a["key1"]["key5"] = "abc"
	fmt.Println(a)

}
func modify(a map[string]map[string]string) {
	_, ok := a["zhangsan"]
	if !ok {
		a["zhangsan"] = make(map[string]string)
	}

	a["zhangsan"]["passwd"] = "123456"
	a["zhangsan"]["nickname"] = "pangpang"

	return
}
func testMap3() {

	a := make(map[string]map[string]string, 100)

	modify(a)

	fmt.Println(a)
}

func trans(a map[string]map[string]string) {
	for k, v := range a {
		fmt.Println(k)
		for k1, v1 := range v {
			fmt.Println("\t", k1, v1)
		}
	}
}

func testMap4() {

	a := make(map[string]map[string]string, 100)
	a["key1"] = make(map[string]string)
	a["key1"]["key2"] = "abc"
	a["key1"]["key3"] = "abc"
	a["key1"]["key4"] = "abc"
	a["key1"]["key5"] = "abc"

	a["key2"] = make(map[string]string)
	a["key2"]["key2"] = "abc"
	a["key2"]["key3"] = "abc"

	trans(a)
	delete(a, "key1")
	fmt.Println()
	trans(a)

	fmt.Println(len(a))
}

func testMap5() {
	var a []map[int]int
	a = make([]map[int]int, 5)

	if a[0] == nil {
		a[0] = make(map[int]int)
	}
	a[0][10] = 10
	fmt.Println(a)
}

func main() {
	//testMap()
	//testMap2()
	//testMap3()
	//testMap4()
	testMap5()
}
