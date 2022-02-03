package main

import (
	"fmt"
	"strings"
	//	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {

	m := make(map[string]int)
	t := strings.Fields(s)

	for i, v := range t {
		fmt.Println(i, v)
		m[v]++
	}
	fmt.Println(m)

	return m
}

func main() {
	wc.Test(WordCount)
}
