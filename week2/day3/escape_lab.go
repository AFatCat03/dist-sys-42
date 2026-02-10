package main

import "fmt"

// 编译器判定 `x` 在函数返回后仍需被访问，因此必须将其分配在堆上
func returnPointer() *int {
	x := 1
	return &x
}

func main() {
	val := *returnPointer()

	/*
		Go 的编译器在处理闭包时非常聪明，它遵循 “最小代价原则”
		能 Copy 就 Copy：如果是基础类型 (int, bool) 且只读，直接拷贝值
		不能 Copy 就 Share：如果会被修改，或者是个大对象，或者是引用类型，那就必须通过指针共享，这时候往往就会导致逃逸
	*/
	go func() {
		val = val + 1
	}()

	fmt.Println(val) // 由于接口方法的动态性，编译器往往无法确定具体调用的实现，保守起见通常会将其分配到堆上(fmt.Println 接收的是 any, 把一个具体类型（如 int）传给 any 时，Go 运行时会发生 Boxing (装箱)，创建一个 runtime.eface 结构体)

	s := make([]int, 1024) // 尽管 Go 的栈是动态伸缩的，但大对象仍倾向于堆(或编译期无法确定大小的变量, make([]byte, n), 其中 n 是一个变量)

	fmt.Println(s)
}
