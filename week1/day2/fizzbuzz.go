package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	bound, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 1; i <= bound; i++ {
		fizzbuzz(i)
	}
}

func fizzbuzz(num int) {
	fmt.Printf("%v:", num)
	if num%3 == 0 {
		fmt.Print("Fizz")
	}
	if num%5 == 0 {
		fmt.Print("Buzz")
	}
	fmt.Println()
}

