// 指针、new 函数的用法
package main

import (
	"fmt"
)

func main() {
	// 指针
	x := 1
	p := &x
	fmt.Println(*p)
	*p = 2
	fmt.Println(x)

	var x1, y1 int
	fmt.Println(&x1, &y1, x1, y1, &x1 == &x1, &x1 == &y1, &x1 == nil)

	p1 := f()
	fmt.Println(*p1, p1, f(), f())

	v := 1
	fmt.Println(incr(&v), incr(&v))

	// new 函数
	p2 := new(int)
	fmt.Println(p2, *p2)
	*p2 = 2
	fmt.Println(p2, *p2)
}

func f() *int {
	v := 1
	return &v
}

func incr(p *int) int {
	*p++
	return *p
}
