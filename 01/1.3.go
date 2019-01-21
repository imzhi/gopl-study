package main

import (
	"bufio"
	"fmt"
	"os"
)

// 打印重复行及重复次数
func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		text := input.Text()
		if text == "end" {
			break
		}
		counts[text]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
