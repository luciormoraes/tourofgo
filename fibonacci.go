// exercise functions
package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {

	fib, next := 0, 1
	if fib == 0 {
		fmt.Println(fib)
	}
	return func() int {
		current := fib
		fib = fib + next
		next = current
		return fib
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
