package main

import (
	"fmt"
	"io"
	"os"
)

// runCopy 是一个纯粹的逻辑函数。
// 它不关心数据来自磁盘还是键盘，它只关心 input 满足 io.Reader 接口。
// 这里的 'r' 既可以是 *os.File，也可以是 os.Stdin。
func runCopy(r io.Reader) error {
	_, err := io.Copy(os.Stdout, r)
	return err
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		if err := runCopy(os.Stdin); err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(1)
		}
	}

	for _, filename := range args {
		if filename == "-" {
			if err := runCopy(os.Stdin); err != nil {
				fmt.Fprint(os.Stderr, err)
				//os.Exit(1)
				continue
			}
		} else {

			f, err := os.Open(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "go-cat: %s: %v\n", filename, err)
				continue
			}

			if err = runCopy(f); err != nil {
				fmt.Fprint(os.Stderr, err)
				os.Exit(1)
			}

			// 显式关闭
			f.Close()
		}
	}
}
