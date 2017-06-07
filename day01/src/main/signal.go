package main


func test_signal(a int, b int) {
	sum := a + b
	pipe <- sum
}
