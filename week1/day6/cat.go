package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

// ---------------------------------------------------------
// 全局配置：使用 flag 包定义命令行参数
// ---------------------------------------------------------
var (
	// 定义 -n 参数，默认值为 false
	nFlag = flag.Bool("n", false, "number all output lines")
	line  = 1
)

// ---------------------------------------------------------
// Helper 1: 原始拷贝 (Fast Path)
// Day 5 的成果，用于不带 -n 的情况
// ---------------------------------------------------------
func runRawCopy(r io.Reader) error {
	_, err := io.Copy(os.Stdout, r)
	return err
}

// ---------------------------------------------------------
// Helper 2: 带行号拷贝 (Buffered Path)
// TODO: 使用 bufio 处理文本流
// ---------------------------------------------------------
func runWithLineNumbers(r io.Reader) error {
	scanner := bufio.NewScanner(r)

	// the default split function breaks the input into lines with line termination stripped.
	for scanner.Scan() {
		fmt.Printf("%6d\t%s\n", line, scanner.Text())
		line++
	}

	return scanner.Err()
}

func main() {
	flag.Parse()

	// 获取非 Flag 的参数 (即文件名列表)
	// flag.Args() 返回的是解析掉 -n 剩下的参数
	filenames := flag.Args()

	// 逻辑分发：如果没有文件参数，读 Stdin
	if len(filenames) == 0 {
		processStream(os.Stdin)
		return
	}

	// 遍历文件
	for _, fname := range filenames {
		if fname == "-" {
			processStream(os.Stdin)
		} else {
			f, err := os.Open(fname)
			if err != nil {
				fmt.Fprintf(os.Stderr, "go-cat: %s: %v\n", fname, err)
				continue
			}
			// 处理文件流
			processStream(f)
			f.Close()
		}
	}
}

// processStream 根据全局 nFlag 决定调用哪个处理函数
func processStream(r io.Reader) {
	var err error

	if *nFlag {
		err = runWithLineNumbers(r)
	} else {
		err = runRawCopy(r)
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
