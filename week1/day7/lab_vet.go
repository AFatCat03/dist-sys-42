package main

import "fmt"

func main() {
	name := "Go-Cat"
	// 错误：使用 %d (整数) 打印 string
	fmt.Printf("Hello %d\n", name)
}
