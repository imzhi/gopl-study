package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// 命令行参数网址输出页面源码
// 练习 1.7 io.Copy(dst, src)
// 练习 1.8 自动给 url 加上 http://
// 练习 1.9 打印出状态码
func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		_, err2 := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err2 != nil {
			fmt.Fprintf(os.Stderr, "fetch: copy %v\n", err2)
			os.Exit(1)
		}

		io.Copy(os.Stdout, strings.NewReader(fmt.Sprintf("status: %s\n", resp.Status)))
	}
}
