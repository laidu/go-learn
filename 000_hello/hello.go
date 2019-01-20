package main

import "fmt"

/**
	运行方式:
		1, go build hello.gc
		2, ./hello
	建议方式: go run hello.gc

	注意事项:
		1, 源文件以.go结尾
		2, 包含未使用的import时,代码无法通过编译
		3, 程序执行入口为 main
		4, 行尾不需要;
		5, 按行变异
 */
func main() {
	fmt.Println("hello world")
}
