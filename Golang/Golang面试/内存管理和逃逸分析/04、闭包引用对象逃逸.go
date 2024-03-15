package main

import "fmt"

func Fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	f := Fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Printf("Fibonacci: %d\n", f())
	}
	// Fibonacci() 函数中原本属于局部变量的a和b，由于闭包的引用，不得不将二者放到堆上，所以导致产生逃逸。
}

// .\04、闭包引用对象逃逸.go:16:13: inlining call to fmt.Printf
// .\04、闭包引用对象逃逸.go:6:2: moved to heap: a
// .\04、闭包引用对象逃逸.go:6:5: moved to heap: b
// .\04、闭包引用对象逃逸.go:7:9: func literal escapes to heap
// .\04、闭包引用对象逃逸.go:14:16: func literal does not escape
// .\04、闭包引用对象逃逸.go:16:13: ... argument does not escape
// .\04、闭包引用对象逃逸.go:16:34: ~R0 escapes to heap
