package main

import (
	"fmt"
	"time"
)

func main() {
	c := 0
	// 启动一个 goroutine 修改 c
	go func() {
		c = 1
	}()
	// 主线程读取 c
	// 两个线程同时访问内存，且没有加锁 -> 数据竞争！
	if c == 0 {
		fmt.Println("C is 0")
	}
	time.Sleep(time.Second)
}
