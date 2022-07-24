package main

import "fmt"

// 全局变量
var n7 = 100
var n8 = 0.1 * 0.2

// 全局变量可以一次性声明
var (
	n9  = 10
	n10 = 1.111
)

func main() {
	// 局部变量
	// 第一种声明变量并赋值
	var num int = 18
	fmt.Println(num)
	// 第二种声明变量但是不复制，使用默认值
	var num2 int
	var str2 string
	fmt.Println(num2, str2)
	// 第三种没有写变量类型根据=后面的值自动判定变量类型
	var num3 = 123
	fmt.Println(num3)
	// 第四种用:=省略var
	num4 := 456
	fmt.Println(num4)

	fmt.Println("---------------------------------------------")
	// 声明多个变量
	var n1, n2, n3 int
	fmt.Println(n1, n2, n3)
	var n4, n5, n6 = 10, "jack", 8.8888 / 2
	fmt.Println(n4, n5, n6)
	fmt.Println(n7, n8)
	fmt.Println(n9, n10)
}
