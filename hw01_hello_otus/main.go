package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

const targetStr = "Hello, OTUS!"

func main() {
	result := stringutil.Reverse(targetStr)
	fmt.Println(result)
}
