package main

import "fmt"

//map的值是map
func main() {
	var m map[string]map[string]string

	m = make(map[string]map[string]string,100)

	m["key1"] = make(map[string]string)

	m["key1"]["key2"] = "hahah1"
	m["key1"]["key3"] = "hahah2"
	m["key1"]["key4"] = "hahah3"
	m["key1"]["key5"] = "hahah4"
	m["key1"]["key6"] = "hahah5"
	m["key1"]["key7"] = "hahah6"

	m["key2"] = make(map[string]string)

	m["key2"]["key2"] = "hahah1"
	m["key2"]["key3"] = "hahah2"
	m["key2"]["key4"] = "hahah3"
	m["key2"]["key5"] = "hahah4"
	m["key2"]["key6"] = "hahah5"
	m["key2"]["key7"] = "hahah6"

	fmt.Println(m)

	val,ok := m["key1"]

	if ok {
		fmt.Println(val)
	}

	for k,v := range(m) {
		fmt.Println(k)
		for k1,v1 := range(v) {
			fmt.Println(k1,v1)
		}
	}
}
