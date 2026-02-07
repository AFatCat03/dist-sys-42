package main

import "fmt"

func inspectSlice() {
	var s []int = make([]int, 0)
	prevCap := -1
	for i := 0; i < 10; i++ {
		if cap(s) != prevCap {
			fmt.Printf("Address of 0th element: %p, len: %d, cap: %d\n", s, len(s), cap(s))
			prevCap = cap(s)
		}
		s = append(s, 1)
	}
}

func main() {
	inspectSlice()
}