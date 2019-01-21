package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// 简化读取文件内容打印重复行及重复次数
func main() {
	counts := make(map[string]int)
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "need filename")
		return
	}
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
