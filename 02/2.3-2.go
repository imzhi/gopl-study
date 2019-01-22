package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	n := flag.Bool("n", false, "omit trailing newline")
	sep := flag.String("s", " ", "separator")

	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
