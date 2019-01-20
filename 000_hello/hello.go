package main

import "fmt"

/*
	运行方式:
		1, go build hello.gc
		2, ./hello
	建议方式: go run hello.gc

	注意事项:
		1, 源文件以.go结尾
		2, 包含未使用的import时,代码无法通过编译
		3, 程序执行入口为 main
		4, 行尾不需要;
		5, 按行变异s
 */
func main() {

	// print
	fmt.Println("hello world")

	// \r 回车, 从当前行最前面开始输出, 不换行
	fmt.Println("张三, 你好 \r李四")

}
