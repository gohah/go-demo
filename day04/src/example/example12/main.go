package main

import (
	"sort"
	"fmt"
)
//sort包排序，查找
func testSortInt() {
	a:=[...]int{6,2,8,3,1,5,10}

	sort.Ints(a[:])

	fmt.Println(a)
}

func testSortString() {
	a:=[...]string{"A","a","z","c"}

	sort.Strings(a[:])

	fmt.Println(a)
}


func main() {
	testSortInt()
	testSortString()
}
