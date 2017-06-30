package main

import "fmt"

func testSlice() {
	var a = [5]int{1,2,3,4,5}
	s := a[1:4]
	fmt.Printf("s addr %p, a[1] addr %p, lenght=%d,cap=%d",s, &a[1],len(s),cap(s))
	fmt.Println(s)

	s=append(s,10)
	fmt.Printf("s addr %p, a[1] addr %p, lenght=%d,cap=%d",s, &a[1],len(s),cap(s))
	fmt.Println(s)

	s=append(s,10)
	s=append(s,10)
	s=append(s,10)
	s=append(s,10)
	s=append(s,10)
	s=append(s,10)
	s=append(s,10)
	s=append(s,10)
	s=append(s,10)
	s=append(s,10)
	fmt.Printf("s addr %p, a[1] addr %p, lenght=%d,cap=%d",s, &a[1],len(s),cap(s))
	fmt.Println(s)


}

func main() {
	testSlice()
}
