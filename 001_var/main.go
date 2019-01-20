package main

import (
	"fmt"
)


// 定义全局变量(推荐方式)
var (
	g_a = 123
)

/*
	go 变量定义
 */
func main() {

	fmt.Println(g_a)

	// 使用 var [变量名] [变量类型]
	var i float32 = 10
	fmt.Println("i= ",i)

	// 根据值类型 使用 类型推导 var [变量名] = [value]
	var  j = 10;
	fmt.Println("j= ",j)

	// 省略var, [变量名] := [value]
	name := "xiaoming"
	fmt.Println(name)

	// 批量声明同种类型变量
	var a, b, c  = true,true,false
	// 等价于
	//var a, b, c bool = true,true,false
	fmt.Println(a,b,c)

	// 批量声明多种类型变量
	d,e,f := 0,1.2,"hello"
	fmt.Println(d,e,f)


}
