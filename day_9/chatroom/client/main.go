package main

import (
	"fmt"
	"os"
)

//定义两个变量，一个表示用户id，一个表示用户密码
var userId int
var userPwd string

func main() {
	//接收用户的选择
	var key int
	//判断是否还继续显示菜单
	var loop = true

	for loop {
		fmt.Println("-----------欢迎登陆多人聊天系统-------------")
		fmt.Println("\t\t\t 1 登录聊天室")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("\t\t\t 请选择(1-3)")

		fmt.Scanf("%d\n", &key)

		switch key {
			case 1 :
				fmt.Println("登录聊天室")
				loop = false
			case 2 :
				fmt.Println("注册用户")
				loop = false
			case 3 :
				fmt.Println("退出系统")
				os.Exit(0)
			default:
				fmt.Println("您的输入有误，请重新输出...")
		}
	}
	//增加用户输入，显示新的信息
	if key == 1 {
		//说明用户要登录
		fmt.Println("请输入用户id")
		fmt.Scanf("%d\n", &userId)
		fmt.Println("请输入用户密码")
		fmt.Scanf("%d\n", &userPwd)
		//先把登录函数写在另一个文件，比如login.go
		login(userId, userPwd)
	} else if key ==2 {
		fmt.Println("进行用户注册的逻辑...")
	}
}