package main

import (
	"fmt"
	"io"
	"os"
)

// printFile 打开指定路径的文件并将其内容流式传输到 Stdout
func printFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(os.Stdout, f)
	return err
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: printFile filename")
		os.Exit(1)
	}
	filename := os.Args[1]
	if err := printFile(filename); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}
