package main

import(
	"fmt"
)

//Go考虑到了这一点，所以也可以通过range，像操作slice 或者map一样操作缓存类型的channel，请看下面的例子
//for i := range c能够不断的读取channel里面的数据，直到该channel被显式的关闭。上面代码我们看到可以显 式的关闭channel，生产者通过关键字close函数关闭channel。关闭channel之后就无法再发送任何数据了，在消费 方可以通过语法v, ok := <-ch测试channel是否被关闭。如果ok返回false，那么说明channel已经没有任何数据 并且已经被关闭。
//记住应该在生产者的地方关闭channel，而不是消费的地方去关闭它，这样容易引起panic
//另外记住一点的就是channel不像文件之类的，不需要经常去关闭，只有当你确实没有任何发送数据了，或者你想显 式的结束range循环之类的

func fibonacci(n int, c chan int) {
	x,y := 1, 1

	for i:=0; i<n; i++ {
		c <- x
		x, y = y, x + y
	}
	close(c)
}

//生成斐波拉契数列
func main() {
	c:=make(chan int, 10)

	go fibonacci(cap(c),c)

	for i:=range(c) {
		fmt.Println(i)
	}
}
