package main

import "fmt"

type slice struct {
	ptr *[100]int
	len int
	cap int
}

func make1(s slice, cap int) slice {
	s.ptr = new([100]int)
	s.cap = cap
	s.len = 0
	return s
}

func modify(s slice) {
	s.ptr[1] = 1000
}

func testSlice2() {
	var s1 slice
	s1 = make1(s1, 10)

	s1.ptr[0] = 100
	modify(s1)

	fmt.Println(s1.ptr)
}

func testSlice() {
	var slice []int
	var arr [5]int = [...]int{1, 2, 3, 4, 5}

	slice = arr[:]
	slice = slice[1:]
	slice = slice[:len(slice)-1]
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	slice = slice[0:1]
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

}

func modify1(a []int) {
	a[1] = 1000
}

func testSlice3() {
	var b []int = []int{1, 2, 3, 4}
	modify1(b)
	fmt.Println(b)
}

func testSlice4() {
	var a = [10]int{1, 2, 3, 4}

	b := a[1:5]
	fmt.Printf("%p\n", b)
	fmt.Printf("%p\n", &a[1])
}

func main() {
	//testSlice()
	//testSlice2()
	//testSlice3()
	testSlice4()
}
