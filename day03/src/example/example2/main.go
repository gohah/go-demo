package main

import (
	"strings"
	"fmt"
	"strconv"
)

func main() {

	//字符串替换
	str := strings.Replace("hello,world","ll","hh",1)
	fmt.Println(str)

	//字符串出现次数
	count := strings.Count("abcdabacabcdabcd","abc")
	fmt.Println(count)

	//字符串重复
	repeat := strings.Repeat("abc",3)
	fmt.Println(repeat)

	//字符串转换成小写
	lower := strings.ToLower("ABCDEefaadaSDAFDAFDAFD")
	fmt.Println(lower)

	//字符串转换成大写
	upper := strings.ToUpper("ABCDEefaadaSDAFDAFDAFD")
	fmt.Println(upper)

	//去掉字符串首尾空格
	trim := strings.TrimSpace("    hahahahha   ")
	fmt.Println(trim)

	//去掉字符串首尾cut
	trim1 := strings.Trim("    hahahahha   "," ")
	fmt.Println(trim1)

	//去掉字符串左侧cut
	trim2 := strings.TrimLeft("    hahahahha   "," ")
	fmt.Println(trim2)

	//去掉字符串右侧cut
	trim3 := strings.TrimRight("    hahahahha   "," ")
	fmt.Println(trim3)

	//返回空格分隔的所有子串的slice
	splite := strings.Fields("a b c d e f g")
	fmt.Println(splite)

	//返回sep分隔的所有子串的slice
	splite2 := strings.Split("axbxcxdxexfxg","x")
	fmt.Println(splite2)

	//用sep把数组中的元素连接起来
	join := strings.Join([]string{"1","2","3","4","5"},",")
	fmt.Println(join)


        //数字转换成字符串
	str2 := strconv.Itoa(123)
	fmt.Println(str2)

	//字符串转换成数字
	str3,_ := strconv.Atoi("1234567")
	fmt.Println(str3)

}
