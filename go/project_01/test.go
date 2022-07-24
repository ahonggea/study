package main

import "fmt"

func main() {
	fmt.Println("hello world")
	// 1.变量的声明
	var age int
	// 2.变量的赋值
	age = 18
	// 3.变量的使用
	fmt.Println("age", age)
	// 变量和赋值可以合成一句
	var age2 int = 19
	fmt.Println("age2", age2)
	// int不赋值默认为0
	var age3 int
	fmt.Println("age3", age3)
}
