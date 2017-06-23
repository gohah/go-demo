package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"log"
	"strconv"
	"regexp"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析url传递的参数，对于POST则解析响应包的主体(request body) //注意:如果没有调用ParseForm方法，下面无法获取表单的数据
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, "")) }
	fmt.Fprintf(w, "Hello world!") //这个写入到w的是输出到客户端的
}

func login(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	fmt.Println("method:", r.Method) //获取请求的方法

	if r.Method == "POST" {

		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)

	} else {

		//if len(r.Form.Get("username"))== 0 {
		//	fmt.Println("用户名不能为空")
		//}

		//getint,err := strconv.Atoi(r.Form.Get("age"))
		////
		//if err != nil {
		//	fmt.Println("请输入纯数字")
		//}
		////
		//if getint >100 {
		//	fmt.Println("请输入正确的年龄")
		//}

		//数字
		//你想要确保一个表单输入框中获取的只能是数字，例如，你想通过表单获取某个人的具体年龄是50岁还是10岁，而不 是像“一把年纪了”或“年轻着呢”这种描述
		//if m, _ := regexp.MatchString("^[0-9]+$", r.Form.Get("age")); !m {
		//	fmt.Println("请输入正确的年龄")
		//}

		////中文
		////有时候我们想通过表单元素获取一个用户的中文名字，但是又为了保证获取的是正确的中文，我们需要进行验证，而 不是用户随便的一些输入。对于中文我们目前有效的验证只有正则方式来验证，如下代码所示
		//if m, _ := regexp.MatchString("^[\\x{4e00}-\\x{9fa5}]+$", r.Form.Get("realname")); !m {
		//	fmt.Println("请输入正确的中文名")
		//}
		//
		////英文 我们期望通过表单元素获取一个英文值，例如我们想知道一个用户的英文名，应该是astaxie，而不是asta谢。
		////我们可以很简单的通过正则验证数据:
		//if m, _ := regexp.MatchString("^[a-zA-Z]+$", r.Form.Get("engname")); !m {
		//	fmt.Println("请输入正确的英文名")
		//}
		//
		////电子邮件地址 你想知道用户输入的一个Email地址是否正确，通过如下这个方式可以验证:
		//if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, r.Form.Get("email")); !m{
		//	fmt.Println("no")
		//}else{
		//	fmt.Println("yes")
		//}
		//
		////手机号码 你想要判断用户输入的手机号码是否正确，通过正则也可以验证:
		//if m, _ := regexp.MatchString(`^(1[3|4|5|8][0-9]\d{4,8})$`, r.Form.Get("mobile")); !m {
		//	fmt.Println("请输入正确的手机号码")
		//}
		//
		////身份证号码
		////如果我们想验证表单输入的是否是身份证，通过正则也可以方便的验证，但是身份证有15位和18位，我们两个都需要 验证
		////验证15位身份证，15位的是全部数字
		//if m, _ := regexp.MatchString(`^(\d{15})$`, r.Form.Get("usercard")); !m {
		//	fmt.Println("请输入正确的身份证号码")
		//}
		//
		////验证18位身份证，18位前17位为数字，最后一位是校验位，可能为数字或字符X。
		//if m, _ := regexp.MatchString(`^(\d{17})([0-9]|X)$`, r.Form.Get("usercard")); !m {
		//	fmt.Println("请输入正确的身份证号码")
		//}


		fmt.Fprintf(w, r.Form["username"][0]) //这个写入到w的是输出到客户端的
		fmt.Println()
		fmt.Fprintf(w, r.Form["password"][0]) //这个写入到w的是输出到客户端的

	}
}

func main() {
	http.HandleFunc("/", sayhelloName) //设置访问的路由
	http.HandleFunc("/login", login) //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
