package main

import "fmt"

/**
 * @Author kong
 * @Date 2024/3/15
 * @Description 动态类型逃逸
 */

func main() {
	s := "Escape"
	fmt.Println(s)
	// 很多函数参数为interface或any类型，比如：fmt.Println(a ...any)，编译期间很难确定其参数的具体类型，也会产生逃逸。
}

// .\03、动态类型逃逸.go:13:13: inlining call to fmt.Println
// .\03、动态类型逃逸.go:13:13: ... argument does not escape
// .\03、动态类型逃逸.go:13:13: s escapes to heap
