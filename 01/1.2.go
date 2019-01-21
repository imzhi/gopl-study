package main

import (
    "fmt"
    "os"
    "strings"
)

// 打印命令行参数
func main() {
    // 2)
    s, sep := "", ""
    for _, arg := range(os.Args[1:]) {
        s += sep + arg
        sep = " "
    }
    fmt.Println(s)

    // 3)
    fmt.Println(strings.Join(os.Args[1:], " "))

    fmt.Println(os.Args[1:])
}
